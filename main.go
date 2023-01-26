package main

import (
	"log"

	"github.com/fernandovmedina/primer-api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Declaramos una app de fiber
	var fiberApp = fiber.New()

	// Ruta de la funcion "GetBook"
	fiberApp.Get("/api/book/", routes.GetBooks)

	// Ruta de la funcion "GetBooks"
	fiberApp.Get("/api/book/:id", routes.GetBook)

	// Ruta de la funcion "PostBook"
	fiberApp.Post("/api/book/", routes.PostBook)

	// Ruta de la funcion "PutBook"
	fiberApp.Put("/api/book/:id", routes.PutBook)

	// Ruta de la funcion "DeleteBook"
	fiberApp.Delete("/api/book/:id", routes.DeleteBook)

	// Corremos la app en el puerto 3000
	log.Fatal(fiberApp.Listen(":3000"))
}
