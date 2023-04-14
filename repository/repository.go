package repository

import (
	model "localization/structures"

	"gorm.io/gorm"
)

type RepositoryDB struct {
	DB *gorm.DB
}
type RepositoryInterface interface {
	Application(m model.Application) (model.Application, error)
	Module(m model.Module) (model.Module, error)
	Language(m model.Language) (model.Language, error)
	Version(m model.Version) (model.Version, error)

	ArchiveApp(m model.ArchiveApp) (model.ArchiveApp, error)
	ArchiveMod(m model.ArchiveMod) (model.ArchiveMod, error)
	ArchiveLng(m model.ArchiveLng) (model.ArchiveLng, error)

	TrashApp(m model.TrashApp) (model.TrashApp, error)
	TrashMod(m model.TrashMod) (model.TrashMod, error)
	TrashLang(m model.TrashLang) (model.TrashLang, error)

	LatestApp(m model.Latest_App) (model.Latest_App, error)
	LatestMod(m model.Latest_Module) (model.Latest_Module, error)
	LatestLng(m model.Latest_Lang) (model.Latest_Lang, error)
}

func (t RepositoryDB) Application(m model.Application) (model.Application, error) {
	var err error
	if err = t.DB.Create(&m).Table(m.TableName()).Error; err != nil {
		return m, err
	}
	return m, nil
}
func (t RepositoryDB) Module(m model.Module) (model.Module, error) {
	var err error
	if err = t.DB.Create(&m).Table(m.TableName()).Error; err != nil {
		return m, err
	}
	return m, nil
}
func (t RepositoryDB) Language(m model.Language) (model.Language, error) {
	var err error
	if err = t.DB.Create(&m).Table(m.TableName()).Error; err != nil {
		return m, err
	}
	return m, nil
}
func (t RepositoryDB) Version(m model.Version) (model.Version, error) {
	var err error
	if err = t.DB.Create(&m).Table(m.TableName()).Error; err != nil {
		return m, err
	}
	return m, nil
}
func (t RepositoryDB) LatestApp(m model.Latest_App) (model.Latest_App, error) {
	var err error
	if err = t.DB.Create(&m).Table(m.TableName()).Error; err != nil {
		return m, err
	}
	return m, nil
}
func (t RepositoryDB) LatestMod(m model.Latest_Module) (model.Latest_Module, error) {
	var err error
	if err = t.DB.Create(&m).Table(m.TableName()).Error; err != nil {
		return m, err
	}
	return m, nil
}
func (t RepositoryDB) LatestLng(m model.Latest_Lang) (model.Latest_Lang, error) {
	var err error
	if err = t.DB.Create(&m).Table(m.TableName()).Error; err != nil {
		return m, err
	}
	return m, nil
}

func (t RepositoryDB) ArchiveApp(m model.ArchiveApp) (model.ArchiveApp, error) {
	var err error
	if err = t.DB.Create(&m).Table(m.TableName()).Error; err != nil {
		return m, err
	}
	return m, nil
}
func (t RepositoryDB) ArchiveMod(m model.ArchiveMod) (model.ArchiveMod, error) {
	var err error
	if err = t.DB.Create(&m).Table(m.TableName()).Error; err != nil {
		return m, err
	}
	return m, nil
}
func (t RepositoryDB) ArchiveLng(m model.ArchiveLng) (model.ArchiveLng, error) {
	var err error
	if err = t.DB.Create(&m).Table(m.TableName()).Error; err != nil {
		return m, err
	}
	return m, nil
}
func (t RepositoryDB) TrashApp(m model.TrashApp) (model.TrashApp, error) {
	var err error
	if err = t.DB.Create(&m).Table(m.TableName()).Error; err != nil {
		return m, err
	}
	return m, nil
}
func (t RepositoryDB) TrashMod(m model.TrashMod) (model.TrashMod, error) {
	var err error
	if err = t.DB.Create(&m).Table(m.TableName()).Error; err != nil {
		return m, err
	}
	return m, nil
}
func (t RepositoryDB) TrashLang(m model.TrashLang) (model.TrashLang, error) {
	var err error
	if err = t.DB.Create(&m).Table(m.TableName()).Error; err != nil {
		return m, err
	}
	return m, nil
}
