package config

import (
	"log"

	"github.com/knadh/koanf/parsers/dotenv"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type AppConfig struct {
	AppEnv      string `koanf:"APP_ENV"`
	AppPort     string `koanf:"APP_PORT"`
	AppGrpcPort string `koanf:"APP_GRPC_PORT"`
}

func SetupConfig() *AppConfig {
	// Inisialisasi instance koanf
	k := koanf.New(".")

	// Load dari .env
	if err := k.Load(file.Provider(".env"), dotenv.Parser()); err != nil {
		log.Println("error loading .env file : ", err)
	}

	// Load dari OS
	if err := k.Load(env.Provider("", ".", func(s string) string {
		return s
	}), nil); err != nil {
		log.Println("error loading env provider", err)
	}

	// Set Default value
	config := AppConfig{
		AppEnv:      "development",
		AppPort:     "8080",
		AppGrpcPort: "7000",
	}

	// Pindah data koanf ke struct go
	if err := k.Unmarshal("", &config); err != nil {
		log.Fatalf("Gagal memetakan konfigurasi: %v\n", err)
	}

	return &config
}
