package main

import (
	"golang-web-native/config"
	categorycontroller "golang-web-native/controllers/category-controller"
	homecontroller "golang-web-native/controllers/home-controller"
	"log"
	"net/http"
)

func main() {
	log.Println("Connect..")
	config.ConnectDB()

	// panggil home page

	http.HandleFunc("/", homecontroller.Welcome)

	// panggil categories page
	http.HandleFunc("/categories", categorycontroller.Index)
	http.HandleFunc("/categories/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			categorycontroller.AddCategoryForm(w, r)
		} else if r.Method == http.MethodPost {
			categorycontroller.Add(w, r)
		}
	})
	// http.HandleFunc("/categories", categorycontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
