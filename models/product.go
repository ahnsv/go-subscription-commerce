package models

// struct for products
type Product struct {
	name         string `json: "name"`
	url          string `json: "url"`
	price        int32  `json: "price"`
	image        string `json: "image"`
	is_available bool   `json: "availability"`
	is_subable   bool   `json: "subscribabilty"`
}
