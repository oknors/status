package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

var token *string

func toInt(b []byte) (i int) {
	i, err := strconv.Atoi(string(b))
	if err != nil {
		log.Println("Failed to convert string to int")
	}
	return
}

func toFloat(b []byte) (f float64) {
	f, err := strconv.ParseFloat(string(b), 64)
	if err != nil {
		log.Println("Failed to convert string to float")
	}
	return
}

func main() {
	port := flag.String("p", ":8080", "http service address")
	flag.Parse()

	app := fiber.New()

	app.Get("/system", func(c *fiber.Ctx) error {
		return c.JSON(system("/"))
	})

	app.Get("/service/:service/:command", func(c *fiber.Ctx) error {
		return service(c, c.Params("service"), c.Params("command"))
	})
	log.Fatal(app.Listen(*port))
}
