package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/gin-gonic/gin"
)

func getProducts() []Product {
	raw, err := ioutil.ReadFile("./asset/dummy.json")
	if err != nil {
		log.Fatal(err)
	}
	var products []Product
	json.Unmarshal(raw, &products)
	return products
}

func main() {
	products := getProducts()

	for _, product := range products {
		av, err := dynamodbattribute.MarshalMap(product)

		if err != nil {
			fmt.Println("got error on marshalling map:")
			log.Fatal(err)
		}
		input := &dynamodb.PutItemInput{
			Product:   av,
			TableName: aws.String("test_products"),
		}
		_, err := svc.PutItem(input)
		if err != nil {
			fmt.Println("got error on putting item")
			log.Fatal(err)
		}
		fmt.Print("successfully added '", product.name)
	}

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
