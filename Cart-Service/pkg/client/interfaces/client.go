package interfaceClient

type NewProductClient interface {
	CheckProduct(product_id int) (bool, error)
	GetQuantityFromProductID(id int) (int, error)
	GetPriceOfProductFromID(id int) (float64, error)
	ProductStockMinus(productID, stock int) error
}