package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ahnsv/go-subscription-commerce/models"
	"github.com/ahnsv/go-subscription-commerce/db"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Product Product
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
}
