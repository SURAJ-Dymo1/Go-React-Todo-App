package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

var collection *mongo.Collection

func main() {

	if os.Getenv("ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error Loading .env file")
		}
	}

	MONGO_URI := os.Getenv("MONGO_URI")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "5000"
	}
	clientOptions := options.Client().ApplyURI(MONGO_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB Atlas")
	collection = client.Database("Go_React_Todo_App").Collection("Todos")

	// Initialize Fiber app
	app := fiber.New()

	// Add CORS middleware
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "*", // Allow all origins (you can specify specific origins if needed)
	// 	AllowMethods: "GET,POST,PATCH,DELETE",
	// 	AllowHeaders: "Origin, Content-Type, Accept",
	// }))

	// Set up routes
	app.Get("/api/gettodos", GetTodos)
	app.Post("/api/createtodo", CreateTodo)
	app.Patch("/api/updatetodo/:id", UpdateTodo)
	app.Delete("/api/deletetodo/:id", DeleteTodo)

	if os.Getenv("ENV") == "production" {
		app.Static("/", "./client/dist")
	}
	// Start the server
	log.Fatal(app.Listen(":" + PORT))
}

// func GetTodo(c *fiber.Ctx) error {

// }

func GetTodos(c *fiber.Ctx) error {

	var todos []Todo

	cursor, err := collection.Find(context.Background(), bson.M{})
	defer cursor.Close(context.Background())

	if err != nil {

		return err
	}
	for cursor.Next(context.Background()) {
		var todo Todo
		if err := cursor.Decode(&todo); err != nil {

			return err
		}
		todos = append(todos, todo)
	}

	return c.Status(200).JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
	todo := new(Todo)

	if err := c.BodyParser(todo); err != nil {
		return nil
	}
	if todo.Body == "" {
		return c.Status(404).JSON(fiber.Map{"error": "Body Should not empty"})
	}
	insertResult, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return nil
	}
	todo.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(200).JSON(todo)

}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Id"})
	}
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"completed": true}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{"success": true})
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Id"})
	}
	filter := bson.M{"_id": objectID}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{"success": "Delete Successfully"})
}
