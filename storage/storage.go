package storage

import (
	"fmt"
	types "localization/types"
	utils "localization/utils"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type StorageI interface {
	CreateApp(app *types.AppStruct) error
	ReadApp(appName string) (*types.AppStruct, error)
	DeleteApp(appName string) error
	UpdateApp(appName string, app *types.AppStruct) error

	CreateModule(appName string, module *types.ModuleStruct) error
	ReadModule(appName string, moduleName string) (*types.ModuleStruct, error)
	DeleteModule(appName string, moduleName string) error
	UpdateModule(appName string, moduleName string, module *types.ModuleStruct) error

	CreateLanguage(appName string, moduleName string, language *types.LanguageStruct) error
	ReadLanguage(appName string, moduleName string, languageName string) (*types.LanguageStruct, error)
	DeleteLanguage(appName string, moduleName string, languageName string) error
	UpdateLanguage(appName string, moduleName string, languageName string, language *types.LanguageStruct) error

	// Write To File
	WriteToFile(path string, data string) error
	ReadFile(path string) (string, error)
	DeleteFile(path string) error
}

// WorkDir is the working directory of the application
type CrmHandler struct {
	Service utils.UtilsI
}
type Storage struct {
	Utils *utils.Utils
}

/*
 * ===============
 * App
 * ===============
 */

var Appversion = 0

func AutoMaticg(Dir string) int {
	var s utils.Utils
	Appversion = 0
	filepath.WalkDir(Dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			Appversion = 0
			s.Check()
			return nil
		} else {
			entries, _ := os.ReadDir(Dir)

			for _, e := range entries {

				V := strconv.Itoa(Appversion)
				nams := "v" + V + ".json"
				res1 := strings.Contains(e.Name(), nams)

				if res1 == true {
					Appversion = Appversion + 1
				}

			}
		}

		return nil
	})

	return Appversion
}

func CreateApp(apps []byte, appName string) error {
	//	var app types.AppStruct
	var s utils.Utils

	Version := AutoMaticg("./locales/Latest/" + appName)

	s.WorkDir = appName + "/" + "v" + strconv.Itoa(Version) + ".json"
	err := s.Check()
	if err != nil {
		fmt.Println("not file check")
	}

	err = s.WriteToFile(string(apps))

	return err
}

// ReadApp reads an app
func ReadApp(appName string) (string, error) {
	var s utils.Utils
	s.WorkDir = appName + "/" + "v1.json"

	result, err := s.ReadFile()
	if err != nil {
		fmt.Println("Not Read")
	}

	return result, err
}

// DeleteApp deletes an app
func DeleteApp(appName string) error {
	var s utils.Utils
	s.WorkDir = appName + "/" + "v1.json"
	s.DeleteFile()

	return nil
}

// UpdateApp updates an app
func UpdateApp(appName string, apps []byte) error {
	var s utils.Utils
	s.WorkDir = appName + "/" + "v1.json"

	errs := s.MoveToArchive()
	if errs != nil {
		return errs
	}
	s.Check()

	err := s.WriteToFile(string(apps))

	return err
}

/*
 * ===============
 * Module
 * ===============
 */

// CreateModule creates a new module
func CreateModule(appName string, modulename string, module []byte) error {
	var s utils.Utils

	Version := AutoMaticg("./locales/Latest/" + appName + "/" + modulename)

	s.WorkDir = appName + "/" + modulename + "/" + "v" + strconv.Itoa(Version) + ".json"

	err := s.Check()
	if err != nil {
		fmt.Println("not file check")
	}

	err = s.WriteToFile(string(module))

	return err
}

// ReadModule reads a module
func ReadModule(appName string, modulename string) (*string, error) {
	var s utils.Utils
	s.WorkDir = appName + "/" + modulename + "/" + "v1.json"

	result, err := s.ReadFile()
	if err != nil {
		fmt.Println("Not Read")
	}

	return &result, err

}

// DeleteModule deletes a module
func DeleteModule(appName string, moduleName string) error {
	var s utils.Utils
	s.WorkDir = appName + "/" + moduleName + "/" + "v1.json"

	err := s.DeleteFile()

	return err
}

// UpdateModule updates a module
func UpdateModule(appName string, moduleName string, module []byte) error {
	var s utils.Utils
	s.WorkDir = appName + "/" + moduleName + "/" + "v1.json"

	errs := s.MoveToArchive()
	if errs != nil {
		return errs
	}

	s.Check()

	err := s.WriteToFile(string(module))

	return err
}

/*
 * ===============
 * Language
 * ===============
 */

// CreateLanguage creates a new language
func CreateLanguage(appName string, moduleName string, languageName string, language []byte) error {
	var s utils.Utils
	Version := AutoMaticg("./locales/Latest/" + appName + "/" + moduleName + "/" + languageName)

	s.WorkDir = appName + "/" + moduleName + "/" + languageName + "/" + "v" + strconv.Itoa(Version) + ".json"
	s.Check()

	err := s.WriteToFile(string(language))

	return err
}

// ReadLanguage reads a language
func ReadLanguage(appName string, moduleName string, languageName string) (string, error) {
	var s utils.Utils

	s.WorkDir = appName + "/" + moduleName + "/" + languageName + "/" + "v1.json"

	result, err := s.ReadFile()
	if err != nil {
		fmt.Println("Not Read")
	}

	return result, err
}

// DeleteLanguage deletes a language
func DeleteLanguage(appName string, moduleName string, languageName string) error {
	var s utils.Utils
	s.WorkDir = appName + "/" + moduleName + "/" + languageName + "/" + "v1.json"

	err := s.DeleteFile()

	return err
}

// UpdateLanguage updates a language
func UpdateLanguage(appName string, moduleName string, languageName string, language []byte) error {
	var s utils.Utils
	s.WorkDir = appName + "/" + moduleName + "/" + languageName + "/" + "v1.json"

	errs := s.MoveToArchive()
	if errs != nil {
		return errs
	}

	s.Check()
	err := s.WriteToFile(string(language))

	return err

}
