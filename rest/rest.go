package rest

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"GoMusic/middlewares"
	"GoMusic/controllers"
)

func RunAPI(address string) error {
	fmt.Println("entramos en RunAPI");
	//Get gin's default engine
	r := gin.Default()
	r.Use(middlewares.MyCustomLogger())
	//Define a handler
	h, _ := controllers.NewHandler()
	//load homepage
	r.GET("/", h.GetMainPage)
	//get products
	r.GET("/products", h.GetProducts)
	//get promos
	r.GET("/product/:id", h.GetProduct)

		/*
		TO DO
		add product
		remove product
		purchase product
		get orders
		add config files
		securize endpoints with tokens
	
		*/
	

	userGroup := r.Group("/user")
	{
		userGroup.POST("/:id/signout", h.SignOut)
		userGroup.GET("/:firstname/:lastname", h.GetCustomerByName)
		//userGroup.GET("/:id/orders", h.GetOrders)
	}

	usersGroup := r.Group("/users")
	{
		//usersGroup.POST("/charge", h.Charge)
		usersGroup.POST("/signin", h.SignIn)
		usersGroup.POST("", h.AddUser)
		usersGroup.GET("", h.GetUsers)
	}
	return r.Run(address)
}
