package route

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"goproject.com/pkg/config"
	"goproject.com/pkg/controller"
)

// HRMroutes exported
func HRMroutes() {
	if err := config.MongoConnect(); err != nil {
		log.Fatal(err)
	}
	
	app := fiber.New()
	
	hrmSetup(app)

	log.Fatal(app.Listen(":3000"))
}

func hrmSetup(app *fiber.App) {
	app.Get("/employee", 		controller.GetEmployees)
	app.Get("/employee/:id", 	controller.GetEmployee)

	app.Post("/employee", controller.NewEmployee)

	app.Put("/employee/:id", 	controller.UpdateEmployee)
	
	app.Delete("/employee/:id", controller.DeleteEmployee)
}