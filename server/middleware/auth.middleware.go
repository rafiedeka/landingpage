package middleware

import (
	"landingpage/auth"
	"landingpage/helper"
	"landingpage/user"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(userService user.Service, authService auth.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			abortUnauthorized(c)
			return
		}
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			abortUnauthorized(c)
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			abortUnauthorized(c)
			return
		}
		userID := int(claim["user_id"].(float64))

		// Ambil data user dari database
		user, err := userService.GetUserByID(userID)
		if err != nil {
			abortUnauthorized(c)
			return
		}

		// Simpan data user di context
		c.Set("currentUser", user)
		c.Next()
	}
}

func abortUnauthorized(c *gin.Context) {
	response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
}
