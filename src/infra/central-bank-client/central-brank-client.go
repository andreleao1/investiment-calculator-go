package centralbankclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"agls.com.br/constants"
	readjson "agls.com.br/utils/json"
)

type SelicData struct {
	Data  string `json:"data"`
	Value string `json:"valor"`
}

func GetSelicDataFromCentralBank() (*float64, error) {
	bcUrl, err := readjson.GetValueByKey("bcUrl")

	if err != nil {
		return nil, fmt.Errorf("failed to get bcUrl: %w", err)
	}

	response, err := http.Get(bcUrl) //http.Get("http://www.google.com.br")
	if err != nil {
		return nil, fmt.Errorf("failed to get response from Central Bank: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", response.StatusCode)
	}

	var selicData []SelicData
	if err := json.NewDecoder(response.Body).Decode(&selicData); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	if len(selicData) == 0 {
		return nil, errors.New("no data found in response")
	}

	todaySelicRate, err := strconv.ParseFloat(selicData[0].Value, 64)

	if err != nil {
		return nil, fmt.Errorf("failed to parse float: %w", err)
	}

	yearSelicRate := (math.Pow((1+(todaySelicRate/100)), constants.DAYS_PER_YEAR) - 1) * 100

	return &yearSelicRate, nil
}
