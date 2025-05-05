package main

import (
	"fmt"
	"net/http"

	"golang/database"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.Status(http.StatusOK).SendString("hello")
	})

	app.Get("/register", func(c fiber.Ctx) error {
		db, err := database.Connection()
		if err != nil {
			panic(fmt.Sprintf("error: %v", err))
		}

		email := c.FormValue("email")
		password := c.FormValue("password")
		name := c.FormValue("name")

		var user database.User

		if err := db.Unscoped().First(&user, "email = ?", email); err == nil {
			return c.Status(http.StatusBadRequest).JSON(Response{
				Message: fmt.Sprintf("user with email %v alredy exist!", email),
			})
		}

		if err := db.Unscoped().Model(&user).Create(database.User{
			Name:     name,
			Password: password,
			Email:    email,
		}); err != nil {
			return c.Status(http.StatusBadGateway).JSON(Response{
				Message: fmt.Sprintf("error at register %v", err),
			})
		}

		return c.Status(http.StatusOK).JSON(Response{
			Message: "register complete",
		})
	})

	fmt.Print("server running at http://localhost:30000")
	app.Listen(":3000")
}
