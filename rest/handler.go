package rest

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"GoMusic/dblayer"
	"GoMusic/models"
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

func (h *Handler) GetMainPage(c *gin.Context) {
	log.Println("Main page....")
	c.String(http.StatusOK, "Main page for secure API!!")
	//fmt.Fprintf(c.Writer, "Main page for secure API!!")
}

func (h *Handler) GetProducts(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	products, err := h.db.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Found %d products\n", len(products))
	c.JSON(http.StatusOK, products)
}

func (h *Handler) GetPromos(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	promos, err := h.db.GetPromos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, promos)
}

func (h *Handler) GetProduct(c *gin.Context){
	p := c.Params.ByName("id")
	
	if h.db == nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	product, err := h.db.GetProduct(id)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *Handler) GetUsers(c *gin.Context){
	if h.db == nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return 
	}
	customers, err := h.db.GetUsers()
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}

func (h *Handler) AddUser(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	var customer models.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer, err = h.db.AddUser(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func (h *Handler) SignIn(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	var customer models.Customer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer, err = h.db.SignInUser(customer.Email, customer.Pass)
	if err != nil {
		if err == dblayer.ErrINVALIDPASSWORD {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func (h *Handler) SignOut(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.db.SignOutUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func (h *Handler) GetCustomerByName(c *gin.Context){
	firstname := c.Param("firstname")
	lastname := c.Param("lastname")
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	customer, err := h.db.GetCustomerByName(firstname, lastname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)

}
/*func (h *Handler) GetOrders(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orders, err := h.db.GetCustomerOrdersByID(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *Handler) Charge(c *gin.Context) {
	if h.db == nil {
		return
	}

}*/
