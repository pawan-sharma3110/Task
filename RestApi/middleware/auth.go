package middleware

import (
	"fmt"
	"net/http"
	"rest-api/utils"

	"github.com/gin-gonic/gin"
)

func Autenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Not authorized "})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("error : %v", err)})
		return
	}
	context.Set("userId",userId)
	context.Next()
}
