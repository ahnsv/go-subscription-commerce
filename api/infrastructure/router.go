package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func NewRouter() {
	router := gin.Default()

	router.LoadHTMLGlob("api/templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Home Page",
			},
		)
	})
	// router.GET("/products", GetProducts)
	// router.POST("/products", AddProduct)
	// router.GET("/products/:id", GetProductByID)
	// router.PUT("/products/:id", UpdateProduct)
	// router.DELETE("/products/:id", RemoveProduct)

	router.Run()

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
