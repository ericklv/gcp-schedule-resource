package main

import (
	"log"
	"net/http"

	gcp "scheduler/gcp"
	"scheduler/utils"

	"github.com/labstack/echo/v4"
)

const port = ":5432"
const gcp_res = "Action applied"
const gcp_err = "Action cannot be applied"

func execCmd(c echo.Context) error {

	params := new(utils.Params)
	if err := c.Bind(params); err != nil {
		return err
	}

	instance := params.Instance
	values := gcp.Action(*params)

	if values == nil {
		return c.JSON(http.StatusBadRequest, utils.S4xx("Invalid action"))
	}

	values = append(values, instance)
	_, err := gcp.CallGCP(values)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.S4xx(gcp_err))
	}
	return c.JSON(http.StatusOK, utils.S200(gcp_res))

}

func health(c echo.Context) error {
	log.Println("Everything is fine ...")
	return c.JSON(http.StatusOK, utils.S200("up"))
}

func main() {
	e := echo.New()

	e.GET("/:action/:inst_name", execCmd)
	e.GET("/:action/:inst_name/:resize", execCmd)
	e.GET("/health", health)

	err := e.Start(port)

	if err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}
}
