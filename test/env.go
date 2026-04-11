package test

import (
	"os"

	"github.com/joho/godotenv"
)

var env envModel

type envModel struct {
	BaseURL string
	APIKey  string
	Model   string
}

func init() {
	godotenv.Load()
	env = envModel{
		BaseURL: os.Getenv("BASE_URL"),
		APIKey:  os.Getenv("API_KEY"),
		Model:   os.Getenv("MODEL"),
	}
}
