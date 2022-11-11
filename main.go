package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()
	router.GET("/service/rest/v1/status/check", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	router.POST("/service/rest/v1/components", func(c *gin.Context) {
		file, _ := c.FormFile("raw.asset1")
		os.MkdirAll(fmt.Sprintf("./data/%s/%s", c.Query("repository"), c.PostForm("raw.directory")), os.ModePerm)
		c.SaveUploadedFile(file, fmt.Sprintf("./data/%s/%s/%s", c.Query("repository"), c.PostForm("raw.directory"), c.PostForm("raw.asset1.filename")))
		c.Status(http.StatusNoContent)
	})
	router.StaticFS("/repository", http.Dir("./data"))
	router.Run(":8081")
}
