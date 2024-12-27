package fetchData

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/medwards2009/waittimewizard-api/dataModels"
	"github.com/medwards2009/waittimewizard-api/graph/model"
)

func FetchDestinations() (*dataModels.Destinations, error) {
	resp, err := http.Get("https://api.themeparks.wiki/v1/destinations")

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var responseObject *dataModels.Destinations
	json.Unmarshal(responseData, &responseObject)

	return responseObject, nil
}

func FetchLiveData(id string) (*model.LiveData, error) {
	url := fmt.Sprintf("https://api.themeparks.wiki/v1/entity/%s/live", id)
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var responseObject *model.LiveData
	json.Unmarshal(responseData, &responseObject)

	return responseObject, nil
}
