package handler

// import (
// 	"go-api/handlers"
// 	"net/http"
// )

// func HandlerFour(w http.ResponseWriter, r *http.Request) {
// 	mux := http.NewServeMux()

// 	mux.HandleFunc("/api/hashData", handlers.DecryptAES)
// 	handlers.EnableCORS(mux).ServeHTTP(w, r)
// 	// handlers.DecryptAES(w, r)
// }