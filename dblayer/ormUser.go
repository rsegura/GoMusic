package dblayer

import (
	"GoMusic/models"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	//"fmt"
	"errors"
)




func (db *DBORM) AddUser(customer models.Customer) (models.Customer, error){
	hashPassword(&customer.Pass)
	customer.LoggedIn = true
	return customer, db.Create(&customer).Error
}

func (db *DBORM) GetUsers()(users []models.Customer, err error){
	return users, db.Find(&users).Error
}

func (db *DBORM) GetCustomerByName(firstname string, lastname string) (customer models.Customer, err error) {
	return customer, db.Where(&models.Customer{FirstName:firstname, LastName:lastname}).Find(&customer).Error
}

func (db *DBORM) GetCustomerByID(id int) (customer models.Customer, err error) {
	return customer, db.First(&customer, id).Error
}

func (db *DBORM) SignInUser(email, pass string)(customer models.Customer, err error){
	//Obtain a *gorm.DB object representing our customer's row
	result := db.Table("Customers").Where(&models.Customer{Email: email})
	err = result.First(&customer).Error
	if err != nil {
		return customer, err
	}

	if !checkPassword(customer.Pass, pass) {
		return customer, ErrINVALIDPASSWORD
	}
	//update the loggedin field
	err = result.Update("loggedin", 1).Error
	if err != nil {
		return customer, err
	}
	//return the new customer row
	return customer, result.Find(&customer).Error
}

func (db *DBORM) SignOutUserById(id int) error {
	customer := models.Customer{
		Model : gorm.Model{
			ID: uint(id),
		},
	}
	return db.Table("customers").Where(&customer).Update("loggedin", 0).Error
}

func hashPassword(s *string) error {
	if s == nil {
		return errors.New("Reference provided for hashing password is nil")
	}
	//converd password string to byte slice
	sBytes := []byte(*s)
	//Obtain hashed password
	hashedBytes, err := bcrypt.GenerateFromPassword(sBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	//update password string with the hashed version
	*s = string(hashedBytes[:])
	return nil
}

func checkPassword(existingHash, incomingPass string) bool{
	return bcrypt.CompareHashAndPassword([]byte(existingHash), []byte(incomingPass)) == nil
}