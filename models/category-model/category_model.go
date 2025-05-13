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

func AddCategory(category entities.Category) error {
	ctx := context.Background()
	script := "INSERT INTO categories(name) VALUES(?)"
	_, err := config.DB.ExecContext(ctx, script, category.Name)
	if err != nil {
		return err
	}
	return nil
}
