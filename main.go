package main

import (
	"github.com/Shunux-Stuxnet/Blog_CRUD/controllers/blogcontroller"
	"github.com/Shunux-Stuxnet/Blog_CRUD/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	models.Connection()
	app := fiber.New()

	api := app.Group("/api")
	blog := api.Group("/blog")

	blog.Get("/", blogcontroller.Index)
	blog.Post("/", blogcontroller.Create)
	blog.Get("/:id", blogcontroller.Show)
	blog.Put("/:id", blogcontroller.Update)
	blog.Delete("/:id", blogcontroller.Delete)

	app.Listen(":8082")

}
