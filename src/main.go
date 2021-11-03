package main

import (
	"fmt"
	"github.com/tishchenkomaksym/hasher/pkg/hasher"
)

//Package hasher implements 2 functions:
//HashPassword(password string) (string, error),
//CheckPasswordHash(password, hash string) bool

func main() {

	hash, _ := hasher.HashPassword("golang")
	checked := hasher.CheckPasswordHash("golang", hash)
	if checked {
		fmt.Println("Password is correct")
	} else {
		fmt.Println("Password is incorrect")
	}

}
