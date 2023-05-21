package controller

import (
	"encoding/json"
	storage "localization/storage"
	_ "localization/structures"

	"github.com/gofiber/fiber/v2"
)

type LocalizationController struct {
	Svc storage.Storage
}

// ShowAccount godoc
// @Summary      CreateApp
// @Description  CreateApp
// @Tags         APP
// @Id			 Localization_App
// @Accept       json
// @Produce      json
// @Param app body string true "app"
// @Success      200  {string} byte
// @Failure      400  {object} error
// @Failure      404  {object} error
// @Failure      500  {object} error
// @Router       /:app [post]
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

// ShowAccount godoc
// @Summary      ReadApp
// @Description  ReadApp
// @Tags         APP
// @Id			 Localization_App
// @Accept       json
// @Produce      json
// @Param app body string true "app"
// @Success      200  {array}  byte
// @Failure      400  {object} error
// @Failure      404  {object} error
// @Failure      500  {object} error
// @Router       /:app [get]
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

// ShowAccount godoc
// @Summary      DeleteApp
// @Description  DeleteApp
// @Tags         APP
// @Id			 Localization_DeleteApp
// @Accept       json
// @Produce      json
// @Param app body string true "app"
// @Success      200  {array}  byte
// @Failure      400  {object} error
// @Failure      404  {object} error
// @Failure      500  {object} error
// @Router       /:app [delete]
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

// ShowAccount godoc
// @Summary       UpdateApp
// @Description   UpdateApp
// @Tags          APP
// @Id					  Localization_App
// @Accept       json
// @Produce      json
// @Param app body string true "app"
// @Success      200  {array}  byte
// @Failure      400  {object} error
// @Failure      404  {object} error
// @Failure      500  {object} error
// @Router       /:app [put]
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

// ShowAccount godoc
// @Summary      CreateModule
// @Description  CreateModule
// @Tags         APP/MODULE
// @Id			 Localization_Module
// @Accept       json
// @Produce      json
// @Param app body string true "app"
// @Param module body string true "module"
// @Success      200  {array}  byte
// @Failure      400  {object} error
// @Failure      404  {object} error
// @Failure      500  {object} error
// @Router       /:app/:module [post]
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

// ShowAccount godoc
// @Summary      ReadModule
// @Description  ReadModule
// @Tags         APP/MODULE
// @Id			 Localization_Module
// @Accept       json
// @Produce      json
// @Param app body string true "app"
// @Param module body string true "module"
// @Success      200  {array}  byte
// @Failure      400  {object} error
// @Failure      404  {object} error
// @Failure      500  {object} error
// @Router       /:app/:module [get]
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

// ShowAccount godoc
// @Summary      DeleteModule
// @Description  DeleteModule
// @Tags         APP/MODULE
// @Id			 Localization_Module
// @Accept       json
// @Produce      json
// @Param app body string true "app"
// @Param module body string true "module"
// @Success      200  {array}  byte
// @Failure      400  {object} error
// @Failure      404  {object} error
// @Failure      500  {object} error
// @Router       /:app/:module [delete]
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

// ShowAccount godoc
// @Summary      UpdateModule
// @Description  UpdateModule
// @Tags         APP/MODULE
// @Id	    	 Localization_Module
// @Accept       json
// @Produce      json
// @Param app body string true "app"
// @Param module body string true "module"
// @Success      200  {array}  byte
// @Failure      400  {object} error
// @Failure      404  {object} error
// @Failure      500  {object} error
// @Router       /:app/:module [put]
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

// ShowAccount godoc
// @Summary      CreateLanguage
// @Description  CreateLanguage
// @Tags         APP/MODULE/LANGUAGE
// @Id			 Localization_Language
// @Accept       json
// @Produce      json
// @Param app body string true "app"
// @Param module body string true "module"
// @Param language body string true "language"
// @Success      200  {array}  byte
// @Failure      400  {object} error
// @Failure      404  {object} error
// @Failure      500  {object} error
// @Router       /:app/:module/:language [post]
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

// ShowAccount godoc
// @Summary      ReadLanguage
// @Description  ReadLanguage
// @Tags         APP/MODULE/LANGUAGE
// @Id			 Localization_ReadLanguage
// @Accept       json
// @Produce      json
// @Param app body string true "app"
// @Param module body string true "module"
// @Param language body string true "language"
// @Success      200  {array}  byte
// @Failure      400  {object} error
// @Failure      404  {object} error
// @Failure      500  {object} error
// @Router       /:app/:module/:language [get]
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

// ShowAccount godoc
// @Summary      DeleteLanguage
// @Description  DeleteLanguage
// @Tags         APP/MODULE/LANGUAGE
// @Id			 Localization_DeleteLanguage
// @Accept       json
// @Produce      json
// @Param app body string true "app"
// @Param module body string true "module"
// @Param language body string true "language"
// @Success      200  {array}  byte
// @Failure      400  {object} error
// @Failure      404  {object} error
// @Failure      500  {object} error
// @Router       /:app/:module/:language [delete]
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

// ShowAccount godoc
// @Summary      UpdateLanguage
// @Description  UpdateLanguage
// @Tags     	 APP/MODULE/LANGUAGE
// @Id			 Localization_UpdateLanguage
// @Accept       json
// @Produce      json
// @Param app body string true "app"
// @Param module body string true "module"
// @Param language body string true "language"
// @Success      200  {array}  byte
// @Failure      400  {object} error
// @Failure      404  {object} error
// @Failure      500  {object} error
// @Router       /:app/:module/:language [put]
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
