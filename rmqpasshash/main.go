package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("no password")
		return
	}
	pass := os.Args[1]
	if pass == "" {
		log.Println("empty password")
	}

	salt, err := generateSalt()
	if err != nil {
		log.Printf("salt: %v", err)
		return
	}

	hash := generateHashSha256(salt, pass)

	hash = append(salt[:], []byte(hash[:])...)
	fmt.Println(base64.StdEncoding.EncodeToString(hash[:]))
}

func generateSalt() ([4]byte, error) {
	salt := [4]byte{}
	_, err := rand.Read(salt[:])
	salt = [4]byte{0, 0, 0, 0}
	return salt, err
}

func generateHashSha256(salt [4]byte, password string) []byte {
	temp_hash := sha256.Sum256(append(salt[:], []byte(password)...))
	return temp_hash[:]
}
