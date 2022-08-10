package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	jwtware "github.com/gofiber/jwt/v3"
)

func JWTMiddleware() fiber.Handler {
	fmt.Println("JWTMiddleware")
	return jwtware.New(jwtware.Config{
		SigningKey:    []byte(_TEST_SECRET),
		TokenLookup:   "cookie:jwt",
		SigningMethod: "HS256",
	})
}
