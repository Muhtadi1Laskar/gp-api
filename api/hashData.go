package handler

import (
	"go-api/handlers"
	"net/http"
)

func HandlerOne(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/hashData", handlers.HashData)
	mux.HandleFunc("/api/aesEncrypt", handlers.EncryptAES)
	mux.HandleFunc("/api/aesDecrypt", handlers.DecryptAES)
	mux.HandleFunc("/api/verifyData", handlers.VerifyData)
	mux.HandleFunc("/api/xorEncrypt", handlers.XOREncrypt)
	mux.HandleFunc("/api/xorDecrypt", handlers.XOREncrypt)
	mux.HandleFunc("/api/ceaserCipher", handlers.CeaserCipherHandler)
	mux.HandleFunc("/api/vignereCipher", handlers.VigenereCipherHandler)
	handlers.EnableCORS(mux).ServeHTTP(w, r)
	// handlers.HashData(w, r)
}

func HandlerTwo(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/aesEncrypt", handlers.EncryptAES)
	handlers.EnableCORS(mux).ServeHTTP(w, r)
}
