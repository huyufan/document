package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type bis struct {
	BigNumber json.Number `json:BigNumber`
}

func main() {

	engin := gin.Default()
	engin.GET("/", func(ctx *gin.Context) {
		data := bis{BigNumber: "12345678901234567890"}
		js, _ := json.Marshal(data)

		ctx.Header("Content-Type", "application/json")
		ctx.Status(200)
		ctx.Writer.Write(js)
	})
	engin.Run(":8888")
}
