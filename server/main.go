package main

import (
	"fmt"
	"landingpage/config"
	"landingpage/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()
	r := gin.Default()

	routes.Routes(r)

	fmt.Println("server running at http://localhost:8080")
	r.Run(":8080")
}
