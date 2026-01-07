package utils

import (
    "golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

// CheckPasswordHash compares a plaintext password with a bcrypt hashed password
func CheckPasswordHash(password, hashed string) bool {
    return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)) == nil
}