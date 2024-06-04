package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"

	"github.com/evgenytr/trade_labels.git/internal/config"
)

func PingHandler(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "text/plain")
	return nil
}

func IndesignPingHandler(c echo.Context) error {
	client := resty.New()
	resp, err := client.R().
		Get(config.IndesignHost + "/ping")

	fmt.Println(err)
	if err != nil {
		return fmt.Errorf("indesign didn't respond %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("indesign response has wrong status %v", resp.StatusCode())
	}

	fmt.Println(resp)
	return nil
}
