package middleware

import (
	"gin/application/repository/contracts"
	"gin/application/utility"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthenticationMiddleware(UnitOfWork contracts.IUnitOfWork, jwtSigningKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		decodedToken, err := utility.Decode(authHeader)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		jwtToken, err := jwt.Parse(decodedToken, func(t *jwt.Token) (interface{}, error) {
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

			user, err := UnitOfWork.Users().GetByID(userID)

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
