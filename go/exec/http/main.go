package main

import (
	"exec/go/exec/http/gee"
)

func main() {

	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.String(200, "%s", "huyufan")
	})
	r.GET("/acount", func(c *gee.Context) {

		c.Json(200, gee.H{
			"nihao": "ni",
		})
	})
	r.Run(":8888")
}
