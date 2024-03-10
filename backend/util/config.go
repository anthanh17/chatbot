package util

import (
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variables.
type Config struct {
	OpenaiApiKey  string `mapstructure:"OPENAI_API_KEY"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ChunkSize     int    `mapstructure:"CHUNK_SIZE"`
	OverlapPct    int    `mapstructure:"OVERLAP_PCT"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
