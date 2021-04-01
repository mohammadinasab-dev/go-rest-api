package configuration

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBUsername    string `mapstructure:"DB_USERNAME"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBAddress     string `mapstructure:"DB_ADDRESS"`
	DBName        string `mapstructure:"DB_NAME"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	JWTKey        string `mapstructure:"JWT_KEY"`
}
type ConfigTest struct {
	DBDriver      string `mapstructure:"TEST_DB_DRIVER"`
	DBUsername    string `mapstructure:"TEST_DB_USERNAME"`
	DBPassword    string `mapstructure:"TEST_DB_PASSWORD"`
	DBAddress     string `mapstructure:"TEST_DB_ADDRESS"`
	DBName        string `mapstructure:"TEST_DB_NAME"`
	ServerAddress string `mapstructure:"TEST_SERVER_ADDRESS"`
	JWTKey        string `mapstructure:"TEST_JWT_KEY"`
}

func LoadConfig(path string) (config Config, err error) {
	config = Config{}
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}
	fmt.Print(config.DBDriver)
	return config, nil
}

func LoadConfigTest(path string) (configtest ConfigTest, err error) {
	configtest = ConfigTest{}
	viper.AddConfigPath(path)
	viper.SetConfigName("testapp")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return ConfigTest{}, err
	}
	if err := viper.Unmarshal(&configtest); err != nil {
		return ConfigTest{}, err
	}
	return configtest, nil
}

func LoadSetup(path string) (string, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("setup")
	viper.SetConfigType("json")
	if err := viper.ReadInConfig(); err != nil {
		return "", err
	}
	if viper.GetString("run.environment") == "test" {
		viper.Set("log.logout", "stdOut")
		viper.Set("log.logformat", "Text")
		return "test", nil
	}
	if viper.GetString("run.environment") == "product" {
		viper.SetDefault("log.logout", "file")
		viper.SetDefault("log.logformat", "json")
		return "product", nil
	}
	return "", errors.New("run.environment variable not set or Not defined")
}
