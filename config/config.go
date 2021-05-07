package config

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Get(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok || len(value) == 0 {
		log.Fatalf("%s is undefined", key)
	}
	return value
}

func GetInt(key string) int {
	value, ok := os.LookupEnv(key)
	if !ok || len(value) == 0 {
		log.Fatalf("%s is undefined", key)
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("unable to cast %#v of type %T to int", key, key)
	}
	return v
}

func GetBool(key string) bool {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("%s is undefined", key)
	}
	v, err := strconv.ParseBool(value)
	if err != nil {
		log.Fatalf("unable to cast %#v of type %T to bool", key, key)
	}
	return v
}

func GetKeys(keys []string) map[string]string {
	values := map[string]string{}
	for _, key := range keys {
		values[key] = Get(key)
	}
	return values
}

func LoadLocalConfig() {
	v := flag.String(APP_ENV, "production", APP_ENV_FLAG_MESSAGE)
	flag.Parse()
	if *v == "development" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("Error %s", err.Error())
		}
	}
}
