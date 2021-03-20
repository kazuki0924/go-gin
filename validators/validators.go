package validators

import (
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

// ValidateCoolTitle func
func ValidateCoolTitle(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "Cool")
}
