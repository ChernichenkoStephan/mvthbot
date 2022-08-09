package solving

import (
	"fmt"

	"github.com/ChernichenkoStephan/mvthbot/internal/converting"
	"github.com/ChernichenkoStephan/mvthbot/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type SolvePackDTO struct {
	Equations []string `validate:"required"`
}

// Solve
type solveHandler struct{}

func NewSolveHandler(userRoute fiber.Router) {
	h := &solveHandler{}

	userRoute.Post("", h.HandleMultipleSolve)
	userRoute.Post("/", h.HandleMultipleSolve)
	userRoute.Post("/:equation", h.HandleSolve)
}

func (h *solveHandler) HandleSolve(c *fiber.Ctx) error {
	decoded := utils.DecodeLF(c.Params("equation"))

	eq, err := converting.ToRPN(decoded)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res, err := Solve(eq, VMap{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": res,
	})
}

func (h *solveHandler) HandleMultipleSolve(c *fiber.Ctx) error {
	p := new(SolvePackDTO)

	if err := c.BodyParser(p); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	vals := make([]float64, 0)
	errs := make([]string, 0)
	for i, e := range p.Equations {
		eStr := ""
		val := 0.0

		if eq, err := converting.ToRPN(e); err != nil {
			eStr = fmt.Sprintf("Error %v\nin %d equation", err, i+1)
		} else {
			val, err = Solve(eq, VMap{})
			if err != nil {
				eStr = fmt.Sprintf("Error %v\nin %d equation", err, i+1)
			}
		}
		errs = append(errs, eStr)
		vals = append(vals, val)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"results": vals,
		"errors":  errs,
	})
}
