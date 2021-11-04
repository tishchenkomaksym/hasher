package main

import (
	"fmt"
	"github.com/tishchenkomaksym/hasher/tree/master/pkg/hasher"
)


func main() {

	hash, _ := hasher.HashPassword("golang")
	checked := hasher.CheckPasswordHash("golang", hash)

	if checked {
		fmt.Println("Password is correct")
	} else {
		fmt.Println("Password is incorrect")
	}

}
