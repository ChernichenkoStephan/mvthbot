package auth

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt/v4"
)

func NewAuthHandler(authRoute fiber.Router, repo AuthRepository, getter IDGetter) {
	handler := &AuthHandler{
		repository: repo,
		idGetter:   getter,
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
		log.Printf("[signInUser] parsing error")
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}
	log.Printf("[signInUser] request: %v", request.Username)

	id, err := strconv.ParseInt(request.Username, 0, 64)

	// if got id in request
	if err != nil {
		id, err = h.idGetter.GetUserID(request.Username)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
				"status":  "fail",
				"message": err.Error(),
			})
		}
	}

	pass, err := h.repository.GetPassword(context.TODO(), id)
	if err != nil {
		log.Printf("User password error")
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	if pass != request.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"status":  "fail",
			"message": "Wrong password!",
		})
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
		log.Printf("Token generation fault")
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    signedToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24),
		Secure:   false,
		HTTPOnly: true,
	})

	log.Printf("'%v' Login success", request.Username)
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
