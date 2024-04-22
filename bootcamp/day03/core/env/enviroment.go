package env

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	dotenv "github.com/joho/godotenv"
)

func LoadEnv[T comparable](config T) (T, error) {
	fields := reflect.VisibleFields(reflect.TypeOf(config))
	dotenv.Load()
	var res T
	for _, field := range fields {
		val := os.Getenv(field.Tag.Get("env"))
		if len(val) == 0 {
			return res, errors.New(
				fmt.Sprintf("Can't find: %s in enviroment", field.Tag.Get("env")),
			)
		}
		reflect.ValueOf(&res).
			Elem().FieldByName(field.Name).Set(reflect.ValueOf(val))
	}
	return res, nil
}
