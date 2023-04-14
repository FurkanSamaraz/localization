package endpoints

import (
	db "localization/config"
	"localization/controller"
	_ "localization/docs"

	repository "localization/repository"
	storage "localization/storage"
	utils "localization/utils"

	swagger "github.com/arsmn/fiber-swagger/v2"

	"github.com/gofiber/fiber/v2"
)

func SetupPoints(app *fiber.App) {

	mode := storage.Storage{Utils: &utils.Utils{}, Repo: repository.RepositoryDB{DB: db.Collection()}}
	point := controller.LocalizationController{Svc: mode}
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "/swagger/doc.json",
		DeepLinking: false,
	}))
	{

		app.Get("/:app", point.ReadApp)
		app.Post("/:app", point.CreateApp)
		app.Put("/:app", point.UpdateApp)
		app.Delete("/:app", point.DeleteApp)

		app.Get("/:app/:module", point.ReadModule)
		app.Post("/:app/:module", point.CreateModule)
		app.Put("/:app/:module", point.UpdateModule)
		app.Delete("/:app/:module", point.DeleteModule)

		app.Get("/:app/:module/:language", point.ReadLanguage)
		app.Post("/:app/:module/:language", point.CreateLanguage)
		app.Put("/:app/:module/:language", point.UpdateLanguage)
		app.Delete("/:app/:module/:language", point.DeleteLanguage)

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
