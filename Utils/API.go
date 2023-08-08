package Utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Products = []Product{
	{ID: "1", Name: "t-shirt", Category: "clothes", Rated: true, Image: "public/1.jpg"},
	{ID: "2", Name: "t-shirt", Category: "clothes", Rated: true},
	{ID: "3", Name: "t-shirt", Category: "clothes", Rated: true},
	{ID: "4", Name: "t-shirt", Category: "clothes", Rated: true},
	{ID: "5", Name: "t-shirt", Category: "clothes", Rated: true},
}

func GetProducts(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, Products)
}

func GetSingleProduct(context *gin.Context) {
	id := context.Query("id")
	product, err := GetProductById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, product)
}

func GetProductById(id string) (*Product, error) {
	for i, p := range Products {
		if p.ID == id {
			return &Products[i], nil
		}
	}

	return nil, errors.New("product not found")
}
