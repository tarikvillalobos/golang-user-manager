package handler

import (
	"github.com/gofiber/fiber/v2"
	sd "golang-user-manager/internal/subscription/domain"
	"golang-user-manager/internal/subscription/usecase"
)

type SubscriptionHandler struct {
	Create usecase.CreateSubscriptionUseCase
}

func NewSubscriptionHandler(c usecase.CreateSubscriptionUseCase) *SubscriptionHandler {
	return &SubscriptionHandler{Create: c}
}

func (h *SubscriptionHandler) RegisterRoutes(r fiber.Router) {
	r.Post("/subscriptions", h.create)
}

func (h *SubscriptionHandler) create(c *fiber.Ctx) error {
	var s sd.Subscription
	if err := c.BodyParser(&s); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.Create.Execute(&s); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(s)
}
