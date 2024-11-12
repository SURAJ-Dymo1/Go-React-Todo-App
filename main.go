package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hello World........suraj")
	app := fiber.New()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	PORT := os.Getenv("PORT")
	todos := []Todo{}

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(201).JSON(todos)
	})
	app.Get("/api/todo/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for _, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				return c.Status(201).JSON(todo)

			}
		}
		return c.Status(404).JSON(fiber.Map{"msg": "Something Went Wrong..."})

	})
	app.Post("/api/todos", func(c *fiber.Ctx) error {

		todo := &Todo{}
		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"err": "Todo body is require"})
		}
		todo.ID = len(todos) + 1

		todos = append(todos, *todo)
		return c.Status(201).JSON(todo)
	})
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = !todos[i].Completed
				return c.Status(200).JSON(todos[i])

			}

		}
		return c.Status(404).JSON(fiber.Map{"msg": "Something Went Wrong..."})
	})
	app.Delete("/api/todo/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(404).JSON(fiber.Map{"success": "Todo Deleted Successfuly.."})

			}

		}
		return c.Status(404).JSON(fiber.Map{"msg": "Something Went Wrong..."})

	})
	log.Fatal(app.Listen(":" + PORT))
}
