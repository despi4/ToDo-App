package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	userdto "todo-app/internal/transport/http/dto/user"
)

// Positioning the middleware
// middleware -> servemux -> application handler
// If position of middleware before the servemux
// it will act on every request that app receives

// Position the middleware after the servemux
// servemux -> middleware -> app handler
// This would cause middleware to only execute for a specific route

func SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-security-policy", "default-src 'self'; style-src 'self';")
		w.Header().Set("referrer-policy", "origin-when-cross-origin")
		w.Header().Set("x-content-type-optoins", "nosniff")
		w.Header().Set("x-frame-options", "deny")
		w.Header().Set("x-xss-protection", "0")

		next.ServeHTTP(w, r)
	})
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())

		next.ServeHTTP(w, r)
	})
}

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Recover", err)
				w.Header().Add("connection", "close")

				errRes := userdto.ErrorResponse{
					Error:     "Internal Server Error",
					ErrorCode: http.StatusInternalServerError,
				}

				w.WriteHeader(http.StatusInternalServerError)

				json.NewEncoder(w).Encode(&errRes)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
