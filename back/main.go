package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var db *sql.DB

func getSecret(path string) string {
    content, err := os.ReadFile(path)
    if err != nil {
        log.Fatalf("Failed to read secret at %s: %v", path, err)
    }
    return strings.TrimSpace(string(content))
}

func main() {
	user := getSecret("/run/secrets/postgres_user")
    pass := getSecret("/run/secrets/postgres_pass")
	dbName := os.Getenv("DATABASE_NAME")
    host := os.Getenv("DATABASE_HOST")
	log.Printf("DEBUG: The host read from env is: '%s'", host) 

	if host == "" {
		log.Fatal("DATABASE_HOST is empty! Check your Docker Compose env_file settings.")
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", user, pass, host, dbName)

	var err error
	for i := 0; i < 5; i++ {
		db, err = sql.Open("postgres", connStr)
		if err == nil {
			err = db.Ping() // Actually test the connection
		}

		if err == nil {
			log.Println("Successfully connected to the database!")
			break
		}

		log.Printf("DB not ready... retrying in 2s (Attempt %d/5): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Could not connect to DB after retries:", err)
	}

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/submit", submitHandler)
	http.HandleFunc("/users", listUsersHandler)

	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	ageStr := r.FormValue("age")

	// validation
	if name == "" || ageStr == "" {
		http.Error(w, "Missing fields", http.StatusBadRequest)
		return
	}

	age, err := strconv.Atoi(ageStr)
	if err != nil || age <= 0 {
		http.Error(w, "Invalid age", http.StatusBadRequest)
		return
	}

	_, err = db.Exec(
		"INSERT INTO users(name, age) VALUES($1, $2)",
		name,
		age,
	)

	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("User saved"))
}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User
		rows.Scan(&u.ID, &u.Name, &u.Age)
		users = append(users, u)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}