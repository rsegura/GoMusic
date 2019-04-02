package dblayer

import (
	"errors"

	"GoMusic/models"
)

type DBLayer interface {
	GetAllProducts() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
	GetUsers() ([]models.Customer, error)
	//GetCustomerByName(string, string) (models.Customer, error)
	//GetCustomerByID(int) (models.Customer, error)
	GetProduct(int) (models.Product, error)
	AddUser(models.Customer) (models.Customer, error)
	SignInUser(username, password string) (models.Customer, error)
	SignOutUserById(int) error
	/*GetCustomerOrdersByID(int) ([]models.Order, error)
	AddOrder(models.Order) error
	GetCreditCardCID(int) (string, error)
	SaveCreditCardForCustomer(int, string) error*/
}

var ErrINVALIDPASSWORD = errors.New("Invalid password")