package common

import (
  "log"
  "os"
  "strings"

  "github.com/joho/godotenv"
)

func LoadEnv() {
  if err := godotenv.Load(".env"); err != nil {
    log.Fatalf("Error loading .env file: %v", err)
  }
}

func ParseEnvArray(key string) []string {
  parts := strings.Split(os.Getenv(key), ",")
  for i := range parts {
    parts[i] = strings.TrimSpace(parts[i])
  }
  return parts
}
