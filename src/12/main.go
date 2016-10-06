package main

import (
	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin/binding"
	//"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type System_user struct {
	dlina    int
	word     string
	site_out []string
}

type Response struct {
	FoundAtSite string
}

// Binding from JSON
type Request struct {
	Site       []string `form:"site" json:"site" binding:"required"` // Slice of strings: https://blog.golang.org/go-slices-usage-and-internals
	SearchText string   `form:"searchtext" json:"searchtext" binding:"required"`
}

func find_string(sayt, word string) bool {

	res, err := http.Get(sayt)
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	content := string(robots)
	if strings.Contains(content, word) {
		return true
	}
	return false
}

func main() { //1
	router := gin.Default()
	router.POST("/checkText", func(c *gin.Context) { //2
		var json Request
		var res Response
		var sys_us System_user
		if c.BindJSON(&json) == nil { //3
			sys_us.dlina = len(json.Site)        //смотрим сколько пришло сайтов
			for i := 0; i <= sys_us.dlina; i++ { //4 цыкл количество сайтов
				if find_string(json.Site[i], json.SearchText) { //5 ищем слово в сайте
					res.FoundAtSite = json.Site[i]
					c.JSON(200, gin.H{
						"FoundAt11Site": res.FoundAtSite,
					})
					break
					//res.FoundAtSite = append(res.FoundAtSite, json.Site[i]) //добавляем если найдено
				} else {
					c.JSON(204, gin.H{
						"status": "No Content",
					})

				} //5
			} //4
			//			if res.FoundAtSite == "" { //6 если не пустое выводим данные
			//				c.JSON(200, gin.H{
			//					"FoundAt22Site": res.FoundAtSite,
			//				})
			//			} else { //в обратном случае ошибка
			//				c.JSON(204, gin.H{
			//					"status": "No vContent",
			//				})
			//			}
		} //3
	}) //2 end
	// Listen and server on 0.0.0.0:8080
	router.Run(":8080")
} //1 end
