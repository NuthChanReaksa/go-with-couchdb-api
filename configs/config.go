package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	CouchDBURL  string
	CouchDBName string
	ServerPort  string
}

// LoadConfig loads environment variables and returns a Config struct
func LoadConfig() (*Config, error) {
	// Load .env file if present
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables.")
	}

	config := &Config{
		CouchDBURL:  getEnv("COUCHDB_URL", "http://localhost:5984"), // Default to localhost if not set
		CouchDBName: getEnv("COUCHDB_NAME", "products"),             // Default database name if not set
		ServerPort:  getEnv("SERVER_PORT", "8080"),                  // Default to port 8080
	}

	return config, nil
}

// getEnv reads an environment variable or returns a fallback value if not set
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
