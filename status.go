package main

import (
	"flag"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	port := flag.String("p", ":11111", "http service address")
	flag.Parse()

	app := fiber.New(fiber.Config{
		Prefork:       true,
		StrictRouting: true,
		ServerHeader:  "OKNO",
		AppName:       "OKNO status",
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PATCH",
	}))
	app.Use(logger.New())
	app.Get("/system", func(c *fiber.Ctx) error {
		return c.JSON(system("/"))
	})

	app.Get("/service/:service/:command", func(c *fiber.Ctx) error {
		return service(c, c.Params("service"), c.Params("command"))
	})

	log.Fatal(app.Listen(*port))
}
