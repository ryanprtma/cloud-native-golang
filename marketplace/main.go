package main

import (
	"log"
	"marketplace/auth"
	"marketplace/config"
	"marketplace/handler"
	"marketplace/helper"
	"marketplace/product"
	"marketplace/role"
	"marketplace/user"
	"net/http"
	"strings"

	_ "marketplace/docs"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	db, err := config.Getdb()
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepo := user.NewRepository(db)
	roleRepo := role.NewRepository(db)
	productRepo := product.NewRepository(db)
	userService := user.NewService(userRepo)
	roleService := role.NewService(roleRepo)
	productService := product.NewService(productRepo)

	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService, roleService)
	productHandler := handler.NewProductHandler(productService, roleService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/profilepictureupload", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.POST("/add_product", authMiddleware(authService, userService), roleMiddleware(), productHandler.CreateProduct)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()

}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIresponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIresponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIresponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := claim["user_id"].(string)

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIresponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}

func roleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		currentUser := c.MustGet("currentUser").(user.User)

		roleId := currentUser.RoleId

		if roleId != 1 {
			response := helper.APIresponse("User are not merchant", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

	}
}
