package categorymodel

import (
	"context"
	"golang-web-native/config"
	"golang-web-native/entities"
)

func GetAll() []entities.Category {
	ctx := context.Background()
	rows, err := config.DB.QueryContext(ctx, "SELECT * FROM categories")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}

	return categories
}

// add
func CreateCategory(category entities.Category) bool {
	ctx := context.Background()
	script := "INSERT INTO categories(name, created_at) VALUES(?, ?)"
	query, err := config.DB.ExecContext(ctx, script, category.Name, category.CreatedAt.Time)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()

	if err != nil {
		panic(err)
	}

	return result > 0
}

// edit
func FindCategoryById(id int) (entities.Category, error) {
	row := config.DB.QueryRow("SELECT id, name FROM categories WHERE id=?", id)
	var category entities.Category
	err := row.Scan(&category.Id, &category.Name)
	return category, err
}

func UpdateCategory(category entities.Category) bool {
	ctx := context.Background()
	script := "UPDATE categories SET name = ?, updated_at = ? WHERE id = ?"
	query, err := config.DB.ExecContext(ctx, script, category.Name, category.UpdatedAt.Time, category.Id)
	if err != nil {
		panic(err)
	}

	// next bakal begini semua
	result, err := query.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

// delete

func DeleteCategory(id int) bool {
	ctx := context.Background()
	script := "DELETE FROM categories WHERE id = ?"
	query, err := config.DB.ExecContext(ctx, script, id)
	if err != nil {
		panic(err)
	}

	result, err := query.RowsAffected()

	if err != nil {
		panic(err)
	}

	return result > 0

}
