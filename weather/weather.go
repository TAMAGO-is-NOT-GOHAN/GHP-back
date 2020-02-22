package weather

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type OWeatherMap struct {
	List []struct {
		Dt_txt  string `json:"dt_txt"`
		Weather []struct {
			Main string `json:"main"`
		} `json:"weather"`
	} `json:"list"`
}

func GetWeather() OWeatherMap {
	var weather OWeatherMap

	url := `http://api.openweathermap.org/data/2.5/forecast?units=metric&q=Naha&APPID=` + os.Getenv("OWM_TOKEN")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := new(http.Client)

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &weather)
	if err != nil {
		log.Fatal(err)
	}

	return weather
}
