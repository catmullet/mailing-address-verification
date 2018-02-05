package main

import (
	"github.com/labstack/echo"
	"International-Address-Validation/api/addresses"
)

	func main() {

		// setup echo
		server := echo.New()

		addresses.Routes(server)

		server.Logger.Fatal(server.Start(":6001"))

	}
