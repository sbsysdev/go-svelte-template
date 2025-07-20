package infrastructure

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sbsysdev/go-svelte-template/internal/adapters/controllers"
	"github.com/sbsysdev/go-svelte-template/internal/adapters/gateways"
	"github.com/sbsysdev/go-svelte-template/internal/adapters/presenters"
	"github.com/sbsysdev/go-svelte-template/internal/application"
)

type Api struct {
	env    *Environment
	dbpool *pgxpool.Pool
	app    *fiber.App
}

func (api *Api) StartApiServer() error {
	v1Group := api.app.Group("/api/v1")

	// Speciality Repository
	specialityRepository := gateways.NewSpecialityRepository(api.dbpool)

	// List Specialities
	listSpecialityPresenter := presenters.NewListSpecialityPresenter()
	listSpecialityUseCase := application.NewListSpecialityUseCase(specialityRepository, listSpecialityPresenter)
	listSpecialityController := controllers.NewListSpecialityController(listSpecialityUseCase)
	v1Group.Get("/specialities", listSpecialityController.Handle)

	createSpecialityPresenter := presenters.NewCreateSpecialityPresenter()
	createSpecialityUseCase := application.NewCreateSpecialityUseCase(specialityRepository, createSpecialityPresenter)
	createSpecialityController := controllers.NewCreateSpecialityController(createSpecialityUseCase)
	v1Group.Post("/specialities", createSpecialityController.Handle)

	// Start the server
	return api.app.Listen(fmt.Sprintf(":%s", api.env.APP_PORT))
}

func NewApiServer(env *Environment, dbpool *pgxpool.Pool) *Api {
	app := fiber.New(fiber.Config{
		AppName: env.APP_NAME,
	})

	app.Use(recover.New(), cors.New(cors.Config{
		AllowOrigins:     env.APP_ORIGIN,
		AllowCredentials: true,
	}), helmet.New(), limiter.New(), favicon.New())

	if env.APP_ENV == "dev" {
		// TODO: cusom logger when production
		app.Use(logger.New())
	}

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("fiberContext", c)
		return c.Next()
	})

	return &Api{
		env:    env,
		dbpool: dbpool,
		app:    app,
	}
}
