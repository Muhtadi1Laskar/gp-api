package main

import (
	"fmt"
	"go-api/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/aes/encrypt", handlers.EncryptAES)
	http.HandleFunc("/aes/decrypt", handlers.DecryptAES)
	http.HandleFunc("/cipher/xor", handlers.XOREncrypt)
	http.HandleFunc("/cipher/xorDecrypt", handlers.XORDecrypt)
	http.HandleFunc("/cipher/ceaserCipher", handlers.CeaserCipher)
	fmt.Println("Server running on 5000")
	http.ListenAndServe(":5000", nil)
}