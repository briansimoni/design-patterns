package facade

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type CurrentWeatherDataRetriever interface {
	GetByCityAndCountryCode(city, countryCode string) (Weather, error)
	GetByGeoCoordinates(lat, lon float32) (Weather, error)
}

type Weather struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Cod   int    `json:"cod"`
	Coord struct {
		Lon float32 `json:"lon"`
		Lat float32 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float32 `json:"temp"`
		Pressure float32 `json:"pressure"`
		Humidity float32 `json:"humidity"`
		TempMin  float32 `json:"temp_min"`
		TempMax  float32 `json:"temp_max"`
	} `json:"main"`
	Wind struct {
		Speed float32 `json:"speed"`
		Deg   float32 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Rain struct {
		ThreeHours float32 `json:"3h"`
	} `json:"rain"`
	Dt  uint32 `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float32 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
}

type CurrentWeatherData struct {
	APIKey string
}

func (p *CurrentWeatherData) responseParser(body io.Reader) (*Weather, error) {
	w := new(Weather)
	err := json.NewDecoder(body).Decode(w)
	if err != nil {
		return nil, err
	}
	return w, nil
}

func (c *CurrentWeatherData) GetByGeoCoordinates(lat, lon float32) (*Weather, error) {
	return c.doRequest(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?lat=%f&lon=%f&APPID=%s", lat, lon, c.APIKey))
}

func (c *CurrentWeatherData) GetByCityAndCountryCode(city, countryCode string) (*Weather, error) {
	return c.doRequest(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weatherq=%s,%s&APPID=%s", city, countryCode, c.APIKey))
}

func (o *CurrentWeatherData) doRequest(uri string) (*Weather, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		byt, err := ioutil.ReadAll(response.Body)
		if err == nil {
			err = fmt.Errorf("%s", string(byt))
		}
		err = fmt.Errorf("Status code was %d, aborting. Error message was: %s", response.StatusCode, err.Error())
		return nil, err
	}
	weather, err := o.responseParser(response.Body)
	if err != nil {
		return nil, err
	}
	return weather, nil
}
