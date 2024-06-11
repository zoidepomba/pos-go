package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	//create
	//db.Create(&Product{
	//	Name: "Notebook",
	//	Price: 1000.00,
	//})
	//
	////create batch
	//products := []Product{
	//	{Name: "Notebook", Price: 1001.00},
	//	{Name: "Mouse", Price: 11.00},
	//	{Name: "Teclado", Price: 51.00},
	//	{Name: "Monitor", Price: 500.00},
	//}
	//db.Create(&products)
	//select one
	//var product Product
	//db.First(&product, 2)
	//fmt.Println(product)

	//db.First(&product, "name = ?", "Mouse")
	//fmt.Println(product)

	//select all
	//var products []Product
	//db.Limit(2).Offset(1).Find(&products)
	//for _, product := range products {
	//	fmt.Println(product)
	//}
	//where
	//var products []Product
	//db.Where("price > ?", 999).Find(&products)
	//for _, product := range products{
	//	fmt.Println(product)
	//}
	//like
	//var products []Product
	//db.Where("name LIKE ?", "%book%").Find(&products)
	//for _, product := range products{
	//	fmt.Println(product)
	//}

	var p Product
	db.First(&p, 1)
	p.Name = "New mouse"
	db.Save(&p)

	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p.Name)
	//
	db.Delete(&p2)
}
