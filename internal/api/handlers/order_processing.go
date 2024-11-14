package handlers

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"order-processing/internal/api"
	"order-processing/internal/api/schema"
	"order-processing/internal/services"
	"strconv"
)

type OrderProcessingHandler struct {
	validationInstance *validator.Validate
	service            *services.OrderProcessingService
}

func NewOrderProcessingHandler(service *services.OrderProcessingService) *OrderProcessingHandler {
	return &OrderProcessingHandler{
		service:            service,
		validationInstance: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (h *OrderProcessingHandler) HandleRoot(c *fiber.Ctx) error {
	return api.APIResponse{
		Code: http.StatusOK,
		Msg:  "Hello, World!",
	}
}

func (h *OrderProcessingHandler) HandleListCustomers(c *fiber.Ctx) error {
	customers, err := h.service.ListCustomers(c.Context())
	if err != nil {
		log.Errorf("HandleListCustomers: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return api.APIError{
				Code: http.StatusNotFound,
				Err:  "Records not found",
			}
		} else {
			return api.APIError{
				Code: http.StatusInternalServerError,
				Err:  "Internal Server Error",
			}
		}
	}

	return api.APIResponse{
		Code: http.StatusOK,
		Msg: fiber.Map{
			"customers": customers,
		},
	}
}

func (h *OrderProcessingHandler) HandleGetCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return api.APIError{
			Code: http.StatusBadRequest,
			Err:  "ID is required",
		}
	}
	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return api.APIError{
			Code: http.StatusBadRequest,
			Err:  "Invalid ID",
		}
	}

	customer, err := h.service.GetCustomer(c.Context(), uint(parsedId))
	if err != nil {
		log.Errorf("HandleGetCustomer: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return api.APIError{
				Code: http.StatusNotFound,
				Err:  "Records not found",
			}
		} else {
			return api.APIError{
				Code: http.StatusInternalServerError,
				Err:  "Internal Server Error",
			}
		}
	}

	return api.APIResponse{
		Code: http.StatusOK,
		Msg: fiber.Map{
			"customer": customer,
		},
	}
}

func (h *OrderProcessingHandler) HandleCreateOrder(c *fiber.Ctx) error {
	var req schema.CreateOrderRequestSchema
	// bind request json with schema
	if err := c.BodyParser(&req); err != nil {
		log.Errorf("HandleCreateOrder: %v", err)
		return api.APIError{
			Code: http.StatusBadRequest,
			Err:  err.Error(),
		}
	}

	// Validate the request schema
	if err := h.validationInstance.Struct(&req); err != nil {
		return api.APIError{
			Code: http.StatusBadRequest,
			Err:  err.Error(),
		}
	}

	createdOrder, err := h.service.CreateOrder(c.Context(), req)
	if err != nil {
		log.Errorf("HandleCreateOrder: %v", err)
		return api.APIError{
			Code: http.StatusInternalServerError,
			Err:  "Internal Server Error",
		}
	}

	return api.APIResponse{
		Code: http.StatusOK,
		Msg: fiber.Map{
			"order": createdOrder,
		},
	}
}

func (h *OrderProcessingHandler) HandleGetOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return api.APIError{
			Code: http.StatusBadRequest,
			Err:  "ID is required",
		}
	}
	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return api.APIError{
			Code: http.StatusBadRequest,
			Err:  "Invalid ID",
		}
	}

	order, err := h.service.GetOrder(c.Context(), uint(parsedId))
	if err != nil {
		log.Errorf("HandleGetOrder: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return api.APIError{
				Code: http.StatusNotFound,
				Err:  "Records not found",
			}
		} else {
			return api.APIError{
				Code: http.StatusInternalServerError,
				Err:  "Internal Server Error",
			}
		}
	}

	return api.APIResponse{
		Code: http.StatusOK,
		Msg: fiber.Map{
			"order": order,
		},
	}
}

func (h *OrderProcessingHandler) HandleListProducts(c *fiber.Ctx) error {
	products, err := h.service.ListProducts(c.Context())
	if err != nil {
		log.Errorf("HandleListProducts: %v", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return api.APIError{
				Code: http.StatusNotFound,
				Err:  "Records not found",
			}
		} else {
			return api.APIError{
				Code: http.StatusInternalServerError,
				Err:  "Internal Server Error",
			}
		}
	}

	return api.APIResponse{
		Code: http.StatusOK,
		Msg: fiber.Map{
			"products": products,
		},
	}
}
