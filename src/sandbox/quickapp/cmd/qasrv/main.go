package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	h "github.com/jocarde/quickapp/pkg/handler"
	pg "github.com/jocarde/quickapp/pkg/postgres"
)

func main() {
	db := pg.GetDB()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return h.IndexHandler(c, db)
	})

	app.Post("/", func(c *fiber.Ctx) error {
		return h.PostHandler(c, db)
	})

	app.Put("/update", func(c *fiber.Ctx) error {
		return h.PutHandler(c, db)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
		return h.DeleteHandler(c, db)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
