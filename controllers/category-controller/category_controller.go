package categorycontroller

import (
	"golang-web-native/entities"
	categorymodel "golang-web-native/models/category-model"
	"log"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()
	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/category/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}

// add
func AddCategoryForm(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/category/add.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}

func Add(w http.ResponseWriter, r *http.Request) {
	log.Println("Masuk sini ==> ")
	name := r.FormValue("product")
	log.Println(r.Form)
	log.Println(name)

	category := entities.Category{Name: name}
	err := categorymodel.AddCategory(category)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}

func Edit(w http.ResponseWriter, r *http.Request) {

}
func Delete(w http.ResponseWriter, r *http.Request) {

}
