package main

import (
	"github.com/gin-gonic/gin"
	// static "github.com/raggaer/gin-static"
)

func main() {
	r := gin.Default()
	// r.Use(static.ServeStaticFiles("/", static.LocalFile("./web/dist", true)))
	// grp1 := r.Group("/api/v1")
	// {
	// }
	r.GET("/api/tab", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "nine",
		})
	})
	r.GET("/api/panic", func(c *gin.Context) {
		panic("aaaaaaaaaaaaaa!")
	})
	r.Static("/web/dist/hello-gin", "./web/dist/hello-gin")
	// r.GET("", func(c *gin.Context) {
	// 	c.Redirect(302, "/hello-gin/dist/hello-gin/index.html")
	// })
	// r.Static("/static2", "./static2")
	// r.StaticFile("/", "./index.html")
	// r.GET("", func(c *gin.Context) {
	// 	c.Redirect(302, "/static/index.html")
	// })
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(302, "/web/dist/hello-gin/index.html")
		// c.Redirect(302, "/")
		// c.Request.URL.Path = "/"
		// c.File("./static/index.html")
	})
	r.Run()
	// listen and serve on 0.0.0.0:8080
}
