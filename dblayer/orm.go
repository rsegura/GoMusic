package dblayer

import (
	"GoMusic/models"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	//"fmt"
)

//get products
//get promos
//post user sign in
//get user orders
//post user sign out
//post purchase charge

type DBORM struct {
	*gorm.DB
}

func NewORM(dbname, con string) (*DBORM, error) {
	db, err := gorm.Open(dbname, con+"?parseTime=true")
	return &DBORM{
		DB: db,
	}, err
}

func (db *DBORM) GetAllProducts() (products []models.Product, err error) {
	/*product := models.Product{
		Image:"/src/test.jpg",
		//SmallImage:"/src/small/test.jpg",
		ImagAlt:"test",
		Price:200,
		Promotion:50,
		ProductName:"Es un test",
		Description:""}
	db.Create(&product)
	tables := []string{}
	db.Select(&tables, "SHOW TABLES")
	fmt.Println(tables)
	fmt.Println("ENTRAMOS EN GET ALL PRODUCTS")
	fmt.Println(db.HasTable(&products))*/
	return products, db.Table("products").Find(&products).Error
}

func (db *DBORM) GetPromos() (products []models.Product, err error) {
	return products, db.Where("promotion IS NOT NULL").Find(&products).Error
}

func (db *DBORM) GetCustomerByName(firstname string, lastname string) (customer models.Customer, err error) {
	return customer, db.Where(&models.Customer{FirstName:firstname, LastName:lastname}).Find(&customer).Error
}

func (db *DBORM) GetCustomerByID(id int) (customer models.Customer, err error) {
	return customer, db.First(&customer, id).Error
}

func (db *DBORM) GetProduct(id int) (product models.Product, err error) {
	return product, db.First(&product, id).Error
}



func (db *DBORM) GetCustomerOrdersByID(id int) (orders []models.Order, err error) {
	return orders, db.Table("orders").Select("*").Joins("join customers on customers.id = customer_id").Joins("join products on products.id = product_id").Where("customer_id=?", id).Scan(&orders).Error
}
