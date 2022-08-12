package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	//"github.com/faztweb/go-react-crud/models"
	"github.com/JulianElisii/Go-React-Crud/models"
)

func main() {

	app := fiber.New()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://127.0.0.1:27017/Go-DATABASE"))

	if err != nil {
		panic(err)
	}

	app.Use(cors.New())

	app.Static("/", "./Frontend/dist")

	app.Post("/users", func(c *fiber.Ctx) error {
		var user models.User
		c.BodyParser(&user)

		coll := client.Database("Go-DATABASE").Collection("users")
		result, err := coll.InsertOne(context.TODO(), bson.D{{
			Key:   "name",
			Value: user.Name,
		}})

		if err != nil {
			panic(err)
		}

		return c.JSON(&fiber.Map{
			"data": result,
		})

	})

	app.Get("/users", func(c *fiber.Ctx) error {
		var users []models.User

		coll := client.Database("Go-DATABASE").Collection("users")
		results, error := coll.Find(context.TODO(), bson.M{})

		if error != nil {
			panic(error)
		}

		for results.Next(context.TODO()) {
			var user models.User
			results.Decode(&user)
			users = append(users, user)
		}

		return c.JSON(&fiber.Map{
			"users": users,
		})

	})

	app.Listen(":3000")
	fmt.Println("Server on port 3000")
}

//El Go-React-Crud.exe se crea con go build .
