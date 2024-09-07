package configs

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

var Configs *Config

type Config struct {
	Server          ServerConfig
	WorkerAConsumer WorkerAConsumerCfg `mapstructure:"worker_a_consumer"`
	Producer        ProducerCfg
}

type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type ProducerCfg struct {
	Host            string `json:"host"`
	User            string `json:"user"`
	Pass            string `json:"pass"`
	AwsAccessKey    string `json:"aws_access_key"          mapstructure:"aws_access_key"`
	AwsSecretKey    string `json:"aws_secret_key"          mapstructure:"aws_secret_key"`
	AwsRole         string `json:"aws_role"                mapstructure:"aws_role"`
	AwsRegion       string `json:"aws_region"              mapstructure:"aws_region"`
	SASLMechanism   string `json:"sasl_mechanism"          mapstructure:"sasl_mechanism"`
	AutoOffsetReset string `json:"auto_offset_reset"              mapstructure:"auto_offset_reset"`
	Version         string `json:"version"`
	MaxRetries      string `json:"max_retries"             mapstructure:"max_retries"`
}

type WorkerAConsumerCfg struct {
	Host            string `json:"host"`
	User            string `json:"user"`
	Pass            string `json:"pass"`
	AwsAccessKey    string `json:"aws_access_key"          mapstructure:"aws_access_key"`
	AwsSecretKey    string `json:"aws_secret_key"          mapstructure:"aws_secret_key"`
	AwsRole         string `json:"aws_role"                mapstructure:"aws_role"`
	AwsRegion       string `json:"aws_region"              mapstructure:"aws_region"`
	SASLMechanism   string `json:"sasl_mechanism"          mapstructure:"sasl_mechanism"`
	AutoOffsetReset string `json:"auto_offset_reset"              mapstructure:"auto_offset_reset"`
	WorkerATopic    string `json:"worker_a_topic"              mapstructure:"worker_a_topic"`
	WokerAGroupId   string `json:"worker_a_group_id"              mapstructure:"worker_a_group_id"`
	WorkerATimeOut  string `json:"worker_a_timeout"            mapstructure:"worker_a_timeout"`
	Version         string `json:"version"`
	MaxRetries      string `json:"max_retries"             mapstructure:"max_retries"`
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
