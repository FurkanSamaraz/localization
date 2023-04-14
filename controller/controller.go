package controller

import (
	"encoding/json"
	storage "localization/storage"

	"github.com/gofiber/fiber/v2"
)

type LocalizationController struct {
	Svc storage.Storage
}

// ShowAccount godoc
// @Summary      Pagination CreateApp
// @Description  Pagination CreateApp
// @Tags         app/
// @Id					  ApiV1LocalizationCreateAppPagination
// @Accept       json
// @Produce      json
// @Param app path string true "app"
// @Success      200  {array} storage.StorageI.CreateApp(bod, appName)
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
// @Summary      Pagination ReadApp
// @Description  Pagination ReadApp
// @Tags         app/
// @Id					  ApiV1LocalizationReadAppPagination
// @Accept       json
// @Produce      json
// @Param app path string true "app"
// @Param app  path string true "app"
// @Success      200  {array}  storage.StorageI.ReadApp(appName)
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
// @Summary      Pagination DeleteApp
// @Description  Pagination DeleteApp
// @Tags         app/
// @Id					  ApiV1LocalizationDeleteAppPagination
// @Accept       json
// @Produce      json
// @Param app  path string true "app"
// @Success      200  {array}  storage.StorageI.DeleteApp(appName)
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
// @Summary      Pagination UpdateApp
// @Description  Pagination UpdateApp
// @Tags         app/
// @Id					  ApiV1LocalizationUpdateAppPagination
// @Accept       json
// @Produce      json
// @Param app  path string true "app"
// @Success      200  {array}  storage.StorageI.UpdateApp(bod, appName)
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
// @Summary      Pagination CreateModule
// @Description  Pagination CreateModule
// @Tags         app/module
// @Id					  ApiV1LocalizationCreateModulePagination
// @Accept       json
// @Produce      json
// @Param app  path string true "app"
// @Param module  path string true "module"
// @Success      200  {array}  storage.StorageI.CreateModule(appName, moduleName, module)
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
// @Summary      Pagination ReadModule
// @Description  Pagination ReadModule
// @Tags         app/module
// @Id					  ApiV1LocalizationReadModulePagination
// @Accept       json
// @Produce      json
// @Param app  path string true "app"
// @Param module  path string true "module"
// @Success      200  {array}  storage.StorageI.ReadModule(appName, moduleName)
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
// @Summary      Pagination DeleteModule
// @Description  Pagination DeleteModule
// @Tags         app/module
// @Id					  ApiV1LocalizationDeleteModulePagination
// @Accept       json
// @Produce      json
// @Param app  path string true "app"
// @Param module  path string true "module"
// @Success      200  {array}  storage.StorageI.DeleteModule(appName, moduleName)
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
// @Summary      Pagination UpdateModule
// @Description  Pagination UpdateModule
// @Tags         app/module
// @Id					  ApiV1LocalizationUpdateModulePagination
// @Accept       json
// @Produce      json
// @Param app  path string true "app"
// @Param module  path string true "module"
// @Success      200  {array}  storage.StorageI.UpdateModule(appName, moduleName, module)
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
// @Summary      Pagination CreateLanguage
// @Description  Pagination CreateLanguage
// @Tags         app/module/language
// @Id					  ApiV1LocalizationCreateLanguagePagination
// @Accept       json
// @Produce      json
// @Param app  path string true "app"
// @Param module  path string true "module"
// @Param language  path string true "language"
// @Success      200  {array}  storage.StorageI.CreateLanguage(appName, moduleName, languageName, language)
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
// @Summary      Pagination ReadLanguage
// @Description  Pagination ReadLanguage
// @Tags         app/module/language
// @Id					  ApiV1LocalizationReadLanguagePagination
// @Accept       json
// @Produce      json
// @Param app  path string true "app"
// @Param module  path string true "module"
// @Param language  path string true "language"
// @Success      200  {array}  storage.StorageI.ReadLanguage(appName, moduleName, languageName)
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
// @Summary      Pagination DeleteLanguage
// @Description  Pagination DeleteLanguage
// @Tags         app/module/language
// @Id					  ApiV1LocalizationDeleteLanguagePagination
// @Accept       json
// @Produce      json
// @Param app  path string true "app"
// @Param module  path string true "module"
// @Param language  path string true "language"
// @Success      200  {array}  storage.StorageI.DeleteLanguage(appName, moduleName, languageName)
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
// @Summary      Pagination UpdateLanguage
// @Description  Pagination UpdateLanguage
// @Tags          app/module/language
// @Id					  ApiV1LocalizationUpdateLanguagePagination
// @Accept       json
// @Produce      json
// @Param app  path string true "app"
// @Param module  path string true "module"
// @Param language  path string true "language"
// @Success      200  {array}  storage.StorageI.UpdateLanguage(appName, moduleName, languageName, language)
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
