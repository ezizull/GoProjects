package route

import (
	"fmt"

	"goproject.com/pkg/config"
	"goproject.com/pkg/controller"
	model "goproject.com/pkg/models"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// CRMroutes exported
func CRMroutes(){
	app := fiber.New(); initDatabase()

	crmSetup(app)
	
	app.Listen(3000)
	
	defer config.DBConn.Close()
}

func crmSetup(app *fiber.App) {
	app.Get("/lead", 		controller.GetLeads)
	app.Get("/lead/:id",	controller.GetLead)

	app.Post("/lead", 		controller.NewLead)
	app.Post("/lead/:id", 	controller.UpdateLead)
	
	app.Delete("/lead/:id", 	controller.DeleteLead)
}

func initDatabase(){
	var err error
	config.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil { panic("failed to connect database") }

	fmt.Println("connection opened to database")
	
	config.DBConn.AutoMigrate(&model.Lead{})
	fmt.Println("database migrate")
}	