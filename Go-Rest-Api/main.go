package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func (r *Repository) CreatBooks(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	r.DB.Create(&book)
	return c.JSON(book)
}

type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_books", r.CreatBooks)
	api.Get("/get_books", r.GetBooks)
	api.Get("/get_books/:id", r.GetBook)
	api.Put("/update_books/:id", r.UpdateBook)
	api.Delete("/delete_books/:id", r.DeleteBook)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := storage.NewConnection(config)
	app := fiber.New()
	r := Repository{
		DB: db,
	}
	r.SetupRoutes(app)
	app.Listen(":3000")
}
