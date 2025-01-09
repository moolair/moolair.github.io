package middleware

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authClient *auth.Client, allowedUID string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authorization 헤더에서 토큰 추출
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		idToken := strings.TrimPrefix(authHeader, "Bearer ")

		// Firebase 토큰 검증
		token, err := authClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// 관리자 UID 확인
		if token.UID != allowedUID {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			return
		}

		c.Next()
	}
}
