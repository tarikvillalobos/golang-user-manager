package handler

import (
	"github.com/gofiber/fiber/v2"
	ua "golang-user-manager/internal/user/usecase/auth"
)

type ProfileHandler struct{ GetProfile ua.GetProfileUseCase }

func NewProfileHandler(g ua.GetProfileUseCase) *ProfileHandler { return &ProfileHandler{GetProfile: g} }

func (h *ProfileHandler) RegisterRoutes(r fiber.Router) {
	r.Get("/profile", h.me)
}

func (h *ProfileHandler) me(c *fiber.Ctx) error {
	// In a real app, extract userID from JWT. For demo, query ?user_id=1
	uid := c.QueryInt("user_id", 0)
	if uid == 0 {
		return c.Status(400).JSON(fiber.Map{"error": "missing user_id"})
	}
	u, err := h.GetProfile.Execute(uint(uid))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(u)
}
