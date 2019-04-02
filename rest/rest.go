package rest

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"GoMusic/middlewares"
)

func RunAPI(address string) error {
	fmt.Println("entramos en RunAPI");
	//Get gin's default engine
	r := gin.Default()
	r.Use(middlewares.MyCustomLogger())
	//Define a handler
	h, _ := NewHandler()
	//load homepage
	r.GET("/", h.GetMainPage)
	//get products
	r.GET("/products", h.GetProducts)
	//get promos
	r.GET("/product/:id", h.GetProduct)

		/*
		//post user sign in
		r.POST("/user/signin", h.SignIn)
		//post user sign out
		r.POST("/user/:id/signout", h.SignOut)
		//get user orders
		r.GET("/user/:id/orders", h.GetOrders)
		//post purchase charge
		//r.POST("/user/charge", h.Charge)
		*/
	

	userGroup := r.Group("/user")
	{
		userGroup.POST("/:id/signout", h.SignOut)
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
