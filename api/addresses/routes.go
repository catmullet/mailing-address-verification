package addresses

import "github.com/labstack/echo"


func Routes(e *echo.Echo) {

	// Languages Group
	g := e.Group("/address/verify/:country/:address")

	// Languages Group Handlers
	g.GET("", VerifyAddress)

}
