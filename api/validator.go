package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/hl/hl_bank/util"
)

var validateCurrency = func(fieldLevel validator.FieldLevel) bool {

	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
