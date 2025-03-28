package configs

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Configs struct {
	DB struct {
		Driver   string `mapstructure:"DB_DRIVER"`
		Username string `mapstructure:"DB_USERNAME"`
		Host     string `mapstructure:"DB_HOST"`
		Port     string `mapstructure:"DB_PORT"`
		Name     string `mapstructure:"DB_DBNAME"`
		SllMode  string `mapstructure:"DB_SSL_MODE"`
		Encoding string `mapstructure:"DB_ENCODING"`
		Psw      string `mapstructure:"DB_PSW"`
	}
	GRPC struct {
		ConnectType string `mapstructure:"GRPC_CONNECT_TYPE"`
		Port        string `mapstructdure:"GRPC_PORT"`
	}
}

const (
	configDir  = "configs"
	configName = "config"
	configType = "env"
)

func NewConfigs() (*Configs, error) {
	configEnvPath := configDir + "/" + configName + "." + configType
	err := godotenv.Load(configEnvPath)
	if err != nil {
		log.Println("No .env file found", err)
		return nil, err
	}
	viper.AutomaticEnv()

	viper.AddConfigPath(configDir)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	var config Configs
	if err := viper.Unmarshal(&config.DB); err != nil {
		log.Println("Error unmarshal config", err)
		return nil, err
	}

	if err := viper.Unmarshal(&config.GRPC); err != nil {
		log.Println("Error unmarshal config", err)
		return nil, err
	}

	return &Configs{}, nil
}
