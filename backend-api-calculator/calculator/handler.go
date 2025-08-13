package calculator

import (
	"net/http"
	"time"

	"github.com/alexmbribeiro/backend-api-calculator/utils"

	"log/slog"
)

func AddHandler(logRepo *Repository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        var req TwoNumbersRequest
        if err := utils.ParseJSON(r, &req); err != nil {
            slog.Error("Invalid JSON", "error", err)
            utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
            return
        }

        res := Add(req)

        logRepo.Save(r.Context(), LogEntry{
            Operation: "add",
            Input:     req,
            Result:    res,
            Time:      time.Now(),
        })

        utils.WriteJSON(w, http.StatusOK, res)
    }
}


func SubtractHandler(logRepo *Repository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

		var req TwoNumbersRequest
		if err := utils.ParseJSON(r, &req); err != nil {
			slog.Error("Invalid JSON", "error", err)
			utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
			return
		}

		res := Subtract(req)

        logRepo.Save(r.Context(), LogEntry{
            Operation: "subtract",
            Input:     req,
            Result:    res,
            Time:      time.Now(),
        })

		utils.WriteJSON(w, http.StatusOK, res)
	}
}

func MultiplyHandler(logRepo *Repository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
	
		var req TwoNumbersRequest
		if err := utils.ParseJSON(r, &req); err != nil {
			slog.Error("Invalid JSON", "error", err)
			utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
			return
		}
	
		res := Multiply(req)

		logRepo.Save(r.Context(), LogEntry{
            Operation: "multiply",
            Input:     req,
            Result:    res,
            Time:      time.Now(),
        })
	
		utils.WriteJSON(w, http.StatusOK, res)
	}
}

func DivideHandler(logRepo *Repository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

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

			logRepo.Save(r.Context(), LogEntry{
        	    Operation: "divide",
    	        Input:     req,
	            Result:    err,
            	Time:      time.Now(),
        	})
			return
		}

		logRepo.Save(r.Context(), LogEntry{
            Operation: "divide",
            Input:     req,
            Result:    res,
            Time:      time.Now(),
        })

		utils.WriteJSON(w, http.StatusOK, res)
	}
}

func SumHandler(logRepo *Repository) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

		var numbers []float32
		if err := utils.ParseJSON(r, &numbers); err != nil {
			slog.Error("Invalid JSON", "error", err)
			utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON"})
			return
		}
	
		res := Sum(numbers)
		

		logRepo.Save(r.Context(), LogEntry{
            Operation: "sum",
            Input:     numbers,
            Result:    res,
            Time:      time.Now(),
        })

		utils.WriteJSON(w, http.StatusOK, res)
	}
}
