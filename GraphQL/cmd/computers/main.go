package main

import (
	"fmt"
	"net/http"
)

type Computers struct {
	ID          int     `json:id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

var computers = []Computers{
	{
		ID:          0,
		Name:        "Asus",
		Description: "A asus computer",
		Price:       500.00,
	},
	{
		ID:          1,
		Name:        "Dell",
		Description: "A dell computer",
		Price:       800.00,
	},
	{
		ID:          2,
		Name:        "HP",
		Description: "A hp computer",
		Price:       700.00,
	},
}

func computerHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/computers", computerHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
