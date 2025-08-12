package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type contextKey string

const requestIDKey contextKey = "requestID"

func RequestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reqID := uuid.New().String()

		ctx := context.WithValue(r.Context(), requestIDKey, reqID)
		r = r.WithContext(ctx)

		w.Header().Set("X-Request-ID", reqID)

		next.ServeHTTP(w, r)
	})
}

func GetRequestID(r *http.Request) string {
	if reqID, ok := r.Context().Value(requestIDKey).(string); ok {
		return reqID
	}
	return ""
}
