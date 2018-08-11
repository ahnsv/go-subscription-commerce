package main

import (
	"github.com/ahnsv/go-subscription-commerce/api"
	"github.com/ahnsv/go-subscription-commerce/db"
	"github.com/ahnsv/go-subscription-commerce/iamport"
)

/*
TODO: define a struct to contain the product data and be sent to the MongoDB storage
~ 7/24
*/

func main() {
	db.Init()
	api.Init()
	iamport.Init()
}
