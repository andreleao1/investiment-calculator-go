package centralbankclient

import (
	"encoding/json"
	"errors"
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

func GetDataFromCentralBank(bcProduct string) (*float64, error) {
	response, err := makeRequest(&bcProduct)

	if err != nil || response == "" {
		return nil, err
	}

	yearPercentageFee := parseToYearRate(extractDailyRate(response))

	return &yearPercentageFee, nil
}

func makeRequest(bcProduct *string) (string, error) {
	bcUrl, err := getBcUrlRequest(bcProduct)

	if err != nil {
		return "", err
	}

	response, err := http.Get(bcUrl)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", err
	}

	var selicData []SelicData
	if err := json.NewDecoder(response.Body).Decode(&selicData); err != nil {
		return "", err
	}

	if len(selicData) == 0 {
		return "", errors.New("No data found")
	}

	return selicData[0].Value, nil
}

func getBcUrlRequest(bcProduct *string) (string, error) {
	bcUrl, err := readjson.GetValueByKey("bcUrl")

	if err != nil {
		return "", err
	}

	bcProductUrl, err := getProductUrl(bcProduct)

	if err != nil {
		return "", err
	}

	return bcUrl + bcProductUrl, nil
}

func getProductUrl(bcProduct *string) (string, error) {
	productUrl, err := readjson.GetValueByKey(*bcProduct)

	if err != nil {
		return "", err
	}

	return productUrl, nil
}

func parseToYearRate(dailyRate float64, err error) float64 {
	if err != nil {
		return 0.0
	}

	return (math.Pow((1+(dailyRate/100)), constants.DAYS_PER_YEAR) - 1) * 100
}

func extractDailyRate(stringSelicRate string) (float64, error) {
	todaySelicRate, err := strconv.ParseFloat(stringSelicRate, 64)
	if err != nil {
		return 0.0, err
	}

	return todaySelicRate, nil
}
