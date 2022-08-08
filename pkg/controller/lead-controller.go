package controller

import (
	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite

	"goproject.com/pkg/config"
	model "goproject.com/pkg/models"
)

// GetLeads exported
func GetLeads(c *fiber.Ctx) {
	db := config.DBConn
	var leads []model.Lead

	db.Find(&leads); c.JSON(leads)
}

// GetLead exported
func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := config.DBConn
	var lead model.Lead
	
	db.Find(&lead,id); c.JSON(lead)
}

// NewLead exported
func NewLead(c *fiber.Ctx) {
	db := config.DBConn
	lead := new(model.Lead)
	
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err); return }

	db.Create(&lead); c.JSON(lead)
}

// UpdateLead exported
func UpdateLead(c *fiber.Ctx) {
	db := config.DBConn
	
	id := c.Params("id"); var lead model.Lead
	db.First(&lead, id)
	
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err); return }

	db.Create(&lead); c.JSON(lead)
}

// DeleteLead exported
func DeleteLead(c *fiber.Ctx) {
	db := config.DBConn
	id := c.Params("id"); var lead model.Lead

	db.First(&lead, id)
	
	if lead.Name == "" { c.Status(500).Send("No lead found with ID") }

	db.Delete(&lead);
	c.Send("Lead Succesfully Deleted")
}

