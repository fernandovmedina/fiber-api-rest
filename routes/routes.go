package routes

import (
	"github.com/fernandovmedina/primer-api/config"
	"github.com/fernandovmedina/primer-api/models"
	"github.com/gofiber/fiber/v2"
)

func GetBook(c *fiber.Ctx) error {
	var id = c.Params("id")

	var book models.Book

	resultado := config.Database.Find(&book, id)

	if resultado.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  404,
			"mensaje": "Libro no encontrado o registrado",
		})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"status":  200,
			"mensaje": "Libro encontrado exitosamente",
			"dato":    book,
		})
	}
}

func GetBooks(c *fiber.Ctx) error {
	var books []models.Book

	config.Database.Model(&models.Book{}).Order("ID asc").Find(&books)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  200,
		"mensaje": "Libros registrados en la base de datos",
		"datos":   books,
	})
}

func PostBook(c *fiber.Ctx) error {
	var book = new(models.Book)

	if err := c.BodyParser(book); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"status":  503,
			"mensaje": "No se pudo crear el libro",
		})
	}

	config.Database.Create(&book)

	return c.Status(201).JSON(fiber.Map{
		"status":  201,
		"mensaje": "Libro creado exitosamente",
		"dato":    book,
	})
}

func PutBook(c *fiber.Ctx) error {
	var id = c.Params("id")

	var book = new(models.Book)

	if err := c.BodyParser(&book); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"status":  503,
			"mensaje": "Libro no encontrado",
		})
	} else {
		config.Database.Where("id = ?", id).Updates(&book)
		return c.Status(200).JSON(fiber.Map{
			"status":  200,
			"mensaje": "Libro actualizado correctamente",
		})
	}
}

func DeleteBook(c *fiber.Ctx) error {
	var id = c.Params("id")

	var book models.Book

	resultado := config.Database.Delete(&book, id)

	if resultado.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"status":  404,
			"mensaje": "Libro no encontrado",
		})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"status":  200,
			"mensaje": "Libro borrado exitosamente",
		})
	}
}
