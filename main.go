package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	Test(router)

	router.Run()
}

func Test(r *gin.Engine) {
	r.GET("GHP/test", func(c *gin.Context) {
		c.String(200, "TEST!?!??!?!?!?!??!?!?!?!?!?!?!?!?!?!?!?!")
	})
}
