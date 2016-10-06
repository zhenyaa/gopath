package main

import (
	"github.com/gin-gonic/gin"
	//"github.com/gin-gonic/gin/binding"
	"fmt"
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
	var content = string(robots)
	fmt.Println(content)
	if strings.Contains(content, word) {
		return true
	}
	return false
}
func send_str(sait_mass []string, word2 string) string {
	len_mass := len(sait_mass)
	for i := 0; i <= len_mass; i++ {
		fmt.Println(sait_mass[i], word2)
		if find_string(sait_mass[i], word2) {
			fmt.Println("return mass")
			return sait_mass[i]
		}
	}
	fmt.Println("return non!!!")
	return "non"
}
func main() { //1
	router := gin.Default()
	router.POST("/checkText", func(c *gin.Context) { //2
		var json Request
		var res Response
		//		var sys_us System_user
		if c.BindJSON(&json) == nil { //3
			//					"FoundAt22Site": res.FoundAtSite,
			res.FoundAtSite = send_str(json.Site, json.SearchText)
			fmt.Println(res.FoundAtSite, "this")
			if res.FoundAtSite == "non" {
				c.JSON(204, gin.H{
					"status": "No Content",
				})
			} else {
				c.JSON(200, gin.H{
					"FoundAt11Site": res.FoundAtSite,
				})
			} //3
		}
	}) //2 end
	router.Run(":8080")
}
