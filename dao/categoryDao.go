package dao

import (
	"database/sql"
	"simple-inventory/models"
)

// CategoryDao  Data Access Object for Categories
type CategoryDao struct {
	Db *sql.DB
}

// GetCategories gets the list of all categories
func (dao *CategoryDao) GetCategories() (categories []models.ItemCategory, err error) {
	rows, err := dao.Db.Query("SELECT * FROM categories;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var read models.ItemCategory
		err = rows.Scan(&read.ID, &read.Name, &read.ParentID)

		if err != nil {
			return nil, err
		}

		categories = append(categories, read)
	}

	return
}

// CreateCategory Creates a new category
func (dao *CategoryDao) CreateCategory(category models.ItemCategory) error {
	query := "INSERT INTO `categories`(`name`, `parentID`) VALUES (?, ?)"
	_, err := dao.Db.Exec(query, category.Name, category.ParentID)

	if err != nil {
		return err
	}

	return nil
}

// DeleteCategory deletes  a category with the specified ID
func (dao *CategoryDao) DeleteCategory(id int) error {
	query := "DELETE FROM `categories` WHERE id=?"
	_, err := dao.Db.Query(query, id)

	if err != nil {
		return err
	}

	return nil
}

// UpdateCategory updates a category with the specified Id and category model
func (dao *CategoryDao) UpdateCategory(id int, category models.ItemCategory) error {
	// TODO: Throw error if category.ID != id
	query := "UPDATE `categories`SET `parentID`=? WHERE `id`=?"
	_, err := dao.Db.Query(query, category.ParentID, id)

	if err != nil {
		return err
	}

	return nil
}

// GetCategory gets a specific category
func (dao *CategoryDao) GetCategory(id int) (cat models.ItemCategory, err error) {
	query := "SELECT * FROM `categories` WHERE id=?"
	rows, err := dao.Db.Query(query, id)
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&cat.ID, &cat.Name, &cat.ParentID)
	}

	return cat, nil
}
