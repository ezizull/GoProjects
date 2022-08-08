package controller

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"goproject.com/pkg/config"
	model "goproject.com/pkg/models"
)

// GetEmployees exported
func GetEmployees(c *fiber.Ctx) error { 
	query := bson.D{{}} // empty json in mongo

	cursor, err := config.Mongo.Db.Collection("employees").Find(c.Context(), query)
	
	if err != nil { return c.Status(500).SendString(err.Error()) }

	var employees []model.Employee = make([]model.Employee, 0)
	
	if err := cursor.All(c.Context(), &employees); err != nil {
		return c.Status(500).SendString(err.Error()) 
	}

	return c.JSON(employees)
}

// GetEmployee exported
func GetEmployee(c *fiber.Ctx) error { 
	employeeID, err := primitive.ObjectIDFromHex(c.Params("id"))

	if err != nil { return c.Status(400).SendString(err.Error())  }

	employee := model.Employee{}

	query := bson.D{{Key: "_id", Value: employeeID}}
	cursor := config.Mongo.Db.Collection("employees").FindOne(c.Context(), query)
	cursor.Decode(&employee)

	if  err != nil {
		if err == mongo.ErrNoDocuments { return c.Status(400).SendString(err.Error())  }
		return c.Status(500).SendString(err.Error()) 
	}

	return c.Status(200).JSON(employee)
}

// NewEmployee exported
func NewEmployee(c *fiber.Ctx) error { 
	collection := config.Mongo.Db.Collection("employees")

	employee := new(model.Employee)

	if err := c.BodyParser(employee); err != nil { return c.Status(400).SendString(err.Error())  }

	employee.ID = ""

	insertionResult, err := collection.InsertOne(c.Context(), employee)

	if err != nil { return c.Status(500).SendString(err.Error()) }

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	createdEmployee := &model.Employee{}
	createdRecord.Decode(createdEmployee)

	return c.Status(201).JSON(createdEmployee)
}

// UpdateEmployee exported
func UpdateEmployee(c *fiber.Ctx) error { 
	employeeID, err := primitive.ObjectIDFromHex(c.Params("id"))

	if err != nil { return c.SendStatus(400) }

	employee := new(model.Employee)

	if err := c.BodyParser(employee); err != nil { return c.SendStatus(400) }

	query := bson.D{{Key: "_id", Value: employeeID}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: employee.Name},
				{Key: "age", Value: employee.Age},
				{Key: "salary", Value: employee.Salary},
			},
		},
	}

	err = config.Mongo.Db.Collection("employees").FindOneAndUpdate(c.Context(), query, update).Err()

	if err != nil {
	if err == mongo.ErrNoDocuments { return c.SendStatus(400)}
		return c.SendStatus(500)
	}

	employee.ID = c.Params("id")

	return c.Status(200).JSON(employee)
}

// DeleteEmployee exported
func DeleteEmployee(c *fiber.Ctx) error { 
	employeeID, err := primitive.ObjectIDFromHex(c.Params("id"))

	if err != nil { return c.SendStatus(400) }

	query := bson.D{{Key: "_id", Value: employeeID}}
	result, err := config.Mongo.Db.Collection("employees").DeleteOne(c.Context(), &query)

	if err != nil { return c.SendStatus(500) }

	if result.DeletedCount < 1 { return c.SendStatus(404) }

	return c.Status(200).JSON("record deleted")
}