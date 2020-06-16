package model

import "github.com/go-ozzo/ozzo-validation/v4"

func requiredIf(val bool) validation.RuleFunc {
	return func(value interface{}) error {
		if val {
			return validation.Validate(value, validation.Required)
		}
		return nil
	}
}
