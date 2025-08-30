package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/champion19/Flighthours_backend/tools/utils"
)

type Config struct {
	Database          Database          `json:"database"`
	JWT               JWTConfig         `json:"jwt"`
	API               APIConfig         `json:"api"`
	VerificationToken VerificationToken `json:"verification_token"`
	Resend            ResendConfig      `json:"resend"`
}

type Database struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Schema   string `json:"schema"`
}

type APIConfig struct {
	BaseURL string `json:"base_url"`
}

type VerificationToken struct {
	ExpirationTime time.Duration `json:"expiration_time"`
}

type ResendConfig struct {
	APIKey    string `json:"api_key"`
	FromEmail string `json:"from_email"`
}

func Load() Config {
	root, err := utils.FindModuleRoot()
	if err != nil {
		log.Fatal("error read config: ", err)
	}

	path := filepath.Join(root, "/config/default-config.json")
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("error read config: ", err)
	}

	var configRaw struct {
		Database          Database     `json:"database"`
		JWT               JWTConfigRaw `json:"jwt"`
		API               APIConfig    `json:"api"`
		VerificationToken struct {
			ExpirationTime string `json:"expiration_time"`
		} `json:"verification_token"`
		Resend ResendConfig `json:"resend"`
	}

	err = json.Unmarshal(file, &configRaw)
	if err != nil {
		log.Fatal("error unmarshal config: ", err)
	}

	jwtDuration, err := time.ParseDuration(configRaw.JWT.ExpirationTime)
	if err != nil {
		log.Fatal("error parsing jwt duration: ", err)
	}

	verificationDuration, err := time.ParseDuration(configRaw.VerificationToken.ExpirationTime)
	if err != nil {
		log.Fatal("error parsing verification token duration: ", err)
	}

	config := Config{
		Database: configRaw.Database,
		JWT: JWTConfig{
			SecretKey:      configRaw.JWT.SecretKey,
			ExpirationTime: jwtDuration,
		},
		API: configRaw.API,
		VerificationToken: VerificationToken{
			ExpirationTime: verificationDuration,
		},
		Resend: configRaw.Resend,
	}
	return config
}

type JWTConfig struct {
	SecretKey      string        `json:"secret_key"`
	ExpirationTime time.Duration `json:"expiration_time"`
}

type JWTConfigRaw struct {
	SecretKey      string `json:"secret_key"`
	ExpirationTime string `json:"expiration_time"`
}
