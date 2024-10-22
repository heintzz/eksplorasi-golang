package config

import "github.com/joho/godotenv"

func LoadConfig(filename string) error {
	err := godotenv.Load(filename)
	return err
}
