package Admin

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go/Utils"
	"log"
	"net/http"
)

func AddProduct(context *gin.Context) {
	var newProduct Utils.Product

	if err := context.BindJSON(&newProduct); err != nil {
		return
	}

	Utils.Products = append(Utils.Products, newProduct)
	context.IndentedJSON(http.StatusCreated, newProduct)
}

func DeleteProduct(context *gin.Context) {
	id := context.Query("id")
	product, err := DelProductById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, product)
}

func RemoveProduct(s []Utils.Product, i int) []Utils.Product {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func DelProductById(id string) (*Utils.Product, error) {
	for i, p := range Utils.Products {
		if p.ID == id {
			Utils.Products = RemoveProduct(Utils.Products, i)
			return &p, nil
		}
	}
	return nil, errors.New("product not found")
}

func UpdateProduct(context *gin.Context) {
	id := context.Query("id")
	updatedProd := &Utils.Product{}

	err := context.BindJSON(updatedProd)
	if err != nil {
		return
	}

	prod, err := UpdateById(id, updatedProd)
	if err != nil {
		log.Fatalln(err)
	}

	context.IndentedJSON(http.StatusOK, prod)
}

func UpdateById(id string, updatedProd *Utils.Product) (*Utils.Product, error) {
	for i, p := range Utils.Products {
		if p.ID == id {
			Utils.Products[i].Name = updatedProd.Name
			Utils.Products[i].Category = updatedProd.Category
			Utils.Products[i].Image = updatedProd.Image
			Utils.Products[i].Rated = updatedProd.Rated
			return &Utils.Products[i], nil
		}
	}
	return nil, errors.New("product not found")
}
