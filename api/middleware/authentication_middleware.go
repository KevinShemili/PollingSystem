package middleware

import (
	"gin/application/repository/contracts"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthenticationMiddleware(UnitOfWork contracts.IUnitOfWork) gin.HandlerFunc {
	return func(c *gin.Context) {

		jwtSigningKey := os.Getenv("SECRET_JWT")

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		jwtToken, err := jwt.Parse(authHeader, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}

			return []byte(jwtSigningKey), nil
		})

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {

			if exp, ok := claims["exp"].(float64); ok && float64(time.Now().Unix()) > exp {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			userID := uint(claims["sub"].(float64))

			user, err := UnitOfWork.IUserRepository().GetByID(userID)

			if err != nil || user == nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			c.Set("user", user)
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
