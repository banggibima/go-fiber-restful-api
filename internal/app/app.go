package app

import (
	"github.com/banggibima/go-fiber-restful-api/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	UserHandler *handlers.UserHandler
}

func NewApp(userHandler *handlers.UserHandler) *App {
	return &App{UserHandler: userHandler}
}

func (a *App) SetupRoutes(app *fiber.App) {
	v1 := app.Group("/v1")
	{
		users := v1.Group("/users")
		{
			users.Get("/", a.UserHandler.GetUsersHandler)
			users.Get("/:id", a.UserHandler.GetUserByIDHandler)
			users.Post("/", a.UserHandler.CreateUserHandler)
			users.Put("/:id", a.UserHandler.UpdateUserHandler)
			users.Delete("/:id", a.UserHandler.DeleteUserHandler)
		}
	}
}
