package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "rahasia"

	hash, _ := HashPassword(password)

	fmt.Println("password: ", password)
	fmt.Println("hash: ", hash)

	match := CheckPassword(password, hash)
	fmt.Println("match:", match)
}

func HashPassword(p string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(p), 14)
	return string(bytes), err
}

func CheckPassword(p string, h string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	return err == nil
}
