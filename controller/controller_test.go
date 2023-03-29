package controller

import (
	storage "localization/storage"
	"localization/utils"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"

	"github.com/stretchr/testify/assert"
)

func TestAppCreate(t *testing.T) {
	mode := storage.Storage{Utils: &utils.Utils{}}
	l := LocalizationController{Svc: mode}

	APP := fiber.New()

	APP.Post("/:app", func(c *fiber.Ctx) error {
		appName := c.Params("app")
		bod := c.Body()
		err := l.Svc.CreateApp(bod, appName)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"type":    "Fetch Data",
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(nil)
	})
	r := strings.NewReader("my request")
	req := httptest.NewRequest(http.MethodPost, "/app", r)
	resp, _ := APP.Test(req, 1)

	assert.Equal(t, 200, resp.StatusCode)

}
func TestAppRead(t *testing.T) {
	mode := storage.Storage{Utils: &utils.Utils{}}
	l := LocalizationController{Svc: mode}

	app := fiber.New()

	app.Get("/:app", func(c *fiber.Ctx) error {
		appName := c.Params("app")

		result, err := l.Svc.ReadApp(appName) //, app)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"type":    "Fetch Data",
				"message": err.Error(),
			})
		}

		res := Json(result)
		return c.JSON(res)
	})

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8081/app", nil)

	resp, _ := app.Test(req, 1)

	assert.Equal(t, 200, resp.StatusCode)

}

func TestAppUpdate(t *testing.T) {
	mode := storage.Storage{Utils: &utils.Utils{}}
	l := LocalizationController{Svc: mode}

	APP := fiber.New()

	APP.Put("/:app", func(c *fiber.Ctx) error {
		bod := c.Body()
		appName := c.Params("app")

		err := l.Svc.UpdateApp(appName, bod)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"type":    "Fetch Data",
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(err)
	})
	r := strings.NewReader("my request")
	req := httptest.NewRequest(http.MethodPut, "/app", r)
	resp, _ := APP.Test(req, 1)

	assert.Equal(t, 200, resp.StatusCode)

}
func TestAppDelete(t *testing.T) {
	mode := storage.Storage{Utils: &utils.Utils{}}
	l := LocalizationController{Svc: mode}

	APP := fiber.New()

	APP.Delete("/:app", func(c *fiber.Ctx) error {
		appName := c.Params("app")
		err := l.Svc.DeleteApp(appName)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"type":    "Fetch Data",
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(err)
	})
	r := strings.NewReader("my request")
	req := httptest.NewRequest(http.MethodDelete, "/app", r)
	resp, _ := APP.Test(req, 1)

	assert.Equal(t, 200, resp.StatusCode)
}

func TestModuleCreate(t *testing.T) {
	mode := storage.Storage{Utils: &utils.Utils{}}
	l := LocalizationController{Svc: mode}

	APP := fiber.New()

	APP.Post("/:app/:module", func(c *fiber.Ctx) error {
		module := c.Body()
		appName := c.Params("app")
		moduleName := c.Params("module")

		err := l.Svc.CreateModule(appName, moduleName, module)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"type":    "Fetch Data",
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(nil)
	})
	r := strings.NewReader("my request")
	req := httptest.NewRequest(http.MethodPost, "/app/module", r)
	resp, _ := APP.Test(req, 1)

	assert.Equal(t, 200, resp.StatusCode)

}
func TestModuleRead(t *testing.T) {
	mode := storage.Storage{Utils: &utils.Utils{}}
	l := LocalizationController{Svc: mode}

	app := fiber.New()

	app.Get("/:app/:module", func(c *fiber.Ctx) error {
		appName := c.Params("app")
		moduleName := c.Params("module")

		result, err := l.Svc.ReadModule(appName, moduleName)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"type":    "Fetch Data",
				"message": err.Error(),
			})
		}

		res := Json(*result)
		return c.JSON(res)
	})

	req := httptest.NewRequest(http.MethodGet, "/app/module", nil)

	resp, _ := app.Test(req, 1)

	assert.Equal(t, 200, resp.StatusCode)

}
func TestModuleUpdate(t *testing.T) {
	mode := storage.Storage{Utils: &utils.Utils{}}
	l := LocalizationController{Svc: mode}

	APP := fiber.New()

	APP.Put("/:app/:module", func(c *fiber.Ctx) error {
		module := c.Body()
		appName := c.Params("app")
		moduleName := c.Params("module")

		err := l.Svc.UpdateModule(appName, moduleName, module)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"type":    "Fetch Data",
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(err)
	})
	r := strings.NewReader("my request")
	req := httptest.NewRequest(http.MethodPut, "/app/module", r)
	resp, _ := APP.Test(req, 1)

	assert.Equal(t, 200, resp.StatusCode)

}
func TestModuleDelete(t *testing.T) {
	mode := storage.Storage{Utils: &utils.Utils{}}
	l := LocalizationController{Svc: mode}

	APP := fiber.New()

	APP.Delete("/:app/:module", func(c *fiber.Ctx) error {
		appName := c.Params("app")
		moduleName := c.Params("module")

		err := l.Svc.DeleteModule(appName, moduleName)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"type":    "Fetch Data",
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(err)
	})
	r := strings.NewReader("my request")
	req := httptest.NewRequest(http.MethodDelete, "/app/module", r)
	resp, _ := APP.Test(req, 1)

	assert.Equal(t, 200, resp.StatusCode)

}

