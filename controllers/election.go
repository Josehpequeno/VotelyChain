// controllers/election.go
package controllers

import (
    "net/http"
    "html/template"
    "database/sql"
    "votelychain/models"
    "log"
)


func ElectionView(db *sql.DB, w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT id, title, description FROM elections")
    if err != nil {
        log.Printf("Error fetching elections: %v", err) // Log the error
        http.Error(w, "Error fetching elections. Please try again.", http.StatusInternalServerError)
        return
    }
    defer rows.Close()
    
    var elections []models.Election
    for rows.Next() {
        var election models.Election
        if err := rows.Scan(&election.ID, &election.Title, &election.Description); err != nil {
            log.Printf("Error processing election data: %v", err) // Log the error
            http.Error(w, "Error processing election data. Please try again.", http.StatusInternalServerError)
            return
        }
        elections = append(elections, election)
    }

    if err = rows.Err(); err != nil {
        log.Printf("Error during row iteration: %v", err) // Log the error
        http.Error(w, "Error during row iteration. Please try again.", http.StatusInternalServerError)
        return
    }

    templates := template.Must(template.ParseFiles("templates/layout.html", "templates/elections.html")) 
    if err := templates.Execute(w, elections); err != nil {
        log.Printf("Error rendering template: %v", err) // Log the error
        http.Error(w, "Error rendering template. Please try again.", http.StatusInternalServerError)
    }
}




func RegisterElection(db *sql.DB, w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        title := r.FormValue("title")
        description := r.FormValue("description")
        
        _, err := db.Exec("INSERT INTO elections (title, description) VALUES (?, ?)", title, description)
        if err != nil {
            http.Error(w, "Error creating election. Please try again.", http.StatusInternalServerError)
            return
        }
        
        http.Redirect(w, r, "/elections", http.StatusSeeOther)
    }
}

