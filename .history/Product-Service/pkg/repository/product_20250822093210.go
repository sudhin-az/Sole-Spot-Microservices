package repository

import (
	"errors"

	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Product-Service/pkg/domain"
	"github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Product-Service/pkg/models"
	interfaceRepository "github.com/sudhin-az/SOLE-SPOT-MICROSERVICES/Product-Service/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaceRepository.ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (pr *ProductRepository) ProductAlreadyExist(Name string) bool {
	var count int
	err := pr.DB.Raw("SELECT count(*) FROM products WHERE name = ?", Name).Scan(&count).Error
	if err != nil {
		return false
	}
	return count > 0
}

func (pr *ProductRepository) StockInvalid(Name string) bool {
	var count int
	if err := pr.DB.Raw("SELECT count(stock) FROM products WHERE name = ? AND stock >= 0", Name).Scan(&count).Error; err != nil {
		return false
	}
	return count > 0
}

func (pr *ProductRepository) AddProduct(product models.Product) (domain.Product, error) {
	var p models.Product
	query := `
	INSERT INTO products(name, description, category_id, size, stock, price)
	VALUES($1, $2, $3, $4, $5, $6)
	RETURNING name, description, category_id, size, stock, price`
	err := pr.DB.Raw(query, product.Name, product.Description, product.CategoryID, product.Size, product.Stock, product.Price).Scan(&p).Error
	if err != nil {
		return domain.Product{}, err
	}
	var productResponses domain.Product
	err = pr.DB.Raw("SELECT * FROM products WHERE name = ?", p.Name).Scan(&productResponses).Error
	if err != nil {
		return domain.Product{}, err
	}
	return productResponses, nil
}

func (pr *ProductRepository) ListProducts(userID int) ([]models.ProductBrief, error) {
	var product []models.ProductBrief
	err := pr.DB.Raw("SELECT id, name, description, category_id, size, stock, price, product_status FROM products WHERE user_id = ?").Scan(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (pr *ProductRepository) DeleteProducts(product_id int) error {
	var count int
	err := pr.DB.Raw("SELECT COUNT(*) FROM products WHERE id = ?", product_id).Scan(&count).Error
	if err != nil {
		return err
	}
	if count < 1 {
		return errors.New("product for given id does not exist")
	}
	if err := pr.DB.Exec("DELETE FROM products WHERE id = ?", product_id).Error; err != nil {
		return err
	}
	return nil
}

func (pr *ProductRepository) CheckProduct(pid int) (bool, error) {
	var count int
	err := pr.DB.Raw("SELECT COUNT(*) FROM products WHERE id = ?", pid).Scan(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, err
	}
	return count > 0, nil
}

func (pr *ProductRepository) UpdateProducts(product models.Product) (domain.Product, error) {
	var pro domain.Product

	//Check if product exists first
	if err := pr.DB.First(&pro, product.ID).Error; err != nil {
		return domain.Product{}, err // product not found
	}
	err := pr.DB.Model(&pro).Updates(domain.Product{
		Name:        product.Name,
		Description: product.Description,
		CategoryID:  product.CategoryID,
		Size:        product.Size,
		Stock:       product.Stock,
		Price:       product.Price,
	}).Error
	if err != nil {
		return domain.Product{}, err
	}
	//Fetch the updated product
	if err := pr.DB.First(&pro, product.ID).Error; err != nil {
		return domain.Product{}, e
	}
	return pro, nil
}

func (pr *ProductRepository) GetQuantityFromProductID(id int) (int, error) {
	var quantity int
	err := pr.DB.Raw("SELECT stock FROM products WHERE id = ?", id).Scan(&quantity).Error
	if err != nil {
		return 0.0, err
	}
	return quantity, nil
}

func (pr *ProductRepository) GetPriceofProductFromID(id int) (float64, error) {
	var productPrice float64
	err := pr.DB.Raw("SELECT price FROM products WHERE id = ?", id).Scan(&productPrice).Error
	if err != nil {
		return 0.0, err
	}
	return productPrice, nil
}

func (pr *ProductRepository) ProductStockMinus(productID, stock int) error {
	err := pr.DB.Exec("UPDATE products SET stock = stock - ? WHERE id = ?", stock, productID).Error
	if err != nil {
		return err
	}
	return nil
}
