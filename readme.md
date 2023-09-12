# WeatherAPI Golang Client

![GitHub issues](https://img.shields.io/github/issues/marcelomatz/weather-api)
![GitHub pull requests](https://img.shields.io/github/issues-pr/marcelomatz/weather-api)
![GitHub](https://img.shields.io/github/license/marcelomatz/weather-api)

This project uses a weather API to fetch data about current weather in a specific location.

## Project setup
- A viable Go environment setup is required.
- Project uses Go SDK version 1.20.5 and Go programming language version 1.20.

## Dependencies
- `github.com/joho/godotenv`

## Project Structure

The project contains main Go file which includes these components:
- `Location`: A struct that holds data about the location namely name, region, country and the local time.
- `Condition`: A struct holds data about the current weather condition.
- `Current`: A struct holds data about the current weather.
- `Weather`: A struct holds data about the location and current weather.

Main function fetches the weather data using an API key and location parameters (city, state and country). It then decodes the JSON API response into the defined text data.

## Setup process
1. You should have a `.env` file in root of your project directory. It should have a `WEATHER_API_TOKEN` field which is the API token for the weather API.

IMPORTANT: Don't forget to replace YOUR_API_TOKEN with your actual API token.
2. Select the city, state and country you want to fetch the current weather of in the main() function.
3. Run go run . to execute the project.

```go
city := "nova-lima"    // Replace 'nova-lima' with your city
state := "mg"          // Replace 'mg' with your state
country := "br"        // Replace 'br' with your country code
```
The resulting output will display weather details of the selected location.

```shell
Weather in <City>, <State>, <Country>
Temperature: <Temperature>°C
Condition: <Weather Condition>
Last updated: <Last Updated Time>
Local time: <Local Time>

```
All the fields enclosed in <> will be replaced with the fetched weather data.

Feel free to explore and improve upon this project.