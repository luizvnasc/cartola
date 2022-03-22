package server_test

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/luizvnasc/cartola/mercado-producer/internal/server"
)

func TestHandlerVersao(t *testing.T) {

	app := fiber.New()

	app.Get("/versao", server.Versao)

	resp, err := app.Test(httptest.NewRequest("GET", "/versao", nil))
	utils.AssertEqual(t, nil, err, "app.Test")
	body, err := ioutil.ReadAll(resp.Body)
	utils.AssertEqual(t, nil, err, "ioutil.ReadAll")
	utils.AssertEqual(t, "v0.0.1", string(body), "request Body")

}
