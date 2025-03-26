# go zipcode weather - otel

## Local environment

You will need to create an API key from [Weather API](https://www.weatherapi.com/) to use in the environment variable `WEATHER_API_KEY`.

### 1. Running with docker
### Build image
```bash
docker build -t egon89/go-zipcode-weather-otel .
```

### Run container with environment variables
```bash
docker run --name gzweather -p 8080:8080 \
  -e WEATHER_API_BASE_URL=http://api.weatherapi.com \
  -e WEATHER_API_KEY=weather_api_key \
  -e VIA_CEP_BASE_URL=https://viacep.com.br/ws \
  egon89/go-zipcode-weather-otel:latest
```
---

### 2. Running using golang locally
Create a `.env` file from `.env.example` and fill the variables.
```bash
cp .env.example .env
```

> Set up the `WEATHER_API_KEY` environment variable with your API key.

To run the application, execute the following command in the root directory of the project:
```bash
go run cmd/main.go
```

### Usage
```bash
curl --location 'http://localhost:8080/weather/{:zipcode}' --verbose
```
