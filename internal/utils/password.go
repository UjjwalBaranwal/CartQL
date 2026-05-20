// Package utils provides utility functions for password hashing and verification, as well as JWT token generation and validation, to support authentication and security features in the CartQL application.
package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword hash password using bcrypt at default cost
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword checks if password hash is correct
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