func TestLanguageCreate(t *testing.T) {
	mode := storage.Storage{Utils: &utils.Utils{}}
	l := LocalizationController{Svc: mode}

	APP := fiber.New()

	APP.Post("/:app/:module/:language", func(c *fiber.Ctx) error {

		language := c.Body()
		appName := c.Params("app")
		moduleName := c.Params("module")
		languageName := c.Params("language")

		err := l.Svc.CreateLanguage(appName, moduleName, languageName, language)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"type":    "Fetch Data",
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(err)
	})
	r := strings.NewReader("my request")
	req := httptest.NewRequest(http.MethodPost, "/app/module/language", r)
	resp, _ := APP.Test(req, 1)

	assert.Equal(t, 200, resp.StatusCode)

}
func TestLanguageRead(t *testing.T) {

	mode := storage.Storage{Utils: &utils.Utils{}}
	l := LocalizationController{Svc: mode}

	app := fiber.New()

	app.Get("/:app/:module/:language", func(c *fiber.Ctx) error {
		appName := c.Params("app")
		moduleName := c.Params("module")
		languageName := c.Params("language")

		result, err := l.Svc.ReadLanguage(appName, moduleName, languageName)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"type":    "Fetch Data",
				"message": err.Error(),
			})
		}
		res := Json(result)
		return c.JSON(res)
	})

	req := httptest.NewRequest(http.MethodGet, "/app/module/language", nil)

	resp, _ := app.Test(req, 1)

	assert.Equal(t, 200, resp.StatusCode)
}
func TestLanguageUpdate(t *testing.T) {
	mode := storage.Storage{Utils: &utils.Utils{}}
	l := LocalizationController{Svc: mode}

	APP := fiber.New()

	APP.Put("/:app/:module/:language", func(c *fiber.Ctx) error {
		language := c.Body()
		appName := c.Params("app")
		moduleName := c.Params("module")
		languageName := c.Params("language")

		err := l.Svc.UpdateLanguage(appName, moduleName, languageName, language)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"type":    "Fetch Data",
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(err)
	})
	r := strings.NewReader("my request")
	req := httptest.NewRequest(http.MethodPut, "/app/module/language", r)
	resp, _ := APP.Test(req, 1)

	assert.Equal(t, 200, resp.StatusCode)

}
func TestLanguageDelete(t *testing.T) {
	mode := storage.Storage{Utils: &utils.Utils{}}
	l := LocalizationController{Svc: mode}

	APP := fiber.New()

	APP.Delete("/:app/:module/:language", func(c *fiber.Ctx) error {
		appName := c.Params("app")
		moduleName := c.Params("module")
		languageName := c.Params("language")

		err := l.Svc.DeleteLanguage(appName, moduleName, languageName)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"type":    "Fetch Data",
				"message": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(200)
	})
	r := strings.NewReader("my request")
	req := httptest.NewRequest(http.MethodDelete, "/app/module/language", r)
	resp, _ := APP.Test(req, 1)

	assert.Equal(t, 200, resp.StatusCode)

}
