package service

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/yamauthi/goexpert-temperature-by-postalcode/internal/application/repository"
	"github.com/yamauthi/goexpert-temperature-by-postalcode/internal/domain/entity"
)

var ErrInvalidPostalCode = errors.New("invalid postal code")
var ErrPostalCodeNotFound = errors.New("can not find postal code")

type TemperatureByPostalCodeService struct {
	PostalAddressRepository repository.PostalAddressRepositoryInterface
	TemperatureRepository   repository.TemperatureRepositoryInterface
}

func NewTemperatureByPostalCodeService(
	postalRepository repository.PostalAddressRepositoryInterface,
	temperatureRepository repository.TemperatureRepositoryInterface,
) *TemperatureByPostalCodeService {
	return &TemperatureByPostalCodeService{
		PostalAddressRepository: postalRepository,
		TemperatureRepository:   temperatureRepository,
	}
}

type PostalCodeInput struct {
	PostalCode string
}

type TemperatureOutput struct {
	TempCelsius    float64 `json:"temp_C"`
	TempFahrenheit float64 `json:"temp_F"`
	TempKelvin     float64 `json:"temp_K"`
}

func (t *TemperatureByPostalCodeService) Execute(input PostalCodeInput) (TemperatureOutput, error) {
	if !isPostalCodeValid(input.PostalCode) {
		return TemperatureOutput{}, ErrInvalidPostalCode
	}

	address, err := t.PostalAddressRepository.GetAddress(input.PostalCode)
	if err != nil {
		return TemperatureOutput{}, fmt.Errorf("unexpected error when trying to get postal code information: %s", err)
	}

	if (address == entity.PostalAddress{}) {
		return TemperatureOutput{}, ErrPostalCodeNotFound
	}

	tempUnits := []entity.TemperatureUnit{
		entity.Celsius,
		entity.Fahrenheit,
		entity.Kelvin,
	}

	temperatures, err := t.TemperatureRepository.GetCityTemperature(address.City, tempUnits)
	if err != nil {
		return TemperatureOutput{}, fmt.Errorf("unexpected error when trying to get temperature information: %s", err)
	}

	return TemperatureOutput{
		TempCelsius:    temperatures[entity.Celsius].Degrees,
		TempFahrenheit: temperatures[entity.Fahrenheit].Degrees,
		TempKelvin:     temperatures[entity.Kelvin].Degrees,
	}, nil
}

func isPostalCodeValid(postalCode string) bool {
	match, err := regexp.MatchString("([0-9]{8})", postalCode)

	if err == nil && len(postalCode) == 8 && match {
		return true
	}

	return false
}
