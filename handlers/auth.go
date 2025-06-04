package handlers

import (
	"app/db"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, _ := template.ParseFiles("templates/login.html")
		tmpl.Execute(w, nil)
		return
	}

	// POST
	username := r.FormValue("username")
	password := r.FormValue("password")

	var role string
	err := db.DB.QueryRow("SELECT role FROM users WHERE username=$1 AND password=$2", username, password).Scan(&role)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// تحويل حسب الدور
	switch role {
	case "admin":
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	case "staff":
		http.Redirect(w, r, "/staff", http.StatusSeeOther)
	case "member":
		http.Redirect(w, r, "/member", http.StatusSeeOther)
	default:
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
