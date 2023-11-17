package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


func main() {
	password := "12345"
	 hash, _ := HashPassword(password) // ignore error for the sake of simplicity
	// hash1, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", hash)
	// fmt.Println("Hash:    ", hash1)

	//match := CheckPasswordHash(hash, hash1)
	
	// fmt.Println()
	fmt.Println("Match: ", CheckPasswordHash("0622", "$2a$10$YC7fwolCEVR7nHGS2Et8r.WfdCezViKE.TEG9Chomo22TooCoHypy"))
	//fmt.Println("Match: ", CheckPasswordHash(password,hash))

	
}
