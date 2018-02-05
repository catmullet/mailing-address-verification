package main

import (
	"github.com/labstack/echo"
	"os"
	"International-Address-Validation/api/addresses"
)

	func main() {

		// setup echo
		server := echo.New()

		addresses.Routes(server)

		// start the server (echo 3 now comes with built in graceful shutdown).
		server.Logger.Fatal(server.Start(":" + os.Getenv("PORT")))

	}
