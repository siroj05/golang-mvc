package categorycontroller

import (
	"database/sql"
	"golang-web-native/entities"
	categorymodel "golang-web-native/models/category-model"
	"net/http"
	"strconv"
	"text/template"
	"time"
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

// ================================add============================
func AddCategoryForm(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/category/add.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}

func Add(w http.ResponseWriter, r *http.Request) {
	var category entities.Category
	category.Name = r.FormValue("product")
	category.CreatedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	err := categorymodel.AddCategory(category)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}

//===============================end of add==========================

// ===============================edit===========================
func EditCategoryForm(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	category, err := categorymodel.FindCategoryById(id)
	if err != nil {
		panic(err)
	}

	temp, err := template.ParseFiles("views/category/edit.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, category)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("product")
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	var category entities.Category
	category.Id = uint(id)
	category.Name = name
	category.UpdatedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	err := categorymodel.EditCategory(category)
	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}

// =======================end of edit ====================

func Delete(w http.ResponseWriter, r *http.Request) {

}
