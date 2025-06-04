package handlers

import (
	"app/db"
	"html/template"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, _ := template.ParseFiles("templates/signup.html")
		tmpl.Execute(w, nil)
		return
	}

	// POST
	username := r.FormValue("username")
	password := r.FormValue("password")
	role := r.FormValue("role")

	_, err := db.DB.Exec("INSERT INTO users (username, password, role) VALUES ($1, $2, $3)",
		username, password, role)
	if err != nil {
		http.Error(w, "Database error", 500)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
