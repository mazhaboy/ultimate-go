package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	a, _ := bcrypt.GenerateFromPassword([]byte("hello"), bcrypt.DefaultCost)
	fmt.Println(string(a))
	
}
