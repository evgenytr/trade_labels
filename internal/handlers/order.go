package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"

	"github.com/evgenytr/trade_labels.git/internal/config"
	"github.com/evgenytr/trade_labels.git/internal/domain"
)

func PostOrderHandler(c echo.Context) error {

	if !IsJSONContentTypeCorrect(c.Request()) {
		c.Response().WriteHeader(http.StatusBadRequest)
		return domain.ErrorJSONTypeIncorrect
	}

	var order domain.Order

	var buf bytes.Buffer

	_, err := buf.ReadFrom(c.Request().Body)

	if err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		fmt.Printf("error reading body in post order request %v\n", err)
		return fmt.Errorf("error reading body in post order request %w", err)

	}

	c.Request().Body = io.NopCloser(bytes.NewBuffer(buf.Bytes()))

	d := json.NewDecoder(c.Request().Body)

	if err = d.Decode(&order); err != nil {
		c.Response().WriteHeader(http.StatusBadRequest)
		fmt.Printf("error decoding body in post order request %v\n", err)
		return fmt.Errorf("error decoding body in post order request %w", err)
	}

	fmt.Println(order)
	//TODO: store order

	//post to indesign
	//TODO: separate it from this handler

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(order).
		Post(fmt.Sprintf("%v/label", config.IndesignHost))

	if err != nil {
		fmt.Printf("indesign didn't respond %v\n", err)
		return fmt.Errorf("indesign didn't respond %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		fmt.Printf("indesign response has wrong status %v\n", resp.StatusCode())
		return fmt.Errorf("indesign response has wrong status %v", resp.StatusCode())
	}

	fmt.Printf("indesign response %v\n", resp)

	return nil
}

func IsJSONContentTypeCorrect(r *http.Request) bool {
	if len(r.Header.Values("Content-Type")) == 0 {
		return false
	}

	for contentTypeCurrentIndex, contentType := range r.Header.Values("Content-Type") {
		if contentType == "application/json" {
			break
		}
		if contentTypeCurrentIndex == len(r.Header.Values("Content-Type"))-1 {
			return false
		}
	}

	return true
}
