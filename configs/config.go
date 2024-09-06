package configs

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

var Configs *Config

type Config struct {
	Server      ServerConfig
	InboundItmx InboundItmxConfig `mapstructure:"inbound_itmx"`
}

type ServerConfig struct {
	Host string
	Port string
}

type InboundItmxConfig struct {
	Host               string `json:"host"`
	User               string `json:"user"`
	Pass               string `json:"pass"`
	AwsAccessKey       string `json:"aws_access_key"          mapstructure:"aws_access_key"`
	AwsSecretKey       string `json:"aws_secret_key"          mapstructure:"aws_secret_key"`
	AwsRole            string `json:"aws_role"                mapstructure:"aws_role"`
	AwsRegion          string `json:"aws_region"              mapstructure:"aws_region"`
	SASLMechanism      string `json:"sasl_mechanism"          mapstructure:"sasl_mechanism"`
	AutoOffsetReset    string `json:"auto_offset_reset"              mapstructure:"auto_offset_reset"`
	InboundItmxTopic   string `json:"inbound_itmx_topic"              mapstructure:"inbound_itmx_topic"`
	InboundItmxGroupId string `json:"inbound_itmx_group"              mapstructure:"inbound_itmx_group_id"`
	Version            string `json:"version"`
	MaxRetries         string `json:"max_retries"             mapstructure:"max_retries"`
}

func LoadConfigs(config interface{}) {
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
		panic(err)
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		panic(err)
	}
}
