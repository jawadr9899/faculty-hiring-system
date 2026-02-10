package config

import "os"

type Config struct {
	Port       string
	DBUrl      string
	JWTSecret  string
	AiApiKey   string
	AiApiModel string
	UploadsPath string
	AdminPassword string
	AdminEmail string
}

func LoadConfig() *Config {
	return &Config{
		Port:       GetEnv("PORT", "8080"),
		DBUrl:      GetEnv("DATABASE_URL", "sqlite:./dev.db"),
		JWTSecret:  GetEnv("JWT_SECRET", "DEFAULT_SECRET_KEY"),
		AiApiKey:   GetEnv("AI_API_KEY", "DEFAULT_AI_API_KEY"),
		AiApiModel: GetEnv("AI_API_MODEL", "DEFAULT_AI_MODEL"),
		UploadsPath: GetEnv("UPLOADS_PATH", "DEFAULT_UPLOADS_PATH"),
		AdminEmail: GetEnv("ADMIN_EMAIL", "DEFAUL_ADMIN_EMAIL"),
		AdminPassword: GetEnv("ADMIN_PASSWORD", "DEFAUL_ADMIN_PASSWORD"),

	}
}

func GetEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
