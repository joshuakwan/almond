package alertmanager

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/astaxie/beego"
	"github.com/joshuakwan/almond/utils"
)

func TestLoadConfig(t *testing.T) {
	filename := utils.GetAlertmanagerConfigFilePath()
	config, err := LoadConfigFromFile(filename)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(config)
	bytes, err := json.Marshal(config)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(bytes))
}

func TestAppConfig(t *testing.T) {
	fmt.Println(beego.BConfig.RunMode)
}
