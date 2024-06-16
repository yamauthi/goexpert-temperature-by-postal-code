package repository

import "github.com/yamauthi/goexpert-temperature-by-postalcode/internal/domain/entity"

type PostalAddressRepositoryInterface interface {
	GetAddress(postalCode string) (entity.PostalAddress, error)
}

type TemperatureRepositoryInterface interface {
	GetCityTemperature(city string, units []entity.TemperatureUnit) (map[entity.TemperatureUnit]entity.Temperature, error)
}
