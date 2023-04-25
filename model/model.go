package model

type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Product struct {
	ID       uint     `json:"id"`
	Category Category `json:"category"`
	Name     string   `json:"name"`
	Price    float64  `json:"price"`
}

type User struct {
	ID        uint   `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Orders struct {
	ID       uint    `json:"id"`
	User     User    `json:"user"`
	Products Product `json:"products"`
	Payment  float64 `json:"payment"`
}
