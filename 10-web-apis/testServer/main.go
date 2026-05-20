package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"time"

	"goLangFirst/utils/myutil"
)

// Response represents a standard API response
type Response struct {
	Message string `json:"message"`
	Time    string `json:"time"`
	Status  string `json:"status"`
}

// User represents a user entity
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "Pawan Kumar", Email: "pawan@example.com"},
	{ID: 2, Name: "John Doe", Email: "john@example.com"},
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (r *responseRecorder) WriteHeader(statusCode int) {
	r.statusCode = statusCode
	r.ResponseWriter.WriteHeader(statusCode)
}

func loggingMiddleware(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		recorder := &responseRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(recorder, r)

		logger.Printf(
			"%s %s | status=%d | remote=%s | duration=%s",
			r.Method,
			r.URL.RequestURI(),
			recorder.statusCode,
			r.RemoteAddr,
			myutil.FormatDuration(time.Since(start)),
		)
	})
}

// Health check endpoint
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message: "Server is running!",
		Time:    time.Now().Format(time.RFC3339),
		Status:  "healthy",
	}
	json.NewEncoder(w).Encode(response)
}

// Get all users
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Get user by ID
func getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID parameter is required", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// Simple ID matching (in real app, convert to int and search)
	fmt.Fprintf(w, `{"message": "Get user by ID: %s", "time": "%s"}`, id, time.Now().Format(time.RFC3339))
}

// Root endpoint
func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message: "Welcome to Go Learning Project API!",
		Time:    time.Now().Format(time.RFC3339),
		Status:  "active",
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	port := ":8080"

	logger, closeLogger, err := myutil.NewExecutionLogger()
	if err != nil {
		log.Fatalf("failed to initialize execution logger: %v", err)
	}
	defer closeLogger()

	mux := http.NewServeMux()

	// Application routes
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/api/users", getUsersHandler)
	mux.HandleFunc("/api/user", getUserByIDHandler)

	// Profiling routes for learning and debugging.
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.Handle("/debug/pprof/allocs", pprof.Handler("allocs"))
	mux.Handle("/debug/pprof/block", pprof.Handler("block"))
	mux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	mux.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	mux.Handle("/debug/pprof/mutex", pprof.Handler("mutex"))
	mux.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))

	server := &http.Server{
		Addr:    port,
		Handler: loggingMiddleware(logger, mux),
	}

	fmt.Printf("Server starting on port %s\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  GET  http://localhost:8080/")
	fmt.Println("  GET  http://localhost:8080/health")
	fmt.Println("  GET  http://localhost:8080/api/users")
	fmt.Println("  GET  http://localhost:8080/api/user?id=1")
	fmt.Println("  GET  http://localhost:8080/debug/pprof/")
	fmt.Println("  GET  http://localhost:8080/debug/pprof/heap")
	fmt.Println("  GET  http://localhost:8080/debug/pprof/goroutine")
	fmt.Println("  GET  http://localhost:8080/debug/pprof/profile?seconds=5")
	fmt.Println("Logs are written to execution-log-YYYY-MM-DD.txt at the project root.")

	logger.Printf("server started on %s", port)
	log.Fatal(server.ListenAndServe())
}
