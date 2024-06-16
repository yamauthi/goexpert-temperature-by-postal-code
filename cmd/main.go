package main

import (
	"github.com/yamauthi/goexpert-temperature-by-postalcode/internal/application/service"
	"github.com/yamauthi/goexpert-temperature-by-postalcode/internal/configs"
	"github.com/yamauthi/goexpert-temperature-by-postalcode/internal/infra/api"
	"github.com/yamauthi/goexpert-temperature-by-postalcode/internal/infra/web"
)

func main() {
	conf, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	postalCodeRepository := api.NewViaCepApiRepository(
		api.ApiConfig{
			BaseURL: conf.ViaCepApiUrl,
		},
	)

	temperatureRepository := api.NewWeatherAPIRepository(
		api.ApiConfig{
			BaseURL: conf.WeatherApiUrl,
			ApiKey:  conf.WeatherApiKey,
		},
	)

	temperatureService := service.NewTemperatureByPostalCodeService(
		postalCodeRepository,
		temperatureRepository,
	)

	temperatureHandler := web.NewTemperatureByPostalCodeHandler(temperatureService)

	webserver := web.NewWebServer(
		conf,
		[]web.Route{
			{
				Path:    "/{postalCode}",
				Method:  "get",
				Handler: temperatureHandler.TemperatureByPostalCode,
			},
		},
	)

	webserver.Start()
}
