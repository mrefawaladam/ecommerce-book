package middleware

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

const SECRET_JWT = "123"

func CreateToken(userID int, email string, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = userID
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(SECRET_JWT))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func RequireRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			userRole := claims["role"].(string)

			if userRole != role {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized access"})
			}

			return next(c)
		}
	}
}

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.Request().Header.Get("Authorization")
			if tokenString == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Token is missing"})
			}

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid token")
				}
				return []byte(SECRET_JWT), nil
			})

			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
			}

			userID, ok := claims["id"].(float64)
			if !ok {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
			}

			c.Set("userID", int(userID))

			return next(c)
		}
	}
}
