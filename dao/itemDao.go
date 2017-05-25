// Package dao provides Data Access Objects for abstracting accesses to the
// Data layers
// TODO : Create method for removing item categories

package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"simple-inventory/models"
	"strconv"
)

// ItemDao Data Access Object for items
type ItemDao struct {
	Db *sql.DB
}

// GetItems retrieves a list of all items
func (dao *ItemDao) GetItems() (items []models.Item, err error) {
	query := "SELECT * from `items`"
	rows, err := dao.Db.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var item models.Item
		err = rows.Scan(&item.Name, &item.ID, &item.Description)

		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

// GetItem retrieves a specific item
func (dao *ItemDao) GetItem(id int) (item *models.Item, err error) {
	query := "SELECT * from `items` WHERE id=?"
	rows, err := dao.Db.Query(query, id)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&item.Name, &item.ID, &item.Description)

		if err != nil {
			return nil, err
		}
	}

	return item, nil
}

// SetCategories assigns categories to an item
func (dao *ItemDao) SetCategories(itemID int, categories []models.ItemCategory) []error {

	errs := dao.mapCategories(itemID, categories)

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (dao *ItemDao) mapCategories(itemID int, categories []models.ItemCategory) []error {
	var errs []error
	catLength := len(categories)
	query := "INSERT INTO `categories_map`(`item_id`, `category_id`) VALUES (? , ?);"
	// Query for checking if the category to be inserted exists
	chckQuery := "SELECT id FROM `categories` WHERE `id`= ?;"

	if catLength > 0 {
		for i := 0; i < catLength; i++ {
			currCat := categories[i]
			// Check if category exists
			rows, err := dao.Db.Query(chckQuery, currCat.ID)

			if err != nil {
				errs = append(errs, err)
			}

			if !rows.Next() {
				conv := strconv.FormatInt(int64(currCat.ID), 10)
				errs = append(errs, errors.New("Category does not exist : "+conv))
			} else {
				// TODO: Perform checks if the itemID-categoryID pair already exists
				_, err = dao.Db.Exec(query, itemID, currCat.ID)

				if err != nil {
					// Fatally exit
					errs = append(errs, err)
					return errs
				}
			}

		}
	} else {
		errs = append(errs, errors.New("Empty data provided"))
		return errs
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

// CreateItem creates a new entry in the `items` table
func (dao *ItemDao) CreateItem(item models.Item) error {
	query := "INSERT INTO `items`(`name`, `description`) VALUES (?, ?)"
	row, err := dao.Db.Exec(query, item.Name, item.Description)

	if err != nil {
		return err
	}

	id, err := row.LastInsertId()
	fmt.Println("Inserted item id: ", id)

	// Map to categories if item.Categories contains elements
	dao.mapCategories(item.ID, item.Categories)
	// TODO : Set initial stock
	// TODO : Set price per unit

	return nil
}
