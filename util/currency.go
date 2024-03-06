package util

const (
	USD  = "USD"
	EURO = "EURO"
	SGD  = "SGD"
)

func IsSupportedCurrency(currency string) bool {

	switch currency {
	case USD, EURO, SGD:
		return true
	}
	return false
}
