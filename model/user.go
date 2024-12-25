package model

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Phone   int    `json:"phone"`
	Email   string `json:"email"`
}
