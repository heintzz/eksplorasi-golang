package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

const APP_PORT = ":4444"
const UPLOAD_PATH = "public/upload"

func main() {
	router := fiber.New()

	router.Post("/", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		location := fmt.Sprintf("%s/%s", UPLOAD_PATH, username)

		err = os.MkdirAll(location, os.ModePerm)
		if err != nil {
			return err
		}

		// CARA PERTAMA
		// siapin dulu pathnya yang mau diisi sama file
		// destination, err := os.Create(fmt.Sprintf("%s/%s", location, file.Filename))
		// if err != nil {
		// 	return err
		// }
		// defer destination.Close()

		// buka file yang dari request
		// f, err := file.Open()
		// if err != nil {
		// 	return err
		// }
		// defer f.Close()

		// copy isi dari file ke path yang udah dibikin tadi
		// if _, err := io.Copy(destination, f); err != nil {
		// 	return err
		// }

		// CARA KEDUA
		// disingkat pake SaveFile doang
		c.SaveFile(file, fmt.Sprintf("%s/%s", location, file.Filename))

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"message": "file upload successfully",
		})
	})

	router.Listen(APP_PORT)
}
