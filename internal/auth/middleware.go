package auth

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware() fiber.Handler {
	fmt.Println("JWTMiddleware")
	return jwtware.New(jwtware.Config{
		SigningKey:    []byte(_TEST_SECRET),
		TokenLookup:   "cookie:jwt",
		SigningMethod: "HS256",
	})
}

func GetDataFromJWT(c *fiber.Ctx) error {
	jwtData := c.Locals("user").(*jwt.Token)
	claims := jwtData.Claims.(jwt.MapClaims)

	parsedUserID := claims["uid"].(string)
	userID, err := strconv.Atoi(parsedUserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	c.Locals("currentUser", userID)
	return c.Next()
}
