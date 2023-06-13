package cli_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jpaulofmsdev/desafio-hexagonal-golang/adapters/cli"
	mock_application "github.com/jpaulofmsdev/desafio-hexagonal-golang/application/mocks"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Product Test"
	productPrice := 21.22
	productStatus := "enabled"
	productId := "abc"

	productMock := mock_application.NewMockIProduct(ctrl)

	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	productServiceMock := mock_application.NewMockIProductService(ctrl)
	productServiceMock.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	productServiceMock.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultExpected := "Product ID abc with the name Product Test has been created with price 21.22 and status enabled"

	result, err := cli.Run(productServiceMock, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = "Product Product Test has been enabled"

	result, err = cli.Run(productServiceMock, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = "Product Product Test has been disabled"

	result, err = cli.Run(productServiceMock, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = "Product ID: abc\nName: Product Test\nPrice: 21.22\nStatus: enabled"

	result, err = cli.Run(productServiceMock, "", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}
