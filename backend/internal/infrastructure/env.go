package infrastructure

import (
	"os"
	"strconv"
)

type Environment struct {
	// Application configuration
	APP_ENV    string
	APP_NAME   string
	APP_PORT   string
	APP_ORIGIN string
	APP_LANG   string
	// Database configuration
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_PORT     string
	// JWT configuration
	JWT_KEY                string
	JWT_EXPIRE_MIN         int
	JWT_REFRESH_EXPIRE_MIN int
}

func getEnv(key string, value ...string) string {
	if len(value) == 0 {
		value = append(value, "")
	}

	result, ok := os.LookupEnv(key)

	if !ok {
		return value[0]
	}

	return result
}

func NewEnvironment() *Environment {
	expMin, expMinErr := strconv.Atoi(getEnv("JWT_EXPIRE_MIN"))

	if expMinErr != nil {
		expMin = 110
	}

	expRefreshMin, expRefreshMinErr := strconv.Atoi(getEnv("JWT_REFRESH_EXPIRE_MIN"))

	if expRefreshMinErr != nil {
		expRefreshMin = 120
	}

	return &Environment{
		// Application configuration
		APP_ENV:    getEnv("APP_ENV"),
		APP_NAME:   getEnv("APP_NAME"),
		APP_PORT:   getEnv("APP_PORT"),
		APP_ORIGIN: getEnv("APP_ORIGIN"),
		APP_LANG:   getEnv("APP_LANG"),
		// Database configuration
		DB_HOST:     getEnv("DB_HOST"),
		DB_USER:     getEnv("DB_USER"),
		DB_PASSWORD: getEnv("DB_PASSWORD"),
		DB_NAME:     getEnv("DB_NAME"),
		DB_PORT:     getEnv("DB_PORT"),
		// JWT configuration
		JWT_KEY:                getEnv("JWT_KEY"),
		JWT_EXPIRE_MIN:         expMin,
		JWT_REFRESH_EXPIRE_MIN: expRefreshMin,
	}
}
