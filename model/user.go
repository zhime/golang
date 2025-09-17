package model

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Address   string    `json:"address"`
	Phone     int       `json:"phone"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`            // 密码不在JSON中返回
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
