package misc

import "github.com/gofiber/fiber/v2"

type MiscHandler struct{}

func NewMiscHandler(miscRoute fiber.Router) {
	handler := &MiscHandler{}

	miscRoute.Get("/health", handler.healthCheck)
}

// Check for the health of the API.
func (h *MiscHandler) healthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Service working fine",
	})
}
