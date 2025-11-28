package config

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	DBSSLMode  string

	// Добавляем настройки приложения
	AppPort string
	AppEnv  string

	// Настройки пула соединений
	DBMaxConns        int
	DBMinConns        int
	DBConnMaxLifetime time.Duration
	DBConnMaxIdleTime time.Duration
}

func Load() (*Config, error) {
	if err := LoadDotEnv(); err != nil {
		fmt.Printf("Note: .env file not found: %v\n", err)
	}

	cfg := &Config {
		DBUser: getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "7071"),
		DBName: getEnv("DB_NAME", "todo-app"),
		DBSSLMode: getEnv("DB_SSL_MODE", "disable"),

		AppPort: getEnv("APP_PORT", "8081"),
		AppEnv: getEnv("APP_ENV", "development"),

		DBMaxConns: getEnvAsInt("DB_MAX_CONNS", 20),
		DBMinConns: getEnvAsInt("DB_MIN_CONNS", 2),
		DBConnMaxLifetime: getEnvAsDuration("DB_CONN_MAX_LIFETIME", time.Hour),
		DBConnMaxIdleTime: getEnvAsDuration("DB_CONN_MAX_IDLE_TIME", 30*time.Minute),
	}
}

// validate проверяет обязательные поля
func (c *Config) validate() error {
	if c.DBPassword == "" {
		return errors.New("DB_PASSWORD is required")
	}
	if c.DBName == "" {
		return errors.New("DB_NAME is required")
	}
	if c.DBMaxConns < c.DBMinConns {
		return errors.New("DB_MAX_CONNS cannot be less than DB_MIN_CONNS")
	}

	return nil
}

// DSN (Data Source Name) — это строка подключения к базе данных.

// GetDSN возвращает DSN строку для подключения
func (c *Config) GetDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBSSLMode,
	)
}

// LoadDotEnv загружает переменные из .env файла
func LoadDotEnv() error {
	file, err := os.Open("./.env")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	linesLoaded := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Пропускаем пустые строки и комментарии
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Разделяем ключ и значение
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Убираем кавычки
		value = strings.Trim(value, `"'`)

		// Устанавливаем только если переменная еще не установлена
		if os.Getenv(key) == "" {
			os.Setenv(key, value)
			linesLoaded++
		}
	}

	return scanner.Err()
}

// getEnv получает переменную окружения или значение по умолчанию
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt получает переменную окружения как int
func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// getEnvAsDuration получает переменную окружения как time.Duration
func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}

// getEnvAsBool получает переменную окружения как bool
func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}
