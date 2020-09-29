package main

import "golang.org/x/crypto/bcrypt"

func ToHash(s string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(s), 14)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
