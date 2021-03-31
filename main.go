// @program:     debug
// @file:        main.go
// @author:      edte
// @create:      2021-03-30 21:59
// @description:
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "test",
		})
	})
	_ = r.Run()
}
