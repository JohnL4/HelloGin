# HelloGin
Horsing around with Go Gin and Angular

The key to serving an Angular (or any, probably) front end is this:

In index.html, specify your base as the base of your generated web app:

```html
<base href="/web/dist/hello-gin/">
```

In your back end, also specify the base path like so:

```go
	r := gin.Default()
	r.Static("/web/dist/hello-gin", "./web/dist/hello-gin") // Serve a static base directory under the given path.
	// All weird routes need a default fallback, per
	// https://angular.io/guide/deployment#routed-apps-must-fall-back-to-indexhtml
	r.NoRoute(func(c *gin.Context) {
		c.Redirect(302, "/web/dist/hello-gin/index.html")
		// c.Redirect(302, "/")
		// c.Request.URL.Path = "/"
		// c.File("./static/index.html")
	})
```
