package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	ua "golang-user-manager/internal/user/usecase/auth"
	"golang-user-manager/internal/user/usecase/crud"
)

type UserHandler struct {
	List           crud.ListUsersUseCase
	Get            crud.GetUserByIDUseCase
	Create         crud.CreateUserUseCase
	Update         crud.UpdateUserUseCase
	Delete         crud.DeleteUserUseCase
	UpdatePassword ua.UpdatePasswordUseCase
}

func NewUserHandler(l crud.ListUsersUseCase, g crud.GetUserByIDUseCase, c crud.CreateUserUseCase, u crud.UpdateUserUseCase, d crud.DeleteUserUseCase, up ua.UpdatePasswordUseCase) *UserHandler {
	return &UserHandler{List: l, Get: g, Create: c, Update: u, Delete: d, UpdatePassword: up}
}

func (h *UserHandler) RegisterRoutes(r fiber.Router) {
	r.Get("/users", h.list)
	r.Get("/users/:id", h.getByID)
	r.Post("/users", h.create)
	r.Put("/users/:id", h.update)
	r.Delete("/users/:id", h.delete)
	r.Put("/users/:id/password", h.updatePassword)
}

func (h *UserHandler) list(c *fiber.Ctx) error {
	res, err := h.List.Execute()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}

func (h *UserHandler) getByID(c *fiber.Ctx) error {
	id64, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	u, err := h.Get.Execute(uint(id64))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(u)
}

func (h *UserHandler) create(c *fiber.Ctx) error {
	var dto crud.CreateUserDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	u, err := h.Create.Execute(dto)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(u)
}

func (h *UserHandler) update(c *fiber.Ctx) error {
	id64, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	var dto crud.UpdateUserDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	dto.ID = uint(id64)
	u, err := h.Update.Execute(dto)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(u)
}

func (h *UserHandler) delete(c *fiber.Ctx) error {
	id64, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	if err := h.Delete.Execute(uint(id64)); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}

func (h *UserHandler) updatePassword(c *fiber.Ctx) error {
	id64, _ := strconv.ParseUint(c.Params("id"), 10, 32)
	var body struct {
		Password string `json:"password"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.UpdatePassword.Execute(uint(id64), body.Password); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(204)
}
