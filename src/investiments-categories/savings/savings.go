package savings

import (
	"fmt"
	"strconv"

	"agls.com.br/constants"
	"agls.com.br/investiments-categories/selic"
	readjson "agls.com.br/utils/json"
)

type Savings struct {
	InitialContribution float64
	MonthlyContribution float64
	InvestimentTime     float64
	InvestimentRate     float64
	ReferenceRate       float64
	FutureValue         float64
}

func New(initialContribution, monthlyContribution, investimentTime float64) *Savings {
	return &Savings{
		InitialContribution: initialContribution,
		MonthlyContribution: monthlyContribution,
		InvestimentTime:     investimentTime,
		InvestimentRate:     getSavingRate(),
		FutureValue:         0,
	}
}

func (s *Savings) Calculate() {
	s.FutureValue = s.InitialContribution

	for i := 0; i < int(s.InvestimentTime*constants.PERYOD_PER_YEAR); i++ {
		s.FutureValue += s.MonthlyContribution
		s.FutureValue += s.FutureValue * s.InvestimentRate
	}
}

func getSavingRate() float64 {
	referenceRate := getReferenceRate()
	selicRate := selic.GetCurrentSelicRate()

	if selicRate > 8.5 {
		return 0.5/100.0 + referenceRate/100.0
	} else {
		return (70.0 / 100.0 * (selicRate / 12.0)) / 100.0
	}
}

func getReferenceRate() float64 {
	referenceRateFromFile, err := readjson.GetValueByKey("referenceRate")

	if err != nil {
		return 0.0
	}

	referenceRate, err := strconv.ParseFloat(referenceRateFromFile, 64)

	if err != nil {
		fmt.Println("Error to get reference rate, default value 0.0 will be used")
		return 0.0
	}

	return referenceRate
}
