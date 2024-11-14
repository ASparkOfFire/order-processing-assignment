package repository

import (
	"context"
	"gorm.io/gorm"
	"order-processing/internal/models"
)

type OrderProcessor interface {
	ListCustomers(context.Context) ([]*models.Customer, error)
	GetCustomer(context.Context, uint) (*models.Customer, error)
	CreateOrder(context.Context, models.Order) (*models.Order, error)
	GetOrder(context.Context, uint) (*models.Order, error)

	CreateCustomer(context.Context, *models.Customer) (*models.Customer, error)
	CreateProduct(context.Context, models.Product) (*models.Product, error)
	ListProducts(context.Context) ([]*models.Product, error)
}

type PostgresOrderProcessor struct {
	db *gorm.DB
}

func NewPostgresOrderProcessor(db *gorm.DB) OrderProcessor {
	return PostgresOrderProcessor{db: db}
}

func (p PostgresOrderProcessor) ListCustomers(ctx context.Context) ([]*models.Customer, error) {
	var customers []*models.Customer
	if err := p.db.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}

func (p PostgresOrderProcessor) GetCustomer(ctx context.Context, id uint) (*models.Customer, error) {
	var customer models.Customer
	if err := p.db.First(&customer, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

func (p PostgresOrderProcessor) CreateOrder(ctx context.Context, order models.Order) (*models.Order, error) {
	if err := p.db.Create(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (p PostgresOrderProcessor) GetOrder(ctx context.Context, id uint) (*models.Order, error) {
	var order models.Order
	if err := p.db.First(&order, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (p PostgresOrderProcessor) CreateProduct(ctx context.Context, product models.Product) (*models.Product, error) {
	if err := p.db.Create(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p PostgresOrderProcessor) ListProducts(ctx context.Context) ([]*models.Product, error) {
	var products []*models.Product
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (p PostgresOrderProcessor) CreateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	if err := p.db.Create(customer).Error; err != nil {
		return nil, err
	}
	return customer, nil
}
