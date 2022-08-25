package solving

import (
	"fmt"

	"emperror.dev/errors"
	"github.com/ChernichenkoStephan/mvthbot/internal/converting"
	"github.com/ChernichenkoStephan/mvthbot/internal/utils"
	"github.com/gofiber/fiber/v2"
)

// Solve
type solveHandler struct{}

func NewSolveHandler(userRoute fiber.Router) {
	h := &solveHandler{}

	userRoute.Post("", h.HandleMultipleSolve)
	userRoute.Post("/", h.HandleMultipleSolve)
	userRoute.Post("/:equation", h.HandleSolve)
}

// Solve godoc
// @Summary      Basic equations solving
// @Description  Solves equation given in url param
// @Tags         solve,api
// @Produce      json
// @Param        equation path string true "equation for solve (2+2) encoded in LF"
// @Success      200
// @Router       /solve/{equation} [post]
func (h *solveHandler) HandleSolve(c *fiber.Ctx) error {
	decoded := utils.DecodeLF(c.Params("equation"))

	eq, err := converting.ToRPN(decoded)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return errors.Wrap(err, "Converting to RPN fail")
	}

	res, err := Solve(eq, map[string]float64{})
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return errors.Wrap(err, "Solving failed")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": res,
	})
}

// MultiSolve godoc
// @Summary      Multiple solving of equations
// @Description  Solves equations array given in request body
// @Tags         solve,api
// @Accept       json
// @Produce      json
// @Success      200
// @Router       /solve [post]
func (h *solveHandler) HandleMultipleSolve(c *fiber.Ctx) error {
	p := new(PackDTO)

	if err := c.BodyParser(p); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
		return errors.Wrap(err, "Body parsing failed")
	}

	vals := make([]float64, 0)
	errs := make([]string, 0)
	for i, e := range p.Equations {
		eStr := ""
		val := 0.0

		if eq, err := converting.ToRPN(e); err != nil {
			eStr = fmt.Sprintf("Error %v\nin %d equation", err, i+1)
		} else {
			val, err = Solve(eq, map[string]float64{})
			if err != nil {
				eStr = fmt.Sprintf("Error %v\nin %d equation", err, i+1)
			}
		}
		errs = append(errs, eStr)
		vals = append(vals, val)
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"results": vals,
		"error":   errs,
	})

	if len(errs) != 0 {
		return MultySolveError{errs}
	}
	return nil

}
