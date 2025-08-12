package main

import (
	"log/slog"
	"net/http"

	"github.com/alexmbribeiro/backend-api-calculator/calculator"
	"github.com/alexmbribeiro/backend-api-calculator/middleware"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/add", calculator.AddHandler)
	mux.HandleFunc("/subtract", calculator.SubtractHandler)
	mux.HandleFunc("/multiply", calculator.MultiplyHandler)
	mux.HandleFunc("/divide", calculator.DivideHandler)
	mux.HandleFunc("/sum", calculator.SumHandler)

	handler := cors.Default().Handler(middleware.AuthMiddleware(mux))

	slog.Info("Starting server on :3000")
	if err := http.ListenAndServe(":3000", handler); err != nil {
		slog.Error("Server failed", "error", err)
	}
}
