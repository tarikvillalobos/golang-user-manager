package handler

import (
	"github.com/gofiber/fiber/v2"
	"golang-user-manager/internal/plan/usecase"
)

type PlanHandler struct{ List usecase.ListPlansUseCase }

func NewPlanHandler(l usecase.ListPlansUseCase) *PlanHandler { return &PlanHandler{List: l} }

func (h *PlanHandler) RegisterRoutes(r fiber.Router) {
	r.Get("/plans", h.list)
}

func (h *PlanHandler) list(c *fiber.Ctx) error {
	res, err := h.List.Execute()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}
