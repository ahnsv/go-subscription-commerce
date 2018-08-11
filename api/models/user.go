package models

type User struct {
	ID            int           `json:"id"`
	Name          string        `json:"name"`
	Role          string        `json:"role"`
	Subscriptions []interface{} `json:"subscriptions"`
}
