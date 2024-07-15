package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	gcp "scheduler/gcp"
	"scheduler/utils"

	"github.com/labstack/echo/v4"
)

const port = ":5432"
const gcp_res = "Action sent to gcp, may take a few minutes to apply"

type handler struct {
	ch chan []string
}

func (h *handler) execCmd(c echo.Context) error {

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
	h.ch <- values

	return c.JSON(http.StatusOK, utils.S200(gcp_res))

}

func (h *handler) health(c echo.Context) error {
	log.Println("Everything is fine ...")
	return c.JSON(http.StatusOK, utils.S200("up"))
}

func main() {
	ch := make(chan []string, 1)

	h := &handler{ch: ch}
	e := echo.New()

	e.GET("/:action/:inst_name", h.execCmd)
	e.GET("/:action/:inst_name/:resize", h.execCmd)
	e.GET("/health", h.health)

	go func() {
		err := e.Start(port)
		// Close the chan once e.Start returns.
		close(ch)

		if err != http.ErrServerClosed {
			e.Logger.Fatal(err)
		}
	}()

	ctx, _ := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	// This goroutine listens for the signals and shutdown the echo server.
	go func() {
		<-ctx.Done()

		shutDownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := e.Shutdown(shutDownCtx); err != nil {
			e.Logger.Error(err)
		}
	}()

	for cmd := range ch {
		gcp.CallGCP(cmd)
	}
}
