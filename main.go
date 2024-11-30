package main

import (
	"fmt"
	"go-api/handlers"
	"net/http"
)

func enableCORS(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        h.ServeHTTP(w, r)
    })
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hash", handlers.HashData)
	mux.HandleFunc("/aes/encrypt", handlers.CipherAES)
	mux.HandleFunc("/cipher/xor", handlers.CipherXOR)
	mux.HandleFunc("/cipher/ceaserCipher", handlers.CeaserCipherHandler)
	mux.HandleFunc("/cipher/vignereCipher", handlers.VigenereCipherHandler)
	mux.HandleFunc("/cipher/atbashCipher", handlers.AtbashHandler)
	mux.HandleFunc("/cipher/rot13Cipher", handlers.Rot13Handler)
	mux.HandleFunc("/cipher/baconCipher", handlers.BaconCipherHandler)
	mux.HandleFunc("/encoder/binaryEncoder", handlers.BinaryEncoderHander)
	fmt.Println("Server running on 5000")
	http.ListenAndServe(":5000", enableCORS(mux))
}