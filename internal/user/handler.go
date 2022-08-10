package user

import (
	"context"
	"fmt"

	"github.com/ChernichenkoStephan/mvthbot/internal/converting"
	"github.com/ChernichenkoStephan/mvthbot/internal/solving"
	"github.com/ChernichenkoStephan/mvthbot/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// Variables
type variableHandler struct {
	userService     UserService
	variableService VariableService
}

func NewVariableHandler(userRoute fiber.Router, us UserService, vs VariableService) {
	h := &variableHandler{
		userService:     us,
		variableService: vs,
	}

	userRoute.Post("/:name/:equation", GetUserIDFromJWT, h.HandleVariable)
	userRoute.Post("/", GetUserIDFromJWT, h.HandleVariables)

	userRoute.Get("/:name", GetUserIDFromJWT, h.GetVariable)
	userRoute.Get("/", GetUserIDFromJWT, h.GetVariables)

	userRoute.Delete("/:name", GetUserIDFromJWT, h.DeleteVariable)
	userRoute.Delete("", GetUserIDFromJWT, h.DeleteAllVariables)

}

func (h *variableHandler) HandleVariable(c *fiber.Ctx) error {
	uID := c.Locals("userID").(int64)
	n := c.Params("name")
	e := c.Params("equation")

	decoded := utils.DecodeLF(e)

	eq, err := converting.ToRPN(decoded)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	vs, err := h.variableService.FetchAllUserVariables(context.TODO(), uID)
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

	err = h.variableService.AddUserVariable(context.TODO(), uID, n, res)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	st := &solving.Statement{
		Variables: []string{n},
		Equation:  decoded,
		Value:     res,
	}
	fmt.Printf("AddStatement %v", st)
	err = h.userService.AddStatement(context.TODO(), uID, st)
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
	uID := c.Locals("userID").(int64)

	type statement struct {
		Names    []string `json:"names"`
		Equation string   `json:"equation"`
	}

	type setVariablesRequest struct {
		Statements []statement `json:"statements"`
	}

	type variable struct {
		Name  string
		Value float64
	}

	req := &setVariablesRequest{}

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	vs, err := h.variableService.FetchAllUserVariables(context.TODO(), uID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	vals := make([]float64, 0)
	for _, st := range req.Statements {
		eq, err := converting.ToRPN(st.Equation)
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
		vals = append(vals, res)

		ost := &solving.Statement{
			Variables: st.Names,
			Equation:  st.Equation,
			Value:     res,
		}
		err = h.userService.AddStatement(context.TODO(), uID, ost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		err = h.variableService.AddUserVariables(context.TODO(), uID, st.Names, res)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"values": vals,
	})
}

func (h *variableHandler) GetVariable(c *fiber.Ctx) error {
	uID := c.Locals("userID").(int64)
	n := c.Params("name")

	v, err := h.variableService.FetchUserVariable(context.TODO(), uID, n)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"value": v,
	})
}

func (h *variableHandler) GetVariables(c *fiber.Ctx) error {
	uID := c.Locals("userID").(int64)

	type variablesRequest struct {
		Names []string `json:"names"`
	}

	type variable struct {
		Name  string
		Value float64
	}

	req := &variablesRequest{}

	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	vars, err := h.variableService.FetchUserVariables(context.TODO(), uID, req.Names)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"variables": vars,
	})
}

func (h *variableHandler) DeleteVariable(c *fiber.Ctx) error {
	uID := c.Locals("userID").(int64)
	n := c.Params("name")

	if n == "" {
		err := h.variableService.ClearUserVariables(context.TODO(), uID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	err := h.variableService.DeleteUserVariable(context.TODO(), uID, n)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
	})
}

func (h *variableHandler) DeleteAllVariables(c *fiber.Ctx) error {
	uID := c.Locals("userID").(int64)

	if err := h.variableService.ClearUserVariables(context.TODO(), uID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
	})

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

	userRoute.Get("", GetUserIDFromJWT, h.HandleHistory)
	userRoute.Delete("", GetUserIDFromJWT, h.DeleteHistory)
}

func (h *historyHandler) HandleHistory(c *fiber.Ctx) error {
	uID := c.Locals("userID").(int64)

	hist, err := h.userService.FetchHistory(context.TODO(), uID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"history": hist,
	})
}

func (h *historyHandler) DeleteHistory(c *fiber.Ctx) error {
	uID := c.Locals("userID").(int64)
	if err := h.userService.ClearHistory(context.TODO(), uID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
	})
}
