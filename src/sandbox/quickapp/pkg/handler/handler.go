package handler

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

func PostHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func PutHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func DeleteHandler(c *fiber.Ctx, db *sql.DB) error {
	return c.SendString("Hello")
}

func IndexHandler(c *fiber.Ctx, db *sql.DB) error {
	var res string
	var todos []string
	rows, err := db.Query("SELECT * FROM todos")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
	for rows.Next() {
		rows.Scan(&res)
		todos = append(todos, res)
	}
	return c.Render("index", fiber.Map{
		"Todos": todos,
	})
}
