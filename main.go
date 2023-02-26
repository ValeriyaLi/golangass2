package main

import (
	"ass1/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"html/template"
	"net/http"
)

func initDB() (*gorm.DB, error) {
	info := "host=localhost user=postgres password=December1225 dbname=assignment2_http port=5432"
	return gorm.Open(postgres.Open(info), &gorm.Config{})
}

func myStore(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("templates/index.html")
	tmp.Execute(w, nil)
}

func handleRegistration(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db, _ := initDB()
		name := r.FormValue("name")
		username := r.FormValue("username")
		password := r.FormValue("password")
		hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		u := &models.User{Name: name, Login: username, Password: string(hash)}
		db.Create(&u)
		db.AutoMigrate(&models.User{})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	tmp, _ := template.ParseFiles("templates/registration.html")
	tmp.Execute(w, nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db, _ := initDB()
		username := r.FormValue("username")
		password := r.FormValue("password")
		var user models.User
		res := db.Where("login = ?", username).First(&user)
		if res.Error != nil {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		}
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Println("success")
		return
	}
	tmp, _ := template.ParseFiles("templates/login.html")
	tmp.Execute(w, nil)
}

func main() {
	//err := bcrypt.CompareHashAndPassword([]byte(), []byte("fkjjkfs"))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", myStore)
	http.HandleFunc("/registration", handleRegistration)
	http.HandleFunc("/login", handleLogin)

	fmt.Println("server on http//localhost:8000")
	http.ListenAndServe(":8000", nil)

	//info := "host=localhost user=postgres password=December1225 dbname=assignment2_http port=5432"
	//db, err := gorm.Open(postgres.Open(info), &gorm.Config{})
	//db.AutoMigrate(&models.Item{})
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
}
