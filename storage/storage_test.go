package storage

import (
	utils "localization/utils"
	"testing"
)

func TestCreateApp(t *testing.T) {
	mode := Storage{Utils: &utils.Utils{}}

	err := mode.CreateApp([]byte("test"), "test")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

}
func TestReadApp(t *testing.T) {
	mode := Storage{Utils: &utils.Utils{}}

	_, err := mode.ReadApp("test")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

}
func TestUpdateApp(t *testing.T) {
	mode := Storage{Utils: &utils.Utils{}}

	err := mode.UpdateApp("test", []byte("test"))
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
func TestDeleteApp(t *testing.T) {
	mode := Storage{Utils: &utils.Utils{}}

	err := mode.DeleteApp("test")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestCreateModule(t *testing.T) {
	mode := Storage{Utils: &utils.Utils{}}

	err := mode.CreateModule("test", "test", []byte("test"))
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
func TestReadModule(t *testing.T) {
	mode := Storage{Utils: &utils.Utils{}}

	_, err := mode.ReadModule("test", "test")
	if err != nil {
		t.Errorf("Error: %v", err)
	}

}

func TestUpdateModule(t *testing.T) {
	mode := Storage{Utils: &utils.Utils{}}

	err := mode.UpdateModule("test", "test", []byte("test"))
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
func TestDeleteModule(t *testing.T) {
	mode := Storage{Utils: &utils.Utils{}}

	err := mode.DeleteModule("test", "test")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
func TestCreateLanguage(t *testing.T) {
	mode := Storage{Utils: &utils.Utils{}}

	err := mode.CreateLanguage("test", "test", "test", []byte("test"))
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
func TestReadLanguage(t *testing.T) {
	mode := Storage{Utils: &utils.Utils{}}

	_, err := mode.ReadLanguage("test", "test", "test")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestUpdateLanguage(t *testing.T) {
	mode := Storage{Utils: &utils.Utils{}}

	err := mode.UpdateLanguage("test", "test", "test", []byte("test"))
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
func TestDeleteLanguage(t *testing.T) {
	mode := Storage{Utils: &utils.Utils{}}

	err := mode.DeleteLanguage("test", "test", "test")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}
