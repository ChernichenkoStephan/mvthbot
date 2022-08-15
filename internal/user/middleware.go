package user

import (
	"fmt"
	"strconv"

	"emperror.dev/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetUserIDFromJWT(c *fiber.Ctx) error {
	jwtData := c.Locals("user").(*jwt.Token)
	claims := jwtData.Claims.(jwt.MapClaims)

	id, err := strconv.ParseInt(claims["uid"].(string), 0, 64)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
		return errors.Wrap(err,
			fmt.Sprintf("User id shuld be number, got %s", claims["uid"].(string)))
	}
	c.Locals("userID", id)
	return c.Next()
}
