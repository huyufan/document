package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

func AddProduct(c *gin.Context) {

	c.JSON(200, gin.H{
		"v1": "AddProduct",
		"hs": 1000000000000,
	})

}

func AddProductd() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"v1": "AddProduct",
		})
		//c.ShouldBind()
	}

}

func test() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	group1 := r.Group("v1")
	{
		group1.Any("/product/add", AddProduct)
		group1.Any("/product/adt", AddProductd())
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}
func main() {
	js := `{"name":"huyufan","id":1000000000}`
	var data interface{}
	json.Unmarshal([]byte(js), &data)
	fmt.Println(data)
	// st := struct {
	// 	Name string `json:"name"`
	// 	Id   int    `json:"id"`
	// }{"huyufan", 1000000000}
	// sh, err := json.Marshal(st)
	// fmt.Println(err)
	// fmt.Println((string)(sh))
	// json.NewEncoder(os.Stdout).Encode(st)

	// fmt.Println(st)

}
