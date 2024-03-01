package main

import (
	"github.com/gofiber/fiber/v2"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts []Post

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	app.Static("/", "./static")

	// Define a route to handle GET requests to "/"
	app.Get("/", func(c *fiber.Ctx) error {
		// Send the HTML form as a response
		// return c.SendString("Hello, World! This is a GET request.")
		return c.SendFile("./static/post_form.html")
	})

	// Define a route to handle POST requests to "/posts"
	app.Post("/posts", func(c *fiber.Ctx) error {
		// Parse request body into a new Post struct
		var newPost Post
		if err := c.BodyParser(&newPost); err != nil {
			return err
		}

		// Assign a unique ID to the new post
		newPost.ID = len(posts) + 1

		// Add the new post to the posts slice
		posts = append(posts, newPost)

		// Return the newly created post as JSON response
		return c.JSON(newPost)
	})

	// Start the server on port 3000
	app.Listen(":3000")
}
