package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Doload(url string) bool {
	http.Get(url)
	fmt.Println(url)
	return true
}


func main()  {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {

		go Doload("http://www.baidu.com")
		context.JSON(200,"成功")
	})
	r.Run(":8000")
}
