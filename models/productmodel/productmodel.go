package productmodel

import (
	"golang-native/config"
	"golang-native/entities"
)

func GetAll() []entities.Product {
	rows, err := config.DB.Query(`
		SELECT 
			products.id, 
			products.name, 
			categories.name as category_name,
			products.price, 
			products.descr, 
			products.stock, 
			products.created_at,
			products.updated_at FROM products
		JOIN categories ON products.category_id = categories.id
	`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var products []entities.Product

	for rows.Next() {
		var product entities.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Category.Name, &product.Price, &product.Desc, &product.Stock, &product.CreatedAt, &product.UpdatedAt); err != nil {
			panic(err)
		}

		products = append(products, product)
	}

	return products
}

func Create(product entities.Product) bool {
	rows, err := config.DB.Exec(
		`INSERT INTO products(
			name, category_id, price, stock, descr, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		product.Name,
		product.Category.Id,
		product.Price,
		product.Stock,
		product.Desc,
		product.CreatedAt,
		product.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	result, err := rows.LastInsertId()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func GetById(id int) entities.Product {
	row := config.DB.QueryRow(`
		SELECT 
			products.id, 
			products.name, 
			categories.name as category_name,
			products.price,
			products.descr,
			products.stock,
			products.created_at,
			products.updated_at FROM products 
		JOIN categories ON products.category_id = categories.id 
		WHERE products.id = ?
	`, id)

	var product entities.Product

	if err := row.Scan(&product.Id, &product.Name, &product.Category.Name, &product.Price, &product.Desc, &product.Stock, &product.CreatedAt, &product.UpdatedAt); err != nil {
		panic(err)
	}

	return product
}

func Update(id int, product entities.Product) bool {
	rows, err := config.DB.Exec(`
		UPDATE products set 
			name = ?, category_id = ?, price = ?, stock = ?, descr = ?, created_at = ?, updated_at = ? WHERE id = ?
	`, product.Name, product.Category.Id, product.Price, product.Stock, product.Desc, product.CreatedAt, product.UpdatedAt, id)

	if err != nil {
		panic(err)
	}

	result, err := rows.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}

func Delete(id int) bool {
	rows, err := config.DB.Exec(`
		DELETE FROM products WHERE id = ?
	`, id)

	if err != nil {
		panic(err)
	}

	result, err := rows.RowsAffected()
	if err != nil {
		panic(err)
	}

	return result > 0
}
