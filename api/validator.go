package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/vivekcode101/simplebanks-golang/util"
)

var validCurrency validator.Func = func(FieldLevel validator.FieldLevel) bool {
	if currency, ok := FieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
