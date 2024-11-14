package services

import (
	"context"
	"order-processing/internal/api/schema"
	"order-processing/internal/models"
	"order-processing/internal/repository"
)

type OrderProcessingService struct {
	repo repository.OrderProcessor
}

func NewOrderProcessingService(repo repository.OrderProcessor) *OrderProcessingService {
	return &OrderProcessingService{
		repo: repo,
	}
}

func (op OrderProcessingService) CreateCustomer(ctx context.Context, req models.Customer) (*models.Customer, error) {
	customer, err := op.repo.CreateCustomer(ctx, &req)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (op OrderProcessingService) ListCustomers(ctx context.Context) ([]*schema.Customer, error) {
	customers, err := op.repo.ListCustomers(ctx)
	if err != nil {
		return nil, err
	}
	var customerList []*schema.Customer
	for _, customer := range customers {
		customerList = append(customerList, &schema.Customer{
			ID:        customer.ID,
			CreatedAt: customer.CreatedAt,
			UpdatedAt: customer.UpdatedAt,
			Name:      customer.Name,
			Email:     customer.Email,
		})
	}
	return customerList, nil
}

func (op OrderProcessingService) GetCustomer(ctx context.Context, id uint) (*schema.Customer, error) {
	customer, err := op.repo.GetCustomer(ctx, id)
	if err != nil {
		return nil, err
	}
	return &schema.Customer{
		ID:        customer.ID,
		CreatedAt: customer.CreatedAt,
		UpdatedAt: customer.UpdatedAt,
		Name:      customer.Name,
		Email:     customer.Email,
	}, nil
}

func (op OrderProcessingService) GetOrder(ctx context.Context, id uint) (*schema.Order, error) {
	order, err := op.repo.GetOrder(ctx, id)
	if err != nil {
		return nil, err
	}

	var totalPrice float64

	var products []schema.Product
	for _, product := range order.Products {
		products = append(products, schema.Product{
			ID:        product.ID,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
			Name:      product.Name,
			Price:     product.Price,
		})
		totalPrice += product.Price
	}

	return &schema.Order{
		ID:         order.ID,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
		Products:   products,
		CustomerID: order.CustomerID,
		TotalPrice: totalPrice,
	}, nil
}

func (op OrderProcessingService) CreateOrder(ctx context.Context, orderReq schema.CreateOrderRequestSchema) (*schema.Order, error) {
	var totalPrice float64

	var products []models.Product
	for _, product := range orderReq.Products {
		products = append(products, models.Product{
			Name:  product.Name,
			Price: product.Price,
		})

		totalPrice += product.Price
	}

	data := models.Order{
		CustomerID: orderReq.CustomerID,
		Products:   products,
		TotalPrice: totalPrice,
	}

	createdOrder, err := op.repo.CreateOrder(ctx, data)
	if err != nil {
		return nil, err
	}

	return &schema.Order{
		ID:         createdOrder.ID,
		CreatedAt:  createdOrder.CreatedAt,
		UpdatedAt:  createdOrder.UpdatedAt,
		Products:   orderReq.Products,
		CustomerID: createdOrder.CustomerID,
		TotalPrice: totalPrice,
	}, nil
}

func (op OrderProcessingService) CreateProduct(ctx context.Context, prodReq schema.CreateProductRequestSchema) (*schema.Product, error) {
	data := models.Product{
		Name:  prodReq.Name,
		Price: prodReq.Price,
	}

	createdProduct, err := op.repo.CreateProduct(ctx, data)
	if err != nil {
		return nil, err
	}

	return &schema.Product{
		ID:        createdProduct.ID,
		CreatedAt: createdProduct.CreatedAt,
		UpdatedAt: createdProduct.UpdatedAt,
		Name:      createdProduct.Name,
		Price:     createdProduct.Price,
	}, nil
}

func (op OrderProcessingService) ListProducts(ctx context.Context) ([]*schema.Product, error) {
	products, err := op.repo.ListProducts(ctx)
	if err != nil {
		return nil, err
	}
	var productList []*schema.Product
	for _, product := range products {
		productList = append(productList, &schema.Product{
			ID:        product.ID,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
			Name:      product.Name,
			Price:     product.Price,
		})
	}

	return productList, nil
}
