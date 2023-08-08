package main

import (
	"github.com/gin-gonic/gin"
	"go/Admin"
	"go/Utils"
)

func main() {
	router := gin.Default()
	router.GET("/products", Utils.GetProducts)
	router.GET("/product", Utils.GetSingleProduct)
	router.POST("/addProduct", Admin.AddProduct)
	router.DELETE("/deleteProduct", Admin.DeleteProduct)
	router.PUT("/updateProduct", Admin.UpdateProduct)
	err := router.Run("localhost:9090")
	if err != nil {
		return
	}
}
