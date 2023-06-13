package db

import (
	"database/sql"

	"github.com/jpaulofmsdev/desafio-hexagonal-golang/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDB struct {
	db *sql.DB
}

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{db: db}
}
func (p *ProductDB) Get(id string) (application.IProduct, error) {
	var product application.Product
	stmt, err := p.db.Prepare("SELECT id, name, price, status FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDB) Save(product application.IProduct) (application.IProduct, error) {
	if product.GetID() != "" {
		var rows int
		p.db.QueryRow("SELECT COUNT(*) FROM products WHERE id = ?", product.GetID()).Scan(&rows)
		if rows > 0 {
			_, err := p.update(product)
			if err != nil {
				return nil, err
			}
			return product, nil
		}
	}

	_, err := p.create(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDB) create(product application.IProduct) (application.IProduct, error) {
	stmt, err := p.db.Prepare("INSERT INTO products (id, name, price, status) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDB) update(product application.IProduct) (application.IProduct, error) {
	stmt, err := p.db.Prepare("UPDATE products SET name = ?, price = ?, status = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	return product, nil
}
