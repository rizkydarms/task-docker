package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	godotenv.Load()
	// Connect to the database
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Create Database Connection failed: ", err)
		return
	}

	// Test database connection
	fmt.Println("Connecting to:", connStr)
	if err := db.Ping(); err != nil {
		fmt.Printf("Connection to database failed (DB_HOST: %s): %s\n", dbHost, err)
	} else {
		fmt.Println("Successfully connected to the database.")
	}

	// Create Echo instance
	e := echo.New()

	// Middleware for logging and recovering
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Additional route for quote page
	e.GET("/", func(c echo.Context) error {
		fmt.Println("Someone hit me!") // Log to console
		return c.HTML(http.StatusOK, `<h1>“Do not let the noise of others' opinions drown out your own inner voice. Have the courage to follow your heart and intuition; they somehow already know what you truly want to become - Steve Jobs”</h1>`)
	})
	e.GET("/connect", func(c echo.Context) error {
		fmt.Println("Someone hit me!") // Log to console
		return c.HTML(http.StatusOK, `<h1>“Do not let the noise of others' opinions drown out your own inner voice. Have the courage to follow your heart and intuition; they somehow already know what you truly want to become - Steve Jobs”</h1>`)
	})

	// Get port from environment variable or default to 8080
	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	// Start the server on the specified port
	fmt.Println("Starting Echo server on port:", httpPort)
	e.Logger.Fatal(e.Start(":" + httpPort))
}
