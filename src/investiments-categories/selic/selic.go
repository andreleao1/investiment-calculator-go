package selic

import (
	"fmt"
	"math"
	"strconv"

	"agls.com.br/src/constants"
	readjson "agls.com.br/src/utils/json"
)

type Selic struct {
	InitialContribution float64
	MonthlyContribution float64
	InvestimentTime     float64
	InvestimentRate     float64
	FutureValue         float64
}

func New(initialContribution, monthlyContribution, investimentTime float64) *Selic {
	investimentRateFromFile, err := readjson.GetValueByKey("selic")
	var investmentRate float64
	if err == nil {
		investmentRate, err = strconv.ParseFloat(investimentRateFromFile, 64)

		if err != nil {
			defineDefaultInvestimentRate(&investmentRate)
		}
	} else {
		defineDefaultInvestimentRate(&investmentRate)
	}

	return &Selic{
		InitialContribution: initialContribution,
		MonthlyContribution: monthlyContribution,
		InvestimentTime:     investimentTime,
		InvestimentRate:     investmentRate,
		FutureValue:         0,
	}
}

func (s *Selic) Calculate() float64 {
	mounthlyRate := getMonthlyRate(&s.InvestimentRate)
	periods := calculateTotalPeriods(&s.InvestimentTime)
	s.FutureValue += s.InitialContribution * math.Pow(1+mounthlyRate, periods)
	s.FutureValue += s.MonthlyContribution * (math.Pow(1+mounthlyRate, periods) - 1) / mounthlyRate

	return s.FutureValue
}

func GetCurrentSelicRate() float64 {
	investimentRateFromFile, err := readjson.GetValueByKey("selic")
	var investmentRate float64
	if err == nil {
		investmentRate, err = strconv.ParseFloat(investimentRateFromFile, 64)

		if err != nil {
			defineDefaultInvestimentRate(&investmentRate)
		}
	} else {
		defineDefaultInvestimentRate(&investmentRate)
	}

	return investmentRate
}

func getMonthlyRate(anualInterest *float64) float64 {
	return *anualInterest / 100 / constants.PERYOD_PER_YEAR
}

func calculateTotalPeriods(investimentTime *float64) float64 {
	return constants.PERYOD_PER_YEAR * *investimentTime
}

func defineDefaultInvestimentRate(variable *float64) {
	fmt.Println("Error to get investment rate from file, using default value 10.5%")
	*variable = 10.5
}
