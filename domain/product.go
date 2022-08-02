package domain

import "github.com/labstack/echo/v4"

type Product struct {
	ID          int
	ProductName string
	Description string
	Price       int
	ProductPic  string
	Stock       int
	Qty         int
}

type ProductHandler interface {
	// GetItems() echo.HandlerFunc
	// PostCart() echo.HandlerFunc
	PostItem() echo.HandlerFunc
	UpdateItem() echo.HandlerFunc
	// DeleteItem() echo.HandlerFunc
	// GetItem() echo.HandlerFunc
}

type ProductUseCase interface {
	GetAllItems()
	PostCartUser()
	PostItemAdmin(newProduct Product) int
	UpdateItemAdmin(newProduct Product, productID int) int
	DeleteItemAdmin()
	GetItemUser()
}

type ProductData interface {
	GetAllItemData()
	PostCartData()
	PostItemData(newProduct Product) Product
	UpdateItemData(newProduct Product, productID int) Product
	DeleteItemData()
	GetItemData()
}
