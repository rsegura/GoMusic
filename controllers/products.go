package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)


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