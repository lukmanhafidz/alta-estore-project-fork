package usecase

import (
	"altaproject2/domain"
	"altaproject2/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateProduct(t *testing.T) {
	repo := new(mocks.CartData)
	mockData := domain.Cart{ID: 1, Userid: 1, Subtotal: 50000, Quantity: 2, Productid: 1}

	returnData := mockData
	// returnData.ID = 1

	t.Run("Success Update", func(t *testing.T) {
		repo.On("UpdateData", returnData, 1).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res := useCase.UpdateData(mockData, 1)

		assert.Equal(t, 200, res)
	})
}

func TestDeleteCart(t *testing.T) {
	repo := new(mocks.CartData)

	t.Run("Succes delete", func(t *testing.T) {
		repo.On("DeleteData", mock.Anything).Return(true).Once()
		cartcase := New(repo, validator.New())
		delete, err := cartcase.DeleteCart(1)

		assert.Nil(t, err)
		assert.Equal(t, true, delete)
		repo.AssertExpectations(t)
	})
}

func TestPostCart(t *testing.T) {
	repo := new(mocks.CartData)

	mockData := domain.Cart{Quantity: 2, Productid: 1}
	returnData := domain.Cart{ID: 1, Subtotal: 200000, Userid: 1, Productid: 1, Quantity: 2}

	data := mockData
	data.Quantity = 0

	invalidData := returnData
	invalidData.ID = 0

	t.Run("Success post cart", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("PostData", mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		status := useCase.PostCart(mockData)

		assert.Equal(t, 200, status)
		repo.AssertExpectations(t)
	})

	t.Run("Validator error", func(t *testing.T) {
		useCase := New(repo, validator.New())
		status := useCase.PostCart(data)

		assert.Equal(t, 400, status)
		repo.AssertExpectations(t)
	})

	t.Run("Duplicate data", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(true).Once()
		useCase := New(repo, validator.New())
		status := useCase.PostCart(mockData)

		assert.Equal(t, 400, status)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("CheckDuplicate", mock.Anything).Return(false).Once()
		repo.On("PostData", mock.Anything).Return(invalidData).Once()
		useCase := New(repo, validator.New())
		status := useCase.PostCart(mockData)

		assert.Equal(t, 404, status)
		repo.AssertExpectations(t)
	})
}

func TestGetCart(t *testing.T) {
	repo := new(mocks.CartData)

	returnProduct := []domain.CartProduct{{ProductName: "men bag", Description: "this is a men bag", Price: 100000, ProductPic: "bag.jpg", Stock: 5}}
	returnCart := domain.Cart{ID: 1, Subtotal: 200000, Quantity: 2, Userid: 1}

	invalidData := domain.Cart{ID: 0, Subtotal: 0, Quantity: 0, Userid: 0}

	t.Run("succsess get data", func(t *testing.T) {
		repo.On("GetDataProduct", mock.Anything).Return(returnProduct).Once()
		repo.On("GetData", mock.Anything).Return(returnCart).Once()
		useCase := New(repo, validator.New())
		cart, product, status := useCase.GetCart(1)

		assert.Equal(t, returnProduct, product)
		assert.Equal(t, returnCart, cart)
		assert.Equal(t, 200, status)
		repo.AssertExpectations(t)
	})

	t.Run("data not found", func(t *testing.T) {
		repo.On("GetDataProduct", mock.Anything).Return(returnProduct).Once()
		repo.On("GetData", mock.Anything).Return(invalidData).Once()
		useCase := New(repo, validator.New())
		cart, product, status := useCase.GetCart(1)

		assert.Equal(t, []domain.CartProduct(nil), product)
		assert.Equal(t, invalidData, cart)
		assert.Equal(t, 404, status)
		repo.AssertExpectations(t)
	})
}
