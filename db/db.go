package db

import (
	"database/sql"

	"github.com/diyor200/ecommerce/model"
	_ "github.com/xuri/excelize/v2"
)

func createTables(db *sql.DB) error {
	// category jadvali yaratilmoqda
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS category (
		id SERIAL PRIMARY KEY,
		name VARCHAR(200) NOT NULL
	)`)
	if err != nil {
		return err
	}

	// products jadvali yaratilmoqda
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		category INTEGER NOT NULL,
		name VARCHAR(200) NOT NULL,
		price FLOAT NOT NULL
		FOREIGN KEY(category) REFERENCES
	)`)
	if err != nil {
		return err
	}

	// users jadvali yaratilmoqda
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		firstname VARCHAR(200) NOT NULL,
		lastname VARCHAR(200) NOT NULL
	)`)
	if err != nil {
		return err
	}

	// orders jadvali yaratilmoqda
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS orders(
		id SERIAL PRIMARY KEY,
		userid INTEGER NOT NULL,
		productid INTEGER NOT NULL
	);`)
	if err != nil {
		return err
	}

	return nil
}

func AddCategory(name string, db *sql.DB) error {
	_, err := db.Exec(`insert into category (name) VALUES($1)`, name)
	if err != nil {
		return err
	}
	return nil
}

func AddProduct(product model.Product, db *sql.DB) error {
	_, err := db.Exec(`insert into products(name, category, price) values($1, $2, $3)`, product.Name, product.Category, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func AddOrders(order model.Orders, db *sql.DB) error {
	_, err := db.Exec(`INSERT INTO orders(userid, productid) values($1, $2)`, order.User.ID, order.Products.ID)
	if err != nil {
		return err
	}
	return nil
}

func AddUser(db *sql.DB) error {
	var user model.User
	_, err := db.Exec(`insert into users(firstname, lastname) values($1, $2)`, user.Firstname, user.Lastname)
	if err != nil {
		return err
	}
	return nil
}

func GetCategory(db *sql.DB) ([]model.Category, error) {
	rows, err := db.Query(`SELECT * FROM category`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var categories []model.Category
	for rows.Next() {
		category := model.Category{}
		err := rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil

}

func GetPruduct(id int, db *sql.DB) (model.Product, error) {
	row, err := db.Query(`SELECT * FROM products WHERE id=$1 LIMIT 1`, id)
	product := model.Product{}
	if err != nil {
		return product, err
	}

	defer row.Close()

	err = row.Scan(&product.ID, &product.Category, &product.Name, &product.Price)
	if err != nil {
		return product, err
	}

	return product, nil

}

func GetPruducts(db *sql.DB) ([]model.Product, error) {
	rows, err := db.Query(`select * from products`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []model.Product

	for rows.Next() {
		product := model.Product{}
		err := rows.Scan(&product.ID, &product.Category, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func GetUsers(db *sql.DB) ([]model.User, error) {
	rows, err := db.Query(`SELECT * FROM users`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []model.User

	for rows.Next() {
		user := model.User{}

		err := rows.Scan(&user.ID, &user.Firstname, &user.Lastname)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func GetOrders(db *sql.DB) error {
	rows, err := db.Query(`SELECT * FROM orders`)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {

	}
	return nil
}