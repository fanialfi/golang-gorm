package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect database : %s", err.Error()))
	}

	// migrate the schema
	db.AutoMigrate(&Product{})

	// create
	db.Create(&Product{Code: "D42", Price: 100})

	// read
	var product Product
	_ = db.First(&product, 1)
	fmt.Printf("product is %#v\n\n", product)
	_ = db.First(&product, "code = ?", "D42")
	fmt.Printf("product is %#v\n", product)

	// update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	fmt.Printf("%#v\t%#v\t%#v\n", product.Code, int(product.ID), int(product.Price))

	// update - update multiple fields
	db.Model(&product).Updates(Product{Code: "HEXDUMP", Price: 3000})
	fmt.Printf("%#v\t%#v\t%#v\n", product.Code, int(product.ID), int(product.Price))

	db.Model(&product).Updates(map[string]any{"Code": "APAINI", "Price": 1500})
	fmt.Printf("%#v\t%#v\t%#v\n", product.Code, int(product.ID), int(product.Price))

	// delete - delete product
	db.Delete(&product, 1)
}
