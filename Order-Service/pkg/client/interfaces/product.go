package interfaceClient

type ProductClient interface {
	ProductStockMinus(productID, stock int) error
}