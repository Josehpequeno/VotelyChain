// main.go
package main

import (
    "database/sql"
    "log"
    "net/http"
    "votelychain/controllers"
    "os"

    _ "github.com/mattn/go-sqlite3"
    "github.com/joho/godotenv"
)

var db *sql.DB



func initializeDatabase(db *sql.DB) {
    // Lê o conteúdo do arquivo SQL
    sqlFile, err := os.ReadFile("votely.sql")
    if err != nil {
        log.Fatalf("Error reading SQL file: %v", err)
    }

    // Executa o script SQL
    _, err = db.Exec(string(sqlFile))
    if err != nil {
        log.Fatalf("Error executing SQL script: %v", err)
    }
}

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    masterUsername := os.Getenv("MASTER_USERNAME")
    masterPassword := os.Getenv("MASTER_PASSWORD")

    log.Printf("Master Username: %s", masterUsername)
    log.Printf("Master Password: %s", masterPassword)

    db, err = sql.Open("sqlite3", "./votely.db")
    if err != nil {
      log.Fatal("Failed to connect to the database:", err)
    }
    log.Println("Database connection established") 
    defer db.Close()

    initializeDatabase(db)

    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT UNIQUE,
        password TEXT
    )`)
    if err != nil {
        log.Fatal(err)
    }

    var count int
    err = db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", masterUsername).Scan(&count)
    if err != nil {
        log.Fatal(err)
    }

    if count == 0 {
        _, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", masterUsername, masterPassword)
        if err != nil {
            log.Fatal(err)
        }
        log.Println("Master user created")
    } else {
        log.Println("Master user already exists")
    }


    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))


    http.HandleFunc("/", controllers.LoginView)
    http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) { controllers.Login(db, w, r) })
    http.HandleFunc("/elections", func(w http.ResponseWriter, r *http.Request) { controllers.ElectionView(db, w, r) })
    http.HandleFunc("/elections/register", func(w http.ResponseWriter, r *http.Request) { controllers.RegisterElection(db, w, r) })
    
    log.Println("Server started at http://localhost:8080")
    log.Println("Press Ctrl+C to stop the server.")
    http.ListenAndServe("0.0.0.0:8080", nil)
}

