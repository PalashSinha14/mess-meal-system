package main

//import "fmt"

import (
	"github.com/gin-gonic/gin"
	"github.com/PalashSinha14/mess-meal-system/db"
	"github.com/PalashSinha14/mess-meal-system/routes"
)

func main() {
	//fmt.Println("Hello World")
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}