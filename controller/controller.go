package controller

import (
	"encoding/json"
	storage "localization/storage"

	"github.com/gofiber/fiber/v2"
)

type LocalizationController struct {
	Svc storage.Storage
}

/*
 * ===============
 * App
 * ===============
 */

func (l *LocalizationController) CreateApp(c *fiber.Ctx) error {

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
}
func Json(data string) interface{} {
	var Any interface{}
	json.Unmarshal([]byte(data), &Any)

	return Any
}
func (l *LocalizationController) ReadApp(c *fiber.Ctx) error {
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
}

func (l *LocalizationController) DeleteApp(c *fiber.Ctx) error {
	appName := c.Params("app")
	err := l.Svc.DeleteApp(appName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(err)
}

func (l *LocalizationController) UpdateApp(c *fiber.Ctx) error {
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
}

/*
 * ===============
 * Module
 * ===============
 */

func (l *LocalizationController) CreateModule(c *fiber.Ctx) error {
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
}

func (l *LocalizationController) ReadModule(c *fiber.Ctx) error {

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
}

func (l *LocalizationController) DeleteModule(c *fiber.Ctx) error {

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
}

func (l *LocalizationController) UpdateModule(c *fiber.Ctx) error {
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
}

/*
 * ===============
 * Language
 * ===============
 */

func (l *LocalizationController) CreateLanguage(c *fiber.Ctx) error {

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
}

func (l *LocalizationController) ReadLanguage(c *fiber.Ctx) error {
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
}

func (l *LocalizationController) DeleteLanguage(c *fiber.Ctx) error {
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
}

func (l *LocalizationController) UpdateLanguage(c *fiber.Ctx) error {
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
}
