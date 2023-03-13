package controller

import (
	"encoding/json"
	storage "localization/storage"

	"github.com/gofiber/fiber/v2"
)

type LocalizationController struct {
	Svc storage.StorageI
}

/*
 * ===============
 * App
 * ===============
 */

func CreateApp(c *fiber.Ctx) error {

	appName := c.Params("app")
	bod := c.Body()
	err := storage.CreateApp(bod, appName)

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
func ReadApp(c *fiber.Ctx) error {
	appName := c.Params("app")

	result, err := storage.ReadApp(appName) //, app)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": err.Error(),
		})
	}

	res := Json(result)
	return c.JSON(res)
}

func DeleteApp(c *fiber.Ctx) error {
	appName := c.Params("app")
	err := storage.DeleteApp(appName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(err)
}

func UpdateApp(c *fiber.Ctx) error {
	bod := c.Body()
	appName := c.Params("app")

	err := storage.UpdateApp(appName, bod)
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

func CreateModule(c *fiber.Ctx) error {
	module := c.Body()
	appName := c.Params("app")
	moduleName := c.Params("module")

	err := storage.CreateModule(appName, moduleName, module)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(nil)
}

func ReadModule(c *fiber.Ctx) error {

	appName := c.Params("app")
	moduleName := c.Params("module")

	result, err := storage.ReadModule(appName, moduleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": err.Error(),
		})
	}

	res := Json(*result)
	return c.JSON(res)
}

func DeleteModule(c *fiber.Ctx) error {

	appName := c.Params("app")
	moduleName := c.Params("module")

	err := storage.DeleteModule(appName, moduleName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(err)
}

func UpdateModule(c *fiber.Ctx) error {
	module := c.Body()
	appName := c.Params("app")
	moduleName := c.Params("module")

	err := storage.UpdateModule(appName, moduleName, module)
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

func CreateLanguage(c *fiber.Ctx) error {

	language := c.Body()
	appName := c.Params("app")
	moduleName := c.Params("module")
	languageName := c.Params("language")

	err := storage.CreateLanguage(appName, moduleName, languageName, language)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(err)
}

func ReadLanguage(c *fiber.Ctx) error {
	appName := c.Params("app")
	moduleName := c.Params("module")
	languageName := c.Params("language")

	result, err := storage.ReadLanguage(appName, moduleName, languageName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": err.Error(),
		})
	}
	res := Json(result)
	return c.JSON(res)
}

func DeleteLanguage(c *fiber.Ctx) error {
	appName := c.Params("app")
	moduleName := c.Params("module")
	languageName := c.Params("language")

	err := storage.DeleteLanguage(appName, moduleName, languageName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(200)
}

func UpdateLanguage(c *fiber.Ctx) error {
	language := c.Body()
	appName := c.Params("app")
	moduleName := c.Params("module")
	languageName := c.Params("language")

	err := storage.UpdateLanguage(appName, moduleName, languageName, language)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"type":    "Fetch Data",
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(err)
}
