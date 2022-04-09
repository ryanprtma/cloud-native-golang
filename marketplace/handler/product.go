package handler

import (
	"marketplace/helper"
	"marketplace/product"
	"marketplace/role"
	"marketplace/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService product.Service
	roleService    role.Service
}

func NewProductHandler(productService product.Service, roleServie role.Service) *productHandler {
	return &productHandler{productService: productService, roleService: roleServie}
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	var input product.CreateProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIresponse("Failed to add product", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	getMerchant, err := h.roleService.GetMerchantByUserID(currentUser.ID.String())
	if err != nil {
		response := helper.APIresponse("Failed to get merchant", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	input.Merchant = getMerchant

	newProduct, err := h.productService.CreateProduct(input)
	if err != nil {
		response := helper.APIresponse("Failed to add product", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIresponse("Success to add product", http.StatusOK, "success", product.FormatProduct(newProduct))
	c.JSON(http.StatusOK, response)
}
