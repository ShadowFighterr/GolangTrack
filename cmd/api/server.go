package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func teachersHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Teachers endpoint"))
	fmt.Fprintln(w, "Teachers endpoint")
	switch r.Method {

	case http.MethodGet:
		fmt.Println(r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path, "/teachers/")
		userID := strings.TrimPrefix(path, "/")
		fmt.Println("User ID:", userID)
		w.Write([]byte("Get request"))
		fmt.Println("GET request received")

	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}
		response := make(map[string]interface{})
		for k, v := range r.Form {
			response[k] = v[0]
		}
		fmt.Println("Form Data:", r.Form)
		fmt.Println("Parsed Form Data:", response)

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		fmt.Println("Request Body:", string(body))
		defer r.Body.Close()

		w.Write([]byte("Post request"))
		fmt.Println("POST request received")

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func studentsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Students endpoint"))
	// fmt.Fprintln(w, "Students endpoint")
}

func execsHandler(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Executives endpoint"))
	fmt.Fprintln(w, "Executives endpoint")
}

func main() {
	port := ":3000"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
		fmt.Println("Hello, World!")
	})

	http.HandleFunc("/teachers/", teachersHandler)

	http.HandleFunc("/students", studentsHandler)

	http.HandleFunc("/execs", execsHandler)

	fmt.Println("Starting server on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
