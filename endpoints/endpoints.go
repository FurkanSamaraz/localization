package endpoints

import (
	"localization/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupPoints(app *fiber.App) {

	//! fiber app için erişim entegrasyonu ?

	{

		app.Get("/:app", controller.ReadApp)
		app.Post("/:app", controller.CreateApp)
		app.Put("/:app", controller.UpdateApp)
		app.Delete("/:app", controller.DeleteApp)

		app.Get("/:app/:module", controller.ReadModule)
		app.Post("/:app/:module", controller.CreateModule)
		app.Put("/:app/:module", controller.UpdateModule)
		app.Delete("/:app/:module", controller.DeleteModule)

		app.Get("/:app/:module/:language", controller.ReadLanguage)
		app.Post("/:app/:module/:language", controller.CreateLanguage)
		app.Put("/:app/:module/:language", controller.UpdateLanguage)
		app.Delete("/:app/:module/:language", controller.DeleteLanguage)

	}

	// {
	// 	applicationGroup := app.Group("/app")

	// 	applicationGroup.Get("/list", localizationController.ReadApp)
	// 	applicationGroup.Post("/list", localizationController.CreateApp)
	// 	applicationGroup.Put("/list", localizationController.UpdateApp)
	// 	applicationGroup.Delete("/list", localizationController.DeleteApp)

	// 	applicationGroup.Get("/:app", localizationController.ReadApp)
	// 	{
	// 		moduleGroup := applicationGroup.Group("/:app/module")
	// 		moduleGroup.Get("/list", localizationController.ReadModule)
	// 		moduleGroup.Post("/list", localizationController.CreateModule)
	// 		moduleGroup.Put("/list", localizationController.UpdateModule)
	// 		moduleGroup.Delete("/list", localizationController.DeleteModule)

	// 		moduleGroup.Get("/:module", localizationController.ReadModule)
	// 		{
	// 			languagesGroup := moduleGroup.Group("/:module/language")
	// 			languagesGroup.Get("/list", localizationController.ReadLanguage)
	// 			languagesGroup.Post("/list", localizationController.CreateLanguage)
	// 			languagesGroup.Put("/list", localizationController.UpdateLanguage)
	// 			languagesGroup.Delete("/:language", localizationController.DeleteLanguage)
	// 		}

	// 	}

	// }

}
