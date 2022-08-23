package user

import (
	"context"

	"emperror.dev/errors"
	"github.com/ChernichenkoStephan/mvthbot/internal/converting"
	"github.com/ChernichenkoStephan/mvthbot/internal/solving"
	"github.com/ChernichenkoStephan/mvthbot/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// Variables
type variableHandler struct {
	db *Database
}

func NewVariableHandler(userRoute fiber.Router, db *Database) {
	h := &variableHandler{
		db: db,
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
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return errors.Wrap(err, "Convertion to RPN fail")
	}

	vs, err := h.db.GetAllVariables(context.TODO(), uID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return errors.Wrap(err, "Get variables fail")
	}

	res, err := solving.Solve(eq, vs)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return errors.Wrap(err, "Solving fail")
	}

	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return errors.Wrap(err, "Variables saving fail")
	}

	st := &solving.Statement{
		Variables: []string{n},
		Equation:  decoded,
		Value:     res,
	}

	err = h.db.AddStatement(context.TODO(), uID, st)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return errors.Wrap(err, "Statement saving fail")
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
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return errors.Wrap(err, "Parsing fail")
	}

	vs, err := h.db.GetAllVariables(context.TODO(), uID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return errors.Wrap(err, "Getting variables fail")
	}

	vals := make([]float64, 0)
	for _, st := range req.Statements {
		eq, err := converting.ToRPN(st.Equation)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
			return errors.Wrap(err, "Converting to RPN fail")
		}

		res, err := solving.Solve(eq, vs)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
			return errors.Wrap(err, "Solving fail")
		}
		vals = append(vals, res)

		ost := &solving.Statement{
			Variables: st.Names,
			Equation:  st.Equation,
			Value:     res,
		}
		err = h.db.AddStatement(context.TODO(), uID, ost)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
			return errors.Wrap(err, "Statement saving fail")
		}

		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
			return errors.Wrap(err, "Names saving fail")
		}

		for _, n := range st.Names {
			vs[n] = res
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"values": vals,
	})
}

func (h *variableHandler) GetVariable(c *fiber.Ctx) error {
	uID := c.Locals("userID").(int64)
	n := c.Params("name")

	v, err := h.db.GetVariable(context.TODO(), uID, n)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return errors.Wrap(err, "Getting variables fail")
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
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return errors.Wrap(err, "Request parsing fail")
	}

	vars, err := h.db.GetVariablesWithNames(context.TODO(), uID, req.Names)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return errors.Wrap(err, "Getting variables fail")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"variables": vars,
	})
}

func (h *variableHandler) DeleteVariable(c *fiber.Ctx) error {
	uID := c.Locals("userID").(int64)
	n := c.Params("name")

	if n == "" {
		err := h.db.DeleteAllVariables(context.TODO(), uID)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
			return errors.Wrap(err, "Deleting all variables fail")
		}
	}

	err := h.db.DeleteVariable(context.TODO(), uID, n)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return errors.Wrap(err, "Deletiing variables fail")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
	})
}

func (h *variableHandler) DeleteAllVariables(c *fiber.Ctx) error {
	uID := c.Locals("userID").(int64)

	if err := h.db.DeleteAllVariables(context.TODO(), uID); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return errors.Wrap(err, "Deleting variables fail")
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
	db *Database
}

func NewHistoryHandler(userRoute fiber.Router, db *Database) {
	h := &historyHandler{
		db: db,
	}

	userRoute.Get("", GetUserIDFromJWT, h.HandleHistory)
	userRoute.Delete("", GetUserIDFromJWT, h.DeleteHistory)
}

func (h *historyHandler) HandleHistory(c *fiber.Ctx) error {
	uID := c.Locals("userID").(int64)

	hist, err := h.db.GetHistory(context.TODO(), uID)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return errors.Wrap(err, "Solving fail")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"history": hist,
	})
}

func (h *historyHandler) DeleteHistory(c *fiber.Ctx) error {
	uID := c.Locals("userID").(int64)
	if err := h.db.Clear(context.TODO(), uID); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return errors.Wrap(err, "Solving fail")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
	})
}
