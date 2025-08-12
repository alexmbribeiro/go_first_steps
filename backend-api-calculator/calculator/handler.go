package calculator

import (
	"net/http"

	"github.com/alexmbribeiro/backend-api-calculator/utils"

	"log/slog"
)

func AddHandler(w http.ResponseWriter, r *http.Request) {
	var req TwoNumbersRequest
	if err := utils.ParseJSON(r, &req); err != nil {
		slog.Error("Invalid JSON", "error", err)
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
		return
	}
	res := Add(req)
	utils.WriteJSON(w, http.StatusOK, res)
}

func SubtractHandler(w http.ResponseWriter, r *http.Request) {
	var req TwoNumbersRequest
	if err := utils.ParseJSON(r, &req); err != nil {
		slog.Error("Invalid JSON", "error", err)
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
		return
	}
	res := Subtract(req)
	utils.WriteJSON(w, http.StatusOK, res)
}

func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	var req TwoNumbersRequest
	if err := utils.ParseJSON(r, &req); err != nil {
		slog.Error("Invalid JSON", "error", err)
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
		return
	}
	res := Multiply(req)
	utils.WriteJSON(w, http.StatusOK, res)
}

func DivideHandler(w http.ResponseWriter, r *http.Request) {
	var req DivideRequest
	if err := utils.ParseJSON(r, &req); err != nil {
		slog.Error("Invalid JSON", "error", err)
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
		return
	}
	res, err := Divide(req)
	if err != nil {
		slog.Warn("Division error", "error", err)
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	utils.WriteJSON(w, http.StatusOK, res)
}

func SumHandler(w http.ResponseWriter, r *http.Request) {
	var numbers []float32
	if err := utils.ParseJSON(r, &numbers); err != nil {
		slog.Error("Invalid JSON", "error", err)
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
		return
	}
	res := Sum(numbers)
	utils.WriteJSON(w, http.StatusOK, res)
}
