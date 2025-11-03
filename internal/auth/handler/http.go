package handler

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"golang-user-manager/internal/auth/usecase"
)

type AuthHandler struct {
	Register usecase.RegisterUserUseCase
	Login    usecase.LoginUserUseCase
	Refresh  usecase.RefreshTokenUseCase
}

func NewAuthHandler(reg usecase.RegisterUserUseCase, log usecase.LoginUserUseCase, ref usecase.RefreshTokenUseCase) *AuthHandler {
	return &AuthHandler{Register: reg, Login: log, Refresh: ref}
}

func (h *AuthHandler) RegisterRoutes(r fiber.Router) {
	r.Post("/auth/register", h.register)
	r.Post("/auth/login", h.login)
	r.Post("/auth/refresh", h.refresh)
}

func (h *AuthHandler) register(c *fiber.Ctx) error {
	var dto usecase.RegisterUserDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	u, err := h.Register.Execute(dto)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(u)
}

func (h *AuthHandler) login(c *fiber.Ctx) error {
	var dto usecase.LoginDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	tok, err := h.Login.Execute(dto)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"token": tok, "expires_in": int((24 * time.Hour).Seconds())})
}

func (h *AuthHandler) refresh(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return c.Status(400).JSON(fiber.Map{"error": "missing token"})
	}
	tok, err := h.Refresh.Execute(token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"token": tok})
}
