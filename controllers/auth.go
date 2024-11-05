// controllers/auth.go
package controllers

import (
    "net/http"
    "votelychain/models"
    "html/template"
    "database/sql"
    "log"
    "os"
)

func LoginView(w http.ResponseWriter, r *http.Request) {
   log.Println("LoginView called")
    
    // Verifica o diretório atual
    currentDir, err := os.Getwd()
    if err != nil {
        log.Println("Error getting current directory:", err)
    }
    log.Println("Current directory:", currentDir)

    tmpl := template.Must(template.ParseFiles("templates/layout.html", "templates/login.html"))

    if err != nil {
        log.Println("Error loading template:", err)
        http.Error(w, "Error loading template", http.StatusInternalServerError)
        return
    }

    err = tmpl.Execute(w, nil)
    if err != nil {
        log.Println("Error rendering template:", err)
        http.Error(w, "Error rendering template", http.StatusInternalServerError)
        return
    } 
}

func Login(db *sql.DB, w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        username := r.FormValue("username")
        password := r.FormValue("password")
        
        _, err := models.Authenticate(db, username, password)
        if err != nil {
            http.Error(w, "Invalid username or password.", http.StatusUnauthorized)
            return
        }
        
        http.Redirect(w, r, "/elections", http.StatusSeeOther)
    }
}

