package db

import(
	"sql"
	"fmt"
)

type User struct{
	Id int `json: "u_id" db: "u_id"`
	Email string `json: "email" db: "email"`
	Password zero.string `json: "password" db:"email"`
	Name zero.string `json:"name" db: "name"`
	Surname zero.string `json:"surname" db: "surname"`
}