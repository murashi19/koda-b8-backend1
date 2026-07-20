package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/murashi19/koda-b8-backend1/internal/di"
)

func main() {
	router := gin.Default()

	container, err := di.NewContainer()
	if err != nil {
		log.Fatal(err)
	}
	defer container.Close()

	auth := container.AuthHandler()

	router.POST("/register", auth.Register)
	router.POST("/login", auth.Login)
	router.GET("/users", auth.GetUsers)

	log.Fatal(router.Run(":8080"))
}
