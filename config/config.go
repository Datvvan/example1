package config

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Configuration struct {
	PORT               string `mapstructure:"PORT"`
	DB_USER            string `mapstructure:"DB_USER"`
	DB_PASSWORD        string `mapstructure:"DB_PASSWORD"`
	DB_NAME            string `mapstructure:"DB_NAME"`
	DB_HOST            string `mapstructure:"DB_HOST"`
	ENV                string `mapstructure:"ENV"`
	LINECHAT_CLIENT_ID string `mapstructure:"LINECHAT_CLIENT_ID"`

	CHILLPAY_MERCHANT_CODE string `mapstructure:"CHILLPAY_MERCHANT_CODE"`
	CHILLPAY_API_KEY       string `mapstructure:"CHILLPAY_API_KEY"`
	CHILLPAY_MD5_KEY       string `mapstructure:"CHILLPAY_MD5_KEY"`
	CHILLPAY_SANDBOX_URL   string `mapstructure:"CHILLPAY_SANDBOX_URL"`
	CHILLPAY_OFFICIAL_URL  string `mapstructure:"CHILLPAY_OFFICIAL_URL"`

	BUCKET_NAME      string `mapstructure:"BUCKET_NAME"`
	STORAGE_ENDPOINT string `mapstructure:"STORAGE_ENDPOINT"`

	FACE_READING_AI_URL  string `mapstructure:"FACE_READING_AI_URL"`
	FACE_READING_API_KEY string `mapstructure:"FACE_READING_API_KEY"`
}

var Default Configuration

func Init(confPath string) {
	c := Configuration{}
	viper.SetConfigFile(confPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Warning(err, "Config file not found")
	}
	if err := viper.Unmarshal(&c); err != nil {
		log.Panic(err, "Error Unmarshal Viper Config")
	}

	Default = c
}
