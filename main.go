package main

import (
	"fmt"

	 "github.com/gofiber/fiber/v2"
	 "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	
	app := fiber.New()

	app.Use(cors.New())

	app.Static("/", "./Frontend/dist")

	app.Get("/users", func(c *fiber.Ctx ) error {
	  	return c.JSON(&fiber.Map{
			"data": "ususarios desde el backend",
		})
	})

	app.Listen(":3000")
	fmt.Println("Server on port 3000")
}


//El Go-React-Crud.exe se crea con go build .