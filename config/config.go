package config

import (
	"aubbiali/db-import/model"
	"aubbiali/db-import/utility"
	"encoding/json"
	"log"
	"os"

	"github.com/go-playground/validator"
)

var Config model.Config

func ReadConfig() {
	file, err := os.ReadFile("config.json")
	utility.CheckError(err)

	validateConfig(file)

}

func validateConfig(file []byte) {

	v := validator.New()
	//v.RegisterValidation("Fields", checkMap, true)
	utility.CheckError(json.Unmarshal(file, &Config))

	log.Println(Config)
	utility.CheckError(v.Struct(Config))

	// add custom validator for the map

}

func checkMap(f validator.FieldLevel) bool {
	// TODO fix validator map
	log.Println(f.Field())
	return true
}
