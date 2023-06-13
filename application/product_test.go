package application_test

import (
	"testing"

	"github.com/jpaulofmsdev/desafio-hexagonal-golang/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "price must be greater than 0 to enable product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10

	err = product.Disable()
	require.Equal(t, "price must be 0 to disable product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Product 1"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "invalid"

	_, err = product.IsValid()
	require.Equal(t, "status must be either disabled or enabled", err.Error())

	product.Status = application.ENABLED

	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10

	_, err = product.IsValid()
	require.Equal(t, "price must be greater than or equal 0", err.Error())

	product.ID = "123"
	product.Price = 10

	_, err = product.IsValid()
	require.Equal(t, "ID: 123 does not validate as uuidv4", err.Error())
}

func TestProduct_GetID(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.NewV4().String()

	id := product.GetID()

	require.Equal(t, id, product.ID)
}

func TestProduct_GetName(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"

	name := product.GetName()

	require.Equal(t, name, product.Name)
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}
	product.Status = application.DISABLED

	status := product.GetStatus()

	require.Equal(t, status, product.Status)
}

func TestProduct_GetPrice(t *testing.T) {
	product := application.Product{}
	product.Price = 10

	price := product.GetPrice()

	require.Equal(t, price, product.Price)
}

func TestProduct_NewProduct(t *testing.T) {
	product := application.NewProduct()

	require.NotNil(t, product)
	require.NotNil(t, product.ID)
	require.Equal(t, application.DISABLED, product.Status)
}
