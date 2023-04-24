package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/skytsai2/router"
)

func main() {
	r := gin.Default()

	// r.Use(middleware.Auth())

	// 建立 route
	router.Create(r)

	//啟動
	fmt.Println("listen and serve on 0.0.0.0:8080")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
