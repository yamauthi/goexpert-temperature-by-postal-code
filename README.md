# GoExpert - Temperature by Postal

### Challenge description:

Receive a brazilian postal code (usually called CEP) and return the temperature on that location in Celsius, Fahrenheit and Kelvin.
- Postal code to be valid must be only numbers, 8 digits
- If a postal code is invalid should return status code 422 with error "invalid postal code"
- If a postal code is valid, but not found shold return status 404 with error "can not find postal code"
- You can get the location using an address API (like ViaCepAPI https://viacep.com.br/ )
- You can get the temperature using a weather API (like WeatherAPI https://www.weatherapi.com/)

 
### Requirements
- Valid API Key for WeatherAPI

### How to run locally with docker-compose
- Clone the repository
- Open terminal in project folder
- Copy cmd/.env.example to cmd/.env
- Place your weatherAPI apiKey in cmd/.env file.
- Run `docker-compose up --build`
- Default port will be 8080

### Tests
- You can run locally on http://localhost:8080/31275000
- You can test the app running on Google Cloud Run in the url: https://goexpert-temperature-postal-code-jdhq42yxaa-rj.a.run.app/31275000