package user

import "github.com/gofiber/fiber/v2"

// Variables
type variableHandler struct {
	userService UserService
}

func NewVariableHandler(userRoute fiber.Router, us UserService) {
	h := &variableHandler{
		userService: us,
	}

	userRoute.Post("", h.HandleVariable)
}

func (h *variableHandler) HandleVariable(c *fiber.Ctx) error {
	return nil
}

//
//
// History
//
//

type historyHandler struct {
	userService UserService
}

func NewHistoryHandler(userRoute fiber.Router, us UserService) {
	h := &historyHandler{
		userService: us,
	}

	userRoute.Post("", h.HandleHistory)
}

func (h *historyHandler) HandleHistory(c *fiber.Ctx) error {
	return nil
}
