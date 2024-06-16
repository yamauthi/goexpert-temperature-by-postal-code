package configs

import "github.com/spf13/viper"

type Conf struct {
	ViaCepApiUrl  string `mapstructure:"VIACEP_API_URL"`
	WeatherApiUrl string `mapstructure:"WEATHER_API_URL"`
	WeatherApiKey string `mapstructure:"WEATHER_API_KEY"`
	WebServerUrl  string `mapstructure:"WEBSERVER_URL"`
}

func LoadConfig(path string) (*Conf, error) {
	var cfg *Conf

	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, err
}
