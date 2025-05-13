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
func CreateCategoryForm(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/category/create.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var category entities.Category
	category.Name = r.FormValue("product")
	category.CreatedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	err := categorymodel.CreateCategory(category)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}

//===============================end of add==========================

// ===============================edit===========================
func UpdateCategoryForm(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	category, err := categorymodel.FindCategoryById(id)
	if err != nil {
		panic(err)
	}

	temp, err := template.ParseFiles("views/category/update.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, category)
}

func Update(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("product")
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	var category entities.Category
	category.Id = uint(id)
	category.Name = name
	category.UpdatedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	ok := categorymodel.UpdateCategory(category)

	// next bakal begini semua
	if !ok {
		http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/categories", http.StatusSeeOther)
}

// =======================end of edit ====================

func Delete(w http.ResponseWriter, r *http.Request) {

}
