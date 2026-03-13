package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// Name     string `json:"name"`
	Email string `json:"email" gorm:"unique" binding:"required,email"`
	// Password string `json:"password"`
	Name string `json:"name"`
	// Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserCreReq struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name"`
	Password string `json:"password" binding:"required"`
}

type LoginReq struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
