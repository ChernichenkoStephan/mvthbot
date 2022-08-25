package misc

import "github.com/gofiber/fiber/v2"

type MiscHandler struct{}

func NewMiscHandler(miscRoute fiber.Router) {
	handler := &MiscHandler{}

	miscRoute.Get("/health", handler.healthCheck)
}

// HealthCheck godoc
// @Summary      Health check
// @Description  Check for the health of the API.
// @Tags         api
// @Success      200
// @Router       /health [get]
func (h *MiscHandler) healthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Service working fine",
	})
}
