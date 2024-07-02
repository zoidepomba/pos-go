package main

import (
	"net/http"

	"github.com/zoidepomba/pos-go/tree/main/9/configs"
	"github.com/zoidepomba/pos-go/tree/main/9/internal/entity"
	"github.com/zoidepomba/pos-go/tree/main/9/internal/infra/database"
	"github.com/zoidepomba/pos-go/tree/main/9/internal/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8000", nil)
	println(config.DBDriver)
}
