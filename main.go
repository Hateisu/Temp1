package main

import (
	"marketapi/modules/handlers"

	"github.com/gin-gonic/gin"
)

type human struct {
	Name    string
	Surname string
}

func main() {
	// Create database: preferred ORM
	// Create AUTH system with roles, with all saved in db
	//Main GIN Instance
	r := gin.Default()

	//selleer := storage.Seller{Firstname: "New", Lastname: "Nigga"}
	// Create
	//storage.DB.Create(&selleer)
	//storage.DB.Create(&models.Product{Seller: selleer, Name: "rootten", Price: 1000, Description: "LAys"})

	//log.Printf(consts.RESPONSE_OK_200(human{"Timka", "new"}))
	//Simple Router
	r.GET("/users", handlers.HandleGetAllUsers)
	r.GET("/checkforrq", handlers.CheckForRq)
	r.POST("/addc", handlers.HandleAddCountry)
	r.GET("/getc", handlers.HandleGetCountry)

	r.Run() // listen and serve on localhost:8080
}
