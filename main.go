package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {

	root := "./www"

	app := fiber.New(fiber.Config{
		Views: html.New(root, ".html"),
	})
	app.Static("/", root)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	app.Get("/home", func(c *fiber.Ctx) error {
		return c.Render("home", nil)
	})
	err := app.Listen(":3003")
	if err != nil {
		fmt.Println(err)
	}
}
