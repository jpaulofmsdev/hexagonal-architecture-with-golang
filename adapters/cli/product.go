package cli

import (
	"fmt"

	"github.com/jpaulofmsdev/desafio-hexagonal-golang/application"
)

func Run(service application.IProductService, action string, productId string, productName string, productPrice float64) (string, error) {

	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been created with price %2.2f and status %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return "", err
		}
		product, err = service.Enable(product)
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf("Product %s has been enabled", product.GetName())

	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return "", err
		}
		product, err = service.Disable(product)
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf("Product %s has been disabled", product.GetName())

	default:
		product, err := service.Get(productId)
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %2.2f\nStatus: %s", product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	}

	return result, nil
}
