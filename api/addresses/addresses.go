package addresses

import (
	"github.com/labstack/echo"
	"International-Address-Validation/app/addresses"
	"github.com/kyani-inc/kms-api-shop/src/api/app"
)

func VerifyAddress(ctx echo.Context) error{

	country := ctx.Param("country")
	mAddr := ctx.Param("address")

	isValid := addresses.CheckMailingAddress(mAddr, country)

	if isValid {
		return app.OK(ctx, "Valid")
	}

	return app.OK(ctx, "Not Valid Address")

}