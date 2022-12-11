package config

import "github.com/joho/godotenv"

var envMap map[string]string

func SetUpEnvironment() {
	eMap, err := godotenv.Read(".env")
	if err != nil {
		panic(err)
	} else {
		envMap = eMap
	}
}

func GetEnvMap() map[string]string {
	return envMap
}
