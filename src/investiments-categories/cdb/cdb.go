package cdb

import "agls.com.br/constants"

type CDB struct {
	InitialContribution float64
	MonthlyContribution float64
	InvestimentTime     float64
	InvestimentRate     float64
	FutureValue         float64
}

func New(initialContribution, monthlyContribution, investimentTime, investimentRate float64) *CDB {

	if investimentRate != 0 {
		investimentRate = constants.DEAFULT_CDI_RATE * (investimentRate / 100)
	} else {
		investimentRate = constants.DEAFULT_CDI_RATE * (constants.DEAFULT_CBD_RATE / 100)
	}

	return &CDB{
		InitialContribution: initialContribution,
		MonthlyContribution: monthlyContribution,
		InvestimentTime:     investimentTime,
		InvestimentRate:     investimentRate,
		FutureValue:         0,
	}
}

func (c *CDB) Calculate() float64 {
	return 1.0
}
