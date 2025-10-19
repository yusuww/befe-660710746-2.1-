package main

import (
	"fmt"
	"os"
)

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	host := getEnv("DB_HOST", "localhost")
	name := getEnv("DB_NAME", "mydb")
	user := getEnv("DB_USER", "root")
	password := getEnv("DB_PASSWORD", "")
	port := getEnv("DB_PORT", "5432")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, name)
	fmt.Println(connStr)
}