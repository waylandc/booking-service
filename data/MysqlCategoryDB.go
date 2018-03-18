package data

import (
	"database/sql"
	"wchan/booking-service/models"
)

// MysqlCategoryDB is our interface for the Category table
type MysqlCategoryDB struct {
	Conn *sql.DB
}

// Create a Category record
func (db *MysqlCategoryDB) Create(c models.Category) error {
	stmt, err := db.Conn.Prepare("insert into category(name) values (?)")
	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(c.Name)
	return err
}

// Delete a Category from database
func (db *MysqlCategoryDB) Delete(c models.Category) error {
	stmt, err := db.Conn.Prepare("delete from category where name=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(c.Name)
	return err
}

// Update a Category
func (db *MysqlCategoryDB) Update(c models.Category) error {
	stmt, err := db.Conn.Prepare("update category set name=? where ID=?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(c.Name, c.ID)
	return err
}

// Find categories matching name
func (db *MysqlCategoryDB) Find(name string) ([]models.Category, error) {
	rs := []models.Category{}
	rows, err := db.Conn.Query("select id, name from category where name like ?", name)
	if err != nil {
		return rs, err
	}
	defer rows.Close()

	for rows.Next() {
		var cat models.Category
		err := rows.Scan(&cat.ID, &cat.Name)
		if err != nil {
			return rs, err
		}
		rs = append(rs, cat)
	}

	return rs, nil
}
