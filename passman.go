package main

import (
    "fmt"
    "math/rand"
    "time"
)

// Function to generate a random password of specified length
func generatePassword(length int) string {
    rand.Seed(time.Now().UnixNano())

    // Define the character set for the password
    charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+"

    password := make([]byte, length)
    for i := range password {
        password[i] = charSet[rand.Intn(len(charSet))]
    }
    return string(password)
}

func main() {
    // Define the length of the password
    passwordLength := 12

    // Generate and print the password
    password := generatePassword(passwordLength)
    fmt.Println("Generated Password:", password)
}
