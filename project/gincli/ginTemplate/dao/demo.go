package dao

import (
	"ginTemplate/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Demo() {
	db, err := gorm.Open("mysql", "root:123456@tcp(10.229.24.23:3306)/zstest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(fmt.Errorf("%v", err))
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&models.Product{})

	// Create
	db.Create(&models.Product{Code: "L1212", Price: 1000})

	// Read
	var product models.Product
	db.First(&product, 1) // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)
}


