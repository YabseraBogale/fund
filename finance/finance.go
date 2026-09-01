package finance

import "math"

type PresentValue struct {
	CashFlow float64
	Rate     float64
}

func (p *PresentValue) SimplePV(timePeriod int) float64 {
	return p.CashFlow / math.Pow(1+p.Rate, float64(timePeriod))
}
