package config

import (
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Env struct {
	ServiceNumber string
	Password      string
}

var (
	instance *Env
	once     sync.Once
)

func GetEnv() *Env {
	once.Do(func() {
		_ = godotenv.Load()
		instance = &Env{
			ServiceNumber: getField("SERVICE_NUMBER", ""),
			Password:      getField("PASSWORD", ""),
		}
	})
	return instance
}

func getField(field, fallback string) string {
	if value, ok := os.LookupEnv(field); ok {
		return value
	}
	return fallback
}
