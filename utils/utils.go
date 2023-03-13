package utils

import (
	"os"
	"path/filepath"
	"strings"
)

type UtilsI interface {
	Check() error
	CheckFolderTree() error
	WriteToFile(data string) error
	ReadFile() (string, error)
	DeleteFile() error
	MoveToTrash() error
	MoveToArchive() error
}

// WorkDir is the working directory of the application

type Utils struct {
	WorkDir string
	File    *os.File
}

// Check file exists or  open the file
func (s *Utils) Check() error {
	// Check file is exists

	if _, err := os.Stat("./locales/Latest/" + s.WorkDir); os.IsNotExist(err) {

		// if folder is not exists mkdir -p
		folder := filepath.Dir("./locales/Latest/" + s.WorkDir)
		if err := os.MkdirAll(folder, 0755); err != nil {
			return err
		}
		// create file

		file, err := os.Create("./locales/Latest/" + s.WorkDir)
		if err != nil {
			return err
		}
		s.File = file

	}
	//  open the file
	file, err := os.OpenFile("./locales/Latest/"+s.WorkDir, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	s.File = file
	return nil
}

// WriteToFile writes data to a file
func (s *Utils) WriteToFile(data string) error {

	// write to file
	_, err := s.File.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}

// ReadFile reads data from a file
func (s *Utils) ReadFile() (string, error) {

	content, err := os.ReadFile("./locales/Latest/" + s.WorkDir)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

// DeleteFile deletes a file
func (s *Utils) DeleteFile() error {
	oldDir := "./locales/Latest/" + s.WorkDir
	newDir := "./locales/Trash/" + s.WorkDir

	folder := filepath.Dir(newDir)
	if err := os.MkdirAll(folder, 0755); err != nil {
		return err
	}
	err := os.Rename(oldDir, newDir)
	if err != nil {
		return err
	}

	return nil
}

// MoveToTrash moves a file to trash
func (s *Utils) MoveToTrash() error {
	if strings.Contains(s.WorkDir, "Latest") {
		oldDir := s.WorkDir
		s.WorkDir = strings.Replace(s.WorkDir, "Latest", "Trash", 1)
		err := os.Rename(oldDir, s.WorkDir)
		if err != nil {
			return err
		}
	}
	if strings.Contains(s.WorkDir, "Archive") {
		oldDir := s.WorkDir
		s.WorkDir = strings.Replace(s.WorkDir, "Archive", "Trash", 1)
		err := os.Rename(oldDir, s.WorkDir)
		if err != nil {
			return err
		}
	}
	return nil
}

// MoveToArchive moves a file to trash
func (s *Utils) MoveToArchive() error {

	oldDir := "./locales/Latest/" + s.WorkDir
	newDir := "./locales/Archive/" + s.WorkDir

	folder := filepath.Dir(newDir)
	if err := os.MkdirAll(folder, 0755); err != nil {
		return err
	}

	err := os.Rename(oldDir, newDir)
	if err != nil {
		return err
	}

	return nil
}
