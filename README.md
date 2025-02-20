# Country Information Service

A REST web application built in Go that provides information about countries and their historical population data.

## Overview

This service aggregates data from multiple third-party APIs to provide consolidated information about countries, including general details and population history. It's designed to demonstrate effective API integration and data transformation.

### Features
- Country general information (name, continents, population, languages, etc.)
- Historical population data with optional year range filtering
- Service status monitoring

## Setup

### Prerequisites
- Go 1.23 or higher
- Internet connection (to access third-party APIs)

### Installation

Clone the repository:
```bash
git clone <repository-url>
cd countryinfo-api
```

### Building

To compile the application:
```properties
go build -o app
```

### Running

To run the application:
```properties
./app
```

By default, the service runs on port 8080. You can access it at http://localhost:8080/countryinfo/v1/

## API Usage

Base URL:
```
/countryinfo/v1/
```

### Endpoints

#### 1. Country Information
```
GET /countryinfo/v1/info/{two_letter_country_code}{?limit=10}
```

Returns general information for a specified country.

Parameters:
- `two_letter_country_code`: ISO 3166-2 two-letter country code (required)
- `limit`: Optional parameter to limit the number of cities returned (default: all cities)

Example request:
```
GET /countryinfo/v1/info/no
GET /countryinfo/v1/info/no?limit=5
```

Example response:
```json
{
  "name": "Norway",
  "continents": ["Europe"],
  "population": 4700000,
  "languages": {"nno":"Norwegian Nynorsk","nob":"Norwegian Bokmål","smi":"Sami"},
  "borders": ["FIN","SWE","RUS"],
  "flag": "https://flagcdn.com/w320/no.png",
  "capital": "Oslo",
  "cities": ["Abelvaer","Adalsbruk","Adland","Agotnes","Agskardet"]
}
```

#### 2. Country Population History
```
GET /countryinfo/v1/population/{two_letter_country_code}{?limit={startYear-endYear}}
```

Returns historical population data for a specified country.

Parameters:
- `two_letter_country_code`: ISO 3166-2 two-letter country code (required)
- `limit`: Optional parameter to filter population data by year range (format: startYear-endYear)

Example request:
```
GET /countryinfo/v1/population/no
GET /countryinfo/v1/population/no?limit=2010-2015
```

Example response:
```json
{
   "mean": 5044396,
   "values": [
     {"year":2010,"value":4889252},
     {"year":2011,"value":4953088},
     {"year":2012,"value":5018573},
     {"year":2013,"value":5079623},
     {"year":2014,"value":5137232},
     {"year":2015,"value":5188607}
   ]
}
```

#### 3. Service Status
```
GET /countryinfo/v1/status/
```

Returns information about the service's dependencies and uptime.

Example response:
```json
{
   "countriesnowapi": 200,
   "restcountriesapi": 200,
   "version": "v1",
   "uptime": 3600
}
```

## Architecture

This service integrates with two external APIs:
1. CountriesNow API (http://129.241.150.113:3500/api/v0.1/)
2. REST Countries API (http://129.241.150.113:8080/v3.1/)

The service processes and transforms data from these sources to provide a unified interface for country information.

## Deployment

The service is deployed on Render at: [https://countryinfo-api.onrender.com](https://countryinfo-api.onrender.com)

## Development Notes

- Error handling follows REST best practices
- Requests to third-party APIs are made as small as possible
- The service uses only Go standard library to minimize dependencies

## License

[MIT License]
