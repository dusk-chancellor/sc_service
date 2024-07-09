package configs

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)
// читаем yaml конфиг
type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		DbUser     string `yaml:"dbuser"`
		DbPassword string `yaml:"dbpassword"`
		DbName     string `yaml:"dbname"`
		DbHost     string `yaml:"dbhost"`
		DbPort     string `yaml:"dbport"`
	} `yaml:"database"`
}

func ReadConfig() *Config { // если происходят ошибки, то сразу паникуем
	// без конфига приложение не запустится
	err := godotenv.Load(".env") // читаем энв
	if err != nil {
		panic(err)
	}

	configPath := os.Getenv("CONFIG_PATH")
	var cfg Config
	if err = cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic(err)
	}

	return &cfg
}
