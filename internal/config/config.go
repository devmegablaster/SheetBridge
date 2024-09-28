package config

import (
	"log/slog"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
)

type Config struct {
	Database DatabaseConfig `mapstructure:"database"`
	Api      ApiConfig      `mapstructure:"api"`
	Auth     AuthConfig     `mapstructure:"auth"`
	Crypto   CryptoConfig   `mapstructure:"crypto"`
	Kafka    KafkaConfig    `mapstructure:"kafka"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

type ApiConfig struct {
	Port    string `mapstructure:"port"`
	Version string `mapstructure:"version"`
}

type AuthConfig struct {
	JWTSecret string       `mapstructure:"jwt_secret"`
	Google    GoogleConfig `mapstructure:"google"`
}

type GoogleConfig struct {
	ClientID     string   `mapstructure:"client_id"`
	ClientSecret string   `mapstructure:"client_secret"`
	CallbackURL  string   `mapstructure:"callback_url"`
	Scopes       []string `mapstructure:"scopes"`
}

type CryptoConfig struct {
	AESKey string `mapstructure:"aes_key"`
}

type KafkaConfig struct {
	Broker     string `mapstructure:"broker"`
	SynkTopic  string `mapstructure:"synk_topic"`
	SynkGroup  string `mapstructure:"synk_group"`
	WriteTopic string `mapstructure:"write_topic"`
	WriteGroup string `mapstructure:"write_group"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Init() {
	viper.SetConfigFile(os.Getenv("CONFIG_PATH"))
	if err := viper.ReadInConfig(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	if err := viper.Unmarshal(c); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	c.PopulateEnv()

	slog.Info("✅Config Initialized")
}

func (c *Config) PopulateEnv() {
	c.Database.Host = os.ExpandEnv(c.Database.Host)
	c.Database.Port = os.ExpandEnv(c.Database.Port)
	c.Database.User = os.ExpandEnv(c.Database.User)
	c.Database.Password = os.ExpandEnv(c.Database.Password)
	c.Database.Name = os.ExpandEnv(c.Database.Name)

	c.Api.Port = os.ExpandEnv(c.Api.Port)

	c.Auth.JWTSecret = os.ExpandEnv(c.Auth.JWTSecret)
	c.Auth.Google.ClientID = os.ExpandEnv(c.Auth.Google.ClientID)
	c.Auth.Google.ClientSecret = os.ExpandEnv(c.Auth.Google.ClientSecret)

	c.Crypto.AESKey = os.ExpandEnv(c.Crypto.AESKey)

	c.Kafka.Broker = os.ExpandEnv(c.Kafka.Broker)
}
