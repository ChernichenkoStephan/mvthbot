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

// HandleVariable godoc
// @Summary      Sets result of equation to variable
// @Description  Sets result of equation given in path variable to variable
// @Tags         user,solve,api
// @Produce      json
// @Param        name     path string true "variables name"
// @Param        equation path string true "equation for solve (2+2) encoded in LF"
// @Success      200
// @Router       /{name}/{equation} [post]
func (h *variableHandler) HandleVariable(c *fiber.Ctx) error {
	uID := c.Locals("userID").(int64)
	n := c.Params("name")
	e := c.Params("equation")
	ctx := c.UserContext()

	decoded := utils.DecodeLF(e)

	eq, err := converting.ToRPN(decoded)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return errors.Wrap(err, "Convertion to RPN fail")
	}

	var res float64
	err = h.db.WithinTransaction(ctx, func(ctx context.Context) error {
		vs, err := h.db.GetAllVariables(ctx, uID)
		if err != nil {
			return errors.Wrap(err, "Get variables fail")
		}

		res, err = solving.Solve(eq, vs)
		if err != nil {
			return errors.Wrap(err, "Solving fail")
		}

		if err != nil {
			return errors.Wrap(err, "Variables saving fail")
		}

		st := &solving.Statement{
			Variables: []string{n},
			Equation:  decoded,
			Value:     res,
		}

		err = h.db.AddStatement(ctx, uID, st)
		if err != nil {
			return errors.Wrap(err, "Statement saving fail")
		}

		return nil
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": res,
	})
}

// HandleVariables godoc
// @Summary      Processes multiple statements
// @Description  Processes multiple statements given in body
// @Tags         user,solve,api
// @Accept       json
// @Produce      json
// @Success      200
// @Router       / [post]
func (h *variableHandler) HandleVariables(c *fiber.Ctx) error {
	uID := c.Locals("userID").(int64)
	ctx := c.UserContext()

	req := &PackDTO{}

	if err := c.BodyParser(req); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return errors.Wrap(err, "Parsing fail")
	}

	vals := make([]float64, 0)
	err := h.db.WithinTransaction(ctx, func(ctx context.Context) error {

		vs, err := h.db.GetAllVariables(ctx, uID)
		if err != nil {
			return errors.Wrap(err, "Getting variables fail")
		}

		for _, st := range req.Statements {
			eq, err := converting.ToRPN(st.Equation)
			if err != nil {
				return errors.Wrap(err, "Converting to RPN fail")
			}

			res, err := solving.Solve(eq, vs)
			if err != nil {
				return errors.Wrap(err, "Solving fail")
			}
			vals = append(vals, res)

			ost := &solving.Statement{
				Variables: st.Names,
				Equation:  st.Equation,
				Value:     res,
			}
			err = h.db.AddStatement(ctx, uID, ost)
			if err != nil {
				return errors.Wrap(err, "Statement saving fail")
			}

			if err != nil {
				return errors.Wrap(err, "Names saving fail")
			}

			for _, n := range st.Names {
				vs[n] = res
			}
		}

		return nil
	})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"values": vals,
	})
}

// GetVariable godoc
// @Summary      Returns users variable value
// @Description  Returns users variable value (user from given JWT)
// @Tags         user,api
// @Produce      json
// @Param        name     path string true "variables name"
// @Success      200
// @Router       /{name} [get]
func (h *variableHandler) GetVariable(c *fiber.Ctx) error {
	uID := c.Locals("userID").(int64)
	n := c.Params("name")
	ctx := c.UserContext()

	v, err := h.db.GetVariable(ctx, uID, n)
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

// GetVariables godoc
// @Summary      Returns users variables value
// @Description  Returns users variables (you can ask for specific in body) (user from given JWT)
// @Tags         user,api
// @Accept       json
// @Produce      json
// @Success      200
// @Router       / [get]
func (h *variableHandler) GetVariables(c *fiber.Ctx) error {
	uID := c.Locals("userID").(int64)
	ctx := c.UserContext()

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

	vars, err := h.db.GetVariablesWithNames(ctx, uID, req.Names)
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

// GetVariable godoc
// @Summary      Deletes users variable
// @Description  Deletes users variable (user from given JWT)
// @Tags         user,api
// @Accept       json
// @Produce      json
// @Param        name     path string true "variables name"
// @Success      200
// @Router       /{name} [delete]
func (h *variableHandler) DeleteVariable(c *fiber.Ctx) error {
	uID := c.Locals("userID").(int64)
	n := c.Params("name")
	ctx := c.UserContext()

	if n == "" {
		err := h.db.DeleteAllVariables(ctx, uID)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
			return errors.Wrap(err, "Deleting all variables fail")
		}
	}

	err := h.db.DeleteVariable(ctx, uID, n)
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

// DeleteAllVariables godoc
// @Summary      Deletes users variables
// @Description  Deletes users variables (you can ask for specific in body) (user from given JWT)
// @Tags         user,api
// @Accept       json
// @Produce      json
// @Success      200
// @Router       / [delete]
func (h *variableHandler) DeleteAllVariables(c *fiber.Ctx) error {
	uID := c.Locals("userID").(int64)
	ctx := c.UserContext()

	if err := h.db.DeleteAllVariables(ctx, uID); err != nil {
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

// HandleHistory godoc
// @Summary      Returns history
// @Description  Returns variables and statements hiistory for user defined by JWT
// @Tags         user,api
// @Produce      json
// @Success      200
// @Router       /history [get]
func (h *historyHandler) HandleHistory(c *fiber.Ctx) error {
	uID := c.Locals("userID").(int64)
	ctx := c.UserContext()

	hist, err := h.db.GetHistory(ctx, uID)
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

// DeleteHistory godoc
// @Summary      Deletes all users history
// @Description  Clears variables and statements hiistory from user defined by JWT
// @Tags         user,api
// @Produce      json
// @Success      200
// @Router       /history [delete]
func (h *historyHandler) DeleteHistory(c *fiber.Ctx) error {
	uID := c.Locals("userID").(int64)
	ctx := c.UserContext()

	if err := h.db.Clear(ctx, uID); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return errors.Wrap(err, "Solving fail")
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success",
	})
}
