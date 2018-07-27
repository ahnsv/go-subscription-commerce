package backend

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
)

func main() {
	config := &aws.Config{
		Region:   aws.String("us-west-2"),
		Endpoint: aws.String("http://localhost:8080"),
	}

	sess := session.Must(session.NewSession(config))

	svc := dynamodb.New(sess)

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{},
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
