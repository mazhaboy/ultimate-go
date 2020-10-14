package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	a, _ := bcrypt.GenerateFromPassword([]byte("hello"), bcrypt.DefaultCost)
	b := []byte("hello")
	if err := bcrypt.CompareHashAndPassword(a, b); err == nil {
		
	}

}
