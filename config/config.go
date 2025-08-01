package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret                string
	Port                     string
	ProjectID                string
	DBName                   string
	DBUser                   string
	DBHost                   string
	ORDERTOPICID			 string
	REDISHOST				 string
}

func Load() Config {
	_ = godotenv.Load()

	return Config{
		JWTSecret:                os.Getenv("JWT_SECRET"),
		Port:                     os.Getenv("DB_PORT"),
		ProjectID:                os.Getenv("GCP_PROJECT_ID"),
		DBName:                   os.Getenv("DB_NAME"),
		DBUser:                   os.Getenv("DB_USER"),
		DBHost:                   os.Getenv("DB_HOST"),
		ORDERTOPICID: 			  os.Getenv("ORDER_TOPIC_ID"),
		REDISHOST: 				  os.Getenv("REDIS_HOST"),
	}
}
