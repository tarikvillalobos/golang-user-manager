package handler

import (
	"github.com/gofiber/fiber/v2"
	pd "golang-user-manager/internal/profile/domain"
	"golang-user-manager/internal/profile/usecase"
)

type ProfileHandler struct{ Update usecase.UpdateProfileUseCase }

func NewProfileHandlerProfile(u usecase.UpdateProfileUseCase) *ProfileHandler {
	return &ProfileHandler{Update: u}
}

func (h *ProfileHandler) RegisterRoutes(r fiber.Router) {
	r.Put("/profile", h.update)
}

func (h *ProfileHandler) update(c *fiber.Ctx) error {
	var p pd.Profile
	if err := c.BodyParser(&p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.Update.Execute(&p); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
