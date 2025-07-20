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

	// Specialty Repository
	specialtyRepository := gateways.NewSpecialtyRepository(api.dbpool)

	// Create Specialty
	createSpecialtyPresenter := presenters.NewCreateSpecialtyPresenter()
	createSpecialtyUseCase := application.NewCreateSpecialtyUseCase(specialtyRepository, createSpecialtyPresenter)
	createSpecialtyController := controllers.NewCreateSpecialtyController(createSpecialtyUseCase)
	v1Group.Post("/specialties", createSpecialtyController.Handle)

	// List Specialties
	listSpecialtyPresenter := presenters.NewListSpecialtyPresenter()
	listSpecialtyUseCase := application.NewListSpecialtyUseCase(specialtyRepository, listSpecialtyPresenter)
	listSpecialtyController := controllers.NewListSpecialtyController(listSpecialtyUseCase)
	v1Group.Get("/specialties", listSpecialtyController.Handle)

	// Doctor Repository
	doctorRepository := gateways.NewDoctorRepository(api.dbpool, specialtyRepository)

	// Create Doctor
	createDoctorPresenter := presenters.NewCreateDoctorPresenter()
	createDoctorUseCase := application.NewCreateDoctorUseCase(doctorRepository, createDoctorPresenter, specialtyRepository)
	createDoctorController := controllers.NewCreateDoctorController(createDoctorUseCase)
	v1Group.Post("/doctors", createDoctorController.Handle)

	// List Doctors by Specialty
	listDoctorBySpecialtyPresenter := presenters.NewListDoctorBySpecialtyPresenter()
	listDoctorBySpecialtyUseCase := application.NewListDoctorBySpecialtyUseCase(doctorRepository, listDoctorBySpecialtyPresenter)
	listDoctorBySpecialtyController := controllers.NewListDoctorBySpecialtyController(listDoctorBySpecialtyUseCase)
	v1Group.Get("/specialties/:specialtyID/doctors", listDoctorBySpecialtyController.Handle)

	// Patient Repository
	patientRepository := gateways.NewPatientRepository(api.dbpool)

	// Create Patient
	createPatientPresenter := presenters.NewCreatePatientPresenter()
	createPatientUseCase := application.NewCreatePatientUseCase(patientRepository, createPatientPresenter)
	createPatientController := controllers.NewCreatePatientController(createPatientUseCase)
	v1Group.Post("/patients", createPatientController.Handle)

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
