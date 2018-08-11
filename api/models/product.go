package models

// struct for products
type Product struct {
	ID           int32  `json: "id"`
	name         string `json: "name"`
	url          string `json: "url"`
	price        int32  `json: "price"`
	image        string `json: "image"`
	is_available bool   `json: "availability"`
	is_subable   bool   `json: "subscribabilty"`
}

// demo products
var productList = []Product{
	Product{ID: 1, name: "tumblr", url: "http://naver.com", price: 10001, image: "", is_available: true, is_subable: true},
	Product{ID: 2, name: "keyboard", url: "http://daum.net", price: 10000, image: "", is_available: true, is_subable: true},
}

func AllProducts() []Product {
	return productList
}
