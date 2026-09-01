package finance

import (
	"errors"
	"math"
)

type PresentValue struct {
	CashFlow float64
	Rate     float64
}

func (p *PresentValue) SimplePV(timePeriod int) float64 {
	return p.CashFlow / math.Pow(1+p.Rate, float64(timePeriod))
}

func (p *PresentValue) GrowingPerpetuity(growth float64) (float64, error) {
	if p.Rate < growth {
		return -1, errors.New("garbage")
	}
	return p.CashFlow / (p.Rate - growth), nil
}
