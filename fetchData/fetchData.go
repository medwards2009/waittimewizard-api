package fetchData

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Destinations struct {
	Destinations []Destination `json:"destinations"`
}

type Park struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Destination struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Parks []Park `json:"parks"`
}

func FetchDestinations() (Destinations, error) {
	resp, err := http.Get("https://api.themeparks.wiki/v1/destinations")

	if err != nil {
		fmt.Println(err.Error())
		return Destinations{}, err
	}

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return Destinations{}, err
	}

	var responseObject Destinations
	json.Unmarshal(responseData, &responseObject)

	return responseObject, nil
}

// func fetchLiveData(id string) {
// TODO: add fetch with id and will need to build all
// return types of course
// }
