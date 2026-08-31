package finance

import "math"

type Balance struct {
	Principal float64
	Rate      float64
	Period    int64
}

func (b *Balance) PayBack() float64 {
	return b.Principal / math.Pow(1+b.Rate, float64(b.Period))
}
