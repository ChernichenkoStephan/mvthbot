package user

// Variables
type variableHandler struct {
	userService     UserService
	variableService VariableService
}

/*
func NewVariableHandler(userRoute fiber.Router, us UserService) {
	h := &variableHandler{
		userService: us,
	}

	userRoute.Post("/:name/:equation", h.HandleVariable)
	userRoute.Post("/", h.HandleVariables)
	//userRoute.Post("", h.HandleVariables)

	userRoute.Get("/:name", h.GetVariable)
	userRoute.Get("/", h.GetVariables)
	//userRoute.Get("", h.HandleVariables)

	userRoute.Delete("/:name", h.DeleteVariable)
	userRoute.Delete("/", h.DeleteVariables)
	//userRoute.Delete("", h.DeleteVariables)
}

func (h *variableHandler) HandleVariable(c *fiber.Ctx) error {
	n := c.Params("name")
	e := c.Params("equation")

	eq, err := converting.ToRPN(e)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	vs, err := h.variableService.FetchAllUserVariables(0)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res, err := solving.Solve(eq, vs)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = h.variableService.AddUserVariable(context.TODO(), 0, n, res)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": res,
	})
}

func (h *variableHandler) HandleVariables(c *fiber.Ctx) error {
	return nil
}

func (h *variableHandler) GetVariable(c *fiber.Ctx) error {
	return nil
}

func (h *variableHandler) GetVariables(c *fiber.Ctx) error {
	return nil
}

func (h *variableHandler) DeleteVariable(c *fiber.Ctx) error {
	return nil
}

func (h *variableHandler) DeleteVariables(c *fiber.Ctx) error {
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
*/
