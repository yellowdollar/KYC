package configs

import (
	"KYC/iternals/models"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var AppSettings models.Configs

func ReadProjectSettings() error {
	fmt.Println("Load .env file")

	err := godotenv.Load()

	if err != nil {
		fmt.Println(".env file not found")
	}

	fmt.Println("Reading settings file: configs/configs.json")
	confFile, err := os.Open("iternals/configs/configs.json")

	if err != nil {
		return errors.New(fmt.Sprintf("Couldnt open config's file: %s", err.Error()))
	}

	// defer - Execute in the end of a function
	defer func(confFile *os.File) {
		err = confFile.Close()
		if err != nil {
			log.Fatal("Couldnt close config's file: ", err.Error())
		}
	}(confFile)

	if err = json.NewDecoder(confFile).Decode(&AppSettings); err != nil {
		return errors.New(fmt.Sprintf("Couldn't decode json config: %s", err.Error()))
	}

	return nil
}
