package api

import (
	r "github.com/ahnsv/go-subscription-commerce/api/infrastructure"
)

func Init() {
	r.NewRouter()
}
