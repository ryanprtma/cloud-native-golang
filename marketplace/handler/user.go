package handler

import (
	"marketplace/auth"
	"marketplace/helper"
	"marketplace/role"
	"marketplace/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
	roleService role.Service
}

func NewUserHandler(userService user.Service, authService auth.Service, roleService role.Service) *userHandler {
	return &userHandler{userService: userService, authService: authService, roleService: roleService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIresponse("registering account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newuser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.APIresponse("registering account failed", http.StatusBadRequest, "error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(newuser.ID.String())
	if err != nil {
		response := helper.APIresponse("generate token failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	h.roleService.SaveRole(newuser, input)

	formatter := user.FormatUser(newuser, token)

	response := helper.APIresponse("account has been created", http.StatusCreated, "created", formatter)

	c.JSON(http.StatusCreated, response)
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIresponse("login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIresponse("login failed", http.StatusUnauthorized, "error", errorMessage)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	token, err := h.authService.GenerateToken(loggedInUser.ID.String())
	if err != nil {
		response := helper.APIresponse("generate token failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(loggedInUser, token)

	response := helper.APIresponse("your now logged in", http.StatusOK, "succes", formatter)

	c.JSON(http.StatusOK, response)

}

// ShowAccount godoc
// @Summary      check email
// @Description  check email availablelity
// @Tags         user
// @Accept       json
// @Produce      json
// @Param		 user body user.CheckEmailInput true "user email"
// @Success      200  {object}  helper.Response
// @Failure      422  {object}  helper.Response
// @Failure      400  {object}  helper.Response
// @Failure      500  {object}  helper.Response
// @Router       /api/v1/email_checkers [post]
func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"error": errors}

		response := helper.APIresponse("check email failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailible, err := h.userService.CheckEmailAvailiblelity(input)

	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.APIresponse("check email failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var metaMessage string

	if isEmailAvailible {
		metaMessage = "email is available"
	} else {
		metaMessage = "email has been registered"
	}

	data := gin.H{
		"is_available": isEmailAvailible,
	}

	response := helper.APIresponse(metaMessage, http.StatusOK, "succes", data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIresponse("Failed to read data", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID.String()

	path := "images/" + userID + "-" + file.Filename

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIresponse("Failed to upload profile image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIresponse("Failed to save to database", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIresponse("Successfully uploaded images", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}
