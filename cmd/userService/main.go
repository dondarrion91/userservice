package main

import (
	"fmt"
	"os"
	"userservice/cmd/routes"
	database "userservice/internal/userService/database"

	"github.com/labstack/echo/v4"
)

func main() {
	router := echo.New()

	db := database.Init()

	r := routes.Routes(router, db)

	PORT := os.Getenv("PORT")
	port := fmt.Sprintf(":%s", PORT)

	r.Logger.Fatal(r.Start(port))
}
