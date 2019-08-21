package config_test

import (
	"testing"

	nconfig "github.com/adhimaswaskita/AssetManagement/config"
)

func TestNewConfig(t *testing.T) {
	t.Run("NewConfig OK", func(t *testing.T) {
		filepath := "_fixtures/config.yaml"
		config, err := nconfig.NewConfig(filepath)
		if err != nil {
			t.Errorf("This should not be error")
		}
		if config.App.Name != "Asset Management" {
			t.Errorf("This config.App.Name should be equal to 'Asset Management' but have %v", config.App.Name)
		}
		if config.Repo.Name != "db" {
			t.Errorf("This config.Repo.Name should be equal to 'db' but have %v", config.Repo.Name)
		}
	})
	t.Run("NewConfig readfile NOK", func(t *testing.T) {
		filepath := "_fixtures/config.yam"
		_, err := nconfig.NewConfig(filepath)
		if err == nil {
			t.Errorf("This should be error from file paths")
		}
	})
}
