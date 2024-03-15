package handlers

import (
	"net/http"

	"github.com/banggibima/go-fiber-restful-api/internal/entities"
	"github.com/banggibima/go-fiber-restful-api/internal/usecases"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserUseCase *usecases.UserUseCase
}

func NewUserHandler(userUseCase *usecases.UserUseCase) *UserHandler {
	return &UserHandler{UserUseCase: userUseCase}
}

func (h *UserHandler) GetUsersHandler(c *fiber.Ctx) error {
	users, err := h.UserUseCase.GetUsers()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(users)
}

func (h *UserHandler) GetUserByIDHandler(c *fiber.Ctx) error {
	idParam := c.Params("id")

	user, err := h.UserUseCase.GetUserByID(idParam)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

func (h *UserHandler) CreateUserHandler(c *fiber.Ctx) error {
	var input entities.User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.UserUseCase.CreateUser(input)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(user)
}

func (h *UserHandler) UpdateUserHandler(c *fiber.Ctx) error {
	idParam := c.Params("id")

	var input entities.User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.UserUseCase.UpdateUser(idParam, input)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

func (h *UserHandler) DeleteUserHandler(c *fiber.Ctx) error {
	idParam := c.Params("id")

	if err := h.UserUseCase.DeleteUser(idParam); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(http.StatusNoContent)
}
