package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/alexmbribeiro/backend-api-calculator/calculator"
	"github.com/alexmbribeiro/backend-api-calculator/middleware"
	"github.com/alexmbribeiro/backend-api-calculator/db"

	"github.com/rs/cors"
)

func main() {
	redisAddr := os.Getenv("REDIS_ADDR") // e.g. "localhost:6379"
    rdb := db.ConnectRedis(redisAddr, "", 0)
	logRepo := calculator.NewRepository(rdb)


	mux := http.NewServeMux()

	mux.HandleFunc("/add", calculator.AddHandler(logRepo))
	mux.HandleFunc("/subtract", calculator.SubtractHandler(logRepo))
	mux.HandleFunc("/multiply", calculator.MultiplyHandler(logRepo))
	mux.HandleFunc("/divide", calculator.DivideHandler(logRepo))
	mux.HandleFunc("/sum", calculator.SumHandler(logRepo))

	handler := cors.Default().Handler(middleware.AuthMiddleware(middleware.RequestIDMiddleware(mux)))

	slog.Info("Starting server on :3000")
	if err := http.ListenAndServe(":3000", handler); err != nil {
		slog.Error("Server failed", "error", err)
	}
}
