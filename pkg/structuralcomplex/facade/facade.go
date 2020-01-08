package facade

type CurrentWeatherDataReceiver interface {
	GetByCityAndCountryCode(city, countryCode string) (Weather, error)

	GetByGeoCoordinates(lat, lon float32) (Weather, error)
}

type Weather struct {
	ID   int    `json:id`
	Name string `json:name`
}
