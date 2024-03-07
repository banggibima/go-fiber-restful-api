package app

import (
	"github.com/banggibima/go-fiber-restful-api/internal/transport/rest"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	UserHandler *rest.UserHandler
}

func NewApp(userHandler *rest.UserHandler) *App {
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
