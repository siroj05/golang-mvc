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

	// add categories page
	http.HandleFunc("/categories/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			categorycontroller.CreateCategoryForm(w, r)
		} else if r.Method == http.MethodPost {
			categorycontroller.Create(w, r)
		}
	})

	// edit categories page
	http.HandleFunc("/categories/edit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			categorycontroller.UpdateCategoryForm(w, r)
		} else if r.Method == http.MethodPost {
			categorycontroller.Update(w, r)
		}
	})

	// delete categories page
	http.HandleFunc("/categories/delete", func(w http.ResponseWriter, r *http.Request) {
		categorycontroller.Delete(w, r)
	})
	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
