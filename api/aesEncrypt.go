package handler

// import (
// 	"go-api/handlers"
// 	"net/http"
// )

// func HandlerThree(w http.ResponseWriter, r *http.Request) {
// 	mux := http.NewServeMux()

// 	mux.HandleFunc("/api/hashData", handlers.EncryptAES)
// 	handlers.EnableCORS(mux).ServeHTTP(w, r)

// 	handlers.EncryptAES(w, r)
// }