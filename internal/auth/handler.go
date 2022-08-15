package auth

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"emperror.dev/errors"
	"go.uber.org/zap"

	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt/v4"
)

func NewAuthHandler(authRoute fiber.Router, repo AuthRepository, getter IDGetter, lg *zap.SugaredLogger) {
	handler := &AuthHandler{
		repository: repo,
		idGetter:   getter,
		logger:     lg,
	}

	authRoute.Post("/login", handler.signInUser)
	authRoute.Post("/logout", handler.signOutUser)
	authRoute.Get("/private", JWTMiddleware(), handler.privateRoute)
}

func (h *AuthHandler) signInUser(c *fiber.Ctx) error {

	type jwtClaims struct {
		UserID string `json:"uid"`
		User   string `json:"user"`
		jwt.StandardClaims
	}

	request := &loginRequest{}
	if err := c.BodyParser(request); err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
		return errors.Wrap(err, "Error during request parsing")
	}

	id, err := strconv.ParseInt(request.Username, 0, 64)

	// if got id in request
	if err != nil {
		id, err = h.idGetter.GetUserID(request.Username)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			})
			return errors.Wrap(err, "Error during id fetching")
		}
	}

	pass, err := h.repository.GetPassword(context.TODO(), id)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
		return errors.Wrap(err, "Error during password request")
	}

	if pass != request.Password {
		c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Wrong password!",
		})
		return errors.New("Wrong password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtClaims{
		fmt.Sprintf("%d", id),
		request.Username,
		jwt.StandardClaims{
			Audience:  "mvthbot users",
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    "mvthbot service",
		},
	})

	signedToken, err := token.SignedString([]byte(_TEST_SECRET))
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
		return errors.Wrap(err, "Error during JWT generation")
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    signedToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24),
		Secure:   false,
		HTTPOnly: true,
	})

	h.logger.Infof("User '%v' login success\n", request.Username)
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status": "success",
		"token":  signedToken,
	})
}

func (h *AuthHandler) signOutUser(c *fiber.Ctx) error {

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "loggedOut",
		Path:     "/",
		Expires:  time.Now().Add(time.Second * 10),
		Secure:   false,
		HTTPOnly: true,
	})

	jwtData := c.Locals("user").(*jwt.Token)
	claims := jwtData.Claims.(jwt.MapClaims)

	h.logger.Infof("User '%v' logout\n", claims["uid"])
	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Logged out successfully!",
	})
}

func (h *AuthHandler) privateRoute(c *fiber.Ctx) error {
	type jwtResponse struct {
		UserID interface{} `json:"uid"`
		User   interface{} `json:"user"`
		Iss    interface{} `json:"iss"`
		Aud    interface{} `json:"aud"`
		Exp    interface{} `json:"exp"`
	}

	jwtData := c.Locals("user").(*jwt.Token)
	claims := jwtData.Claims.(jwt.MapClaims)

	jwtResp := &jwtResponse{
		UserID: claims["uid"],
		User:   claims["user"],
		Iss:    claims["iss"],
		Aud:    claims["aud"],
		Exp:    claims["exp"],
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "success",
		"message": "Welcome to the private route!",
		"jwtData": jwtResp,
	})
}
