package model

import "time"

// type Category struct {
// 	ID   uint   `json:"id"`
// 	Name string `json:"name"`
// }

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Orders struct {
	ID          uint       `json:"id"`
	User        User       `json:"user"`
	Productname string     `json:"name"`
	Payment     float64    `json:"payment"`
	Time        *time.Time `json:"time"`
}
