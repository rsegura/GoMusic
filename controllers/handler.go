package controllers


import (
	"fmt"
	"GoMusic/dblayer"
	"github.com/gin-gonic/gin"
)


type HandlerInterface interface{
	GetMainPage(c *gin.Context)
	GetProducts(c *gin.Context)
	GetPromos(c *gin.Context)
	GetProduct(c *gin.Context)
	GetUsers(c *gin.Context)
	AddUser(c *gin.Context)
	SignIn(c *gin.Context)
	SignOut(c *gin.Context)
	GetCustomerByName(c *gin.Context)
}
type Handler struct {
	db dblayer.DBLayer
}

func NewHandler() (HandlerInterface, error) {
	return NewHandlerWithParams("mysql", "root:root@tcp(127.0.0.1:3306)/gomusic")
}


func NewHandlerWithParams(dbtype, conn string) (HandlerInterface, error) {
	db, err := dblayer.NewORM(dbtype, conn)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return &Handler{
		db: db,
	}, nil
}