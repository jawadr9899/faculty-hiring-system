package config

import "os"

type Config struct {
	Port  string
	DBUrl string
	SqliteDBPath string
}

func LoadConfig() *Config{
	return &Config{
		Port: getEnv("PORT","8080"),
		DBUrl: getEnv("DATABASE_URL","sqlite:./dev.db"),
		SqliteDBPath: getEnv("SQLITE_PATH","sqlite:./test.db"),
	}
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}