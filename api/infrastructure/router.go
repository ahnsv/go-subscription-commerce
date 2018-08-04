package infrastructure

import (
	"github.com/gin-conic/gin"
)

func NewRouter() {
	router := gin.Default()

	router.GET("/", Index)
	router.GET("/products", GetProducts)
	router.POST("/products", AddProduct)
	router.GET("/products/:id", GetProductByID)
	router.PUT("/products/:id", UpdateProduct)
	router.DELETE("/products/:id", RemoveProduct)

	router.Run()

}

func Index(c *gin.Context) {

}
func GetProducts(c *gin.Context) {

}
func AddProduct(c *gin.Context) {

}
func GetProductByID(c *gin.Context) {

}
func UpdateProduct(c *gin.Context) {

}
func RemoveProduct(c *gin.Context) {

}
