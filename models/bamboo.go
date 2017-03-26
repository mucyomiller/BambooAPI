package models

import (
	"encoding/json"
)

type User struct {
	Id        int64
	Firstname string
	Lastname  string
	Username  string
	password  int64
}

type Menu struct {
	Name    string
	Picture string
	Items   []Item
}
type Item struct {
	Name    string
	Price   uint
	Picture string
}
