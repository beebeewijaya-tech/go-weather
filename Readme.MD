# Project Overview


## Goals
The goals of this project is by able to serve a public API that can be consumed by just passing latitude and longitude and we will return the weather condition.


## How to run
1. First copy `env/config.sample.json` into `env/config.json`
2. Fill the value from API Key from [Openweathermap API](https://openweathermap.org/api)
3. After filling the value try to run `go mod tidy`
4. Afterward you can run `go run cmd/api/main.go`


## Architecture Diagram
You can check our [Architecture Diagram here](./docs/Architecture.MD), 

We love to elaborate everything that been use ( tech stack ) on this project.


## API Docs

### Weather Search

#### Request
---

```http
POST /api/v1/weather/search
```

Body Request:
```json
{
    "lat": 10.99,
    "lon": 44.34
}
```


#### Responses
---
```javascript
{
    "coord": {
        "lat": 10.99,
        "lon": 44.34
    },
    "weather": [
        {
            "id": 802,
            "main": "Clouds",
            "description": "scattered clouds",
            "icon": "03d"
        }
    ],
    "main": {
        "temp": 303.6,
        "temp_min": 303.6,
        "temp_max": 303.6,
        "humidity": 80
    }
}
```