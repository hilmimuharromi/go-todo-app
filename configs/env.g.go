package configs

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type EnvType struct {
	JwtSecret        string
	DatabaseUrl      string
	DatabaseHost     string
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
}

var Env EnvType

func init() {
	setEnv(&Env.DatabasePort, "DB_PORT")
	setEnv(&Env.DatabaseHost, "DB_HOST")
	setEnv(&Env.DatabaseUser, "DB_USER")
	setEnv(&Env.DatabasePassword, "DB_PASSWORD")
	setEnv(&Env.DatabaseName, "DB_NAME")
	setEnv(&Env.JwtSecret, "JWT_SECRET")
}

func setEnv(envVariable interface{}, envName string, defaultValue ...interface{}) {
	var envValue = GetEnv(envName)
	switch envVariable.(type) {
	case *string:
		*envVariable.(*string) = envValue
		if *envVariable.(*string) == "" {
			if len(defaultValue) > 0 {
				*envVariable.(*string) = defaultValue[0].(string)
			}
		}
	case *[]byte:
		if envValue == "" {
			if len(defaultValue) > 0 {
				*envVariable.(*[]byte) = defaultValue[0].([]byte)
			} else {
				*envVariable.(*[]byte) = []byte("")
			}
		} else {
			*envVariable.(*[]byte) = []byte(envValue)
		}
	case *int:
		if envValue == "" {
			if len(defaultValue) > 0 {
				*envVariable.(*int) = defaultValue[0].(int)
			} else {
				*envVariable.(*int) = 0
			}
		} else {
			newInt, _ := strconv.Atoi(os.Getenv(envName))
			*envVariable.(*int) = newInt
		}
	case *bool:
		if envValue == "" {
			if len(defaultValue) > 0 {
				*envVariable.(*bool) = defaultValue[0].(bool)
			} else {
				*envVariable.(*bool) = false
			}
		} else {
			if envValue == "1" {
				*envVariable.(*bool) = true
			} else {
				*envVariable.(*bool) = false
			}
		}
	}
}

func GetEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}
