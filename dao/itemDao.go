package dao

import (
	"database/sql"
	"simple-inventory/models"
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

// CreateItem creates a new entry in the `items` table
func (dao *ItemDao) CreateItem(item models.Item) error {
	query := "INSERT INTO `items`(`name`, `description`) VALUES (?, ?)"
	_, err := dao.Db.Query(query, item.Name, item.Description)

	if err != nil {
		return err
	}

	return nil
}

// SetCategories assigns categories to an item
func (dao *ItemDao) SetCategories(itemID int, ids []int) {

}
