package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	mw "restapi/internal/api/middlewares"
	"restapi/internal/api/utils"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Root endpoint accessed")
	fmt.Fprint(w, "Hello, World!")
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Students endpoint accessed")
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Hello GET method in students routes"))
	case http.MethodPost:
		w.Write([]byte("Hello POST method in students routes"))
	case http.MethodPut:
		w.Write([]byte("Hello PUT method in students routes"))
	case http.MethodPatch:
		w.Write([]byte("Hello PATCH method in students routes"))
	case http.MethodDelete:
		w.Write([]byte("Hello DELETE method in students routes"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Teachers endpoint accessed")
	fmt.Fprint(w, "Teachers Endpoint")
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Executives endpoint accessed")
	fmt.Fprint(w, "Executives Endpoint")
}

func main() {
	port := ":3000"

	cert := "cert.pem"
	key := "key.pem"

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	mux.HandleFunc("/students/", studentsHandler)

	mux.HandleFunc("/teachers/", teachersHandler)

	mux.HandleFunc("/execs/", execsHandler)

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	secureMux := utils.applyMiddlewares(mux,
		mw.SecurityHeaders,
		mw.Cors,
		mw.ResponseTimeMiddleware,
	)

	server := &http.Server{
		Addr:      port,
		Handler:   secureMux,
		TLSConfig: tlsConfig,
	}

	fmt.Printf("Server is running on port %s\n", port)
	if err := server.ListenAndServeTLS(cert, key); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
