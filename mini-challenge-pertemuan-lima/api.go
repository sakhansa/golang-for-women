package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type Account struct {
	Email     string
	Alamat    string
	Pekerjaan string
	Alasan    string
	Nama      string
}

var accounts = []Account{
	{Email: "Khansa@mail.com", Alamat: "Jl. RE Martadinata", Pekerjaan: "Accountant", Alasan: "Menambah wawasan bahasa pemrograman"},
	{Email: "Erri@mail.com", Alamat: "Jl. Fatahilah", Pekerjaan: "BE Developer", Alasan: "Menunjang pekerjaan"},
	{Email: "Rizka@mail.com", Alamat: "Jl. Asia Afrika", Pekerjaan: "Android Developer", Alasan: "Mengisi waktu dengan kegiatan bermanfaat"},
	{Email: "Asrie@mail.com", Alamat: "Jl. Ahmad Yani", Pekerjaan: "FE Developer", Alasan: "Belajar bahasa pemrograman baru"},
	{Email: "Hazrina@mail.com", Alamat: "Jl. Stasiun Wonokromo", Pekerjaan: "IOS Developer", Alasan: "Untuk mencari freelance"},
}

var PORT = ":9090"

func main() {
	http.HandleFunc("/", login)
	http.HandleFunc("/login", getDetailAccount)

	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	// using html/template package
	if r.Method == "GET" {
		// tpl := template.Must(template.New("form").ParseFiles("loginPage.html"))
		// err := tpl.Execute(w, nil)
		tpl, err := template.ParseFiles("loginPage.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, accounts)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}

func getDetailAccount(w http.ResponseWriter, r *http.Request) {
	// using html/template package
	if r.Method == "POST" {
		tpl, err := template.ParseFiles("detailAccount.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		email := r.FormValue("email")
		var account Account

		for _, a := range accounts {
			if a.Email == email {
				nama := strings.SplitN(a.Email, "@", -1)
				a.Nama = nama[0]
				account = a
				break
			}
		}

		tpl.Execute(w, account)
		return
	}

	http.Error(w, "Invalid method", http.StatusBadRequest)
}
