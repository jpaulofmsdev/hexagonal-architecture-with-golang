package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/jpaulofmsdev/desafio-hexagonal-golang/adapters/db"
	"github.com/jpaulofmsdev/desafio-hexagonal-golang/application"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProducts(Db)
}

func createTable(db *sql.DB) {
	table := "CREATE TABLE products (id STRING PRIMARY KEY, name STRING, price FLOAT, status STRING)"
	stmt, err := Db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func createProducts(db *sql.DB) {
	table := "INSERT INTO products (id, name, price, status) VALUES (?, ?, ?, ?)"
	stmt, err := db.Prepare(table)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = stmt.Exec("1", "Product 1", 10.00, "enabled")
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = stmt.Exec("2", "Product 2", 0.00, "disabled")
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDB(Db)
	product, err := productDb.Get("1")

	require.Nil(t, err)
	require.Equal(t, "1", product.GetID())
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 10.00, product.GetPrice())
	require.Equal(t, "enabled", product.GetStatus())

	product, err = productDb.Get("2")

	require.Nil(t, err)
	require.Equal(t, "2", product.GetID())
	require.Equal(t, "Product 2", product.GetName())
	require.Equal(t, 0.00, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDB(Db)

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25

	result, err := productDb.Save(product)
	require.Nil(t, err)
	require.NotEmpty(t, result.GetID())
	require.Equal(t, "Product Test", result.GetName())
	require.Equal(t, 25.0, result.GetPrice())
	require.Equal(t, "disabled", result.GetStatus())

	product.Enable()

	result2, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, "enabled", result2.GetStatus())
	require.Equal(t, result.GetID(), result2.GetID())

}
