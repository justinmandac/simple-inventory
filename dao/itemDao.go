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
func (dao *ItemDao) GetItem(id int) (item models.Item, err error) {
	query := `SELECT * FROM items AS i
	RIGHT JOIN stocks AS s ON i.id = s.item_id
	WHERE i.id = ?`
	catQuery := `SELECT c.id, c.name, c.parentID
	FROM categories AS c RIGHT JOIN categories_map AS m
	ON c.id = m.category_id WHERE m.item_id = ?;`

	rows, err := dao.Db.Query(query, id)
	defer rows.Close()

	if err != nil {
		return item, err
	}

	for rows.Next() {
		err = rows.Scan(&item.ID, &item.Name, &item.Description, &item.Stock.ID, &item.Stock.ItemID, &item.Stock.Quantity)

		if err != nil {
			return item, err
		}
	}

	rows, err = dao.Db.Query(catQuery, id)

	if err != nil {
		return item, err
	}

	for rows.Next() {
		var cat models.ItemCategory

		err = rows.Scan(&cat.ID, &cat.Name, &cat.ParentID)

		if err != nil {
			return item, err
		}

		item.Categories = append(item.Categories, cat)
	}

	return item, nil
}

// SetCategories assigns categories to an item
func (dao *ItemDao) SetCategories(itemID int, categories []models.ItemCategory) []error {
	return dao.mapCategories(itemID, categories)
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

func (dao *ItemDao) mapItemStock(itemID int, stock models.ItemStock) []error {
	var errs []error
	query := "INSERT INTO `stocks`(`item_id`, `quantity`) VALUES (?, ?);"

	_, err := dao.Db.Exec(query, itemID, stock.Quantity)

	if err != nil {
		errs = append(errs, err)
		return errs
	}

	return nil
}

// SetItemStock updates an item's stock
func (dao *ItemDao) SetItemStock(stock models.ItemStock) []error {
	var errs []error
	query := "UPDATE `stocks` SET `quantity` = ? WHERE `id` = ? AND `item_id` = ?;"

	_, err := dao.Db.Exec(query, stock.Quantity, stock.ID, stock.ItemID)

	if err != nil {
		errs = append(errs, err)
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
	dao.mapCategories(int(id), item.Categories)
	dao.mapItemStock(int(id), item.Stock)
	// TODO : Set price per unit

	return nil
}
