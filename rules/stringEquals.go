package rules

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
)

func StringEquals(str string) validation.RuleFunc {
	return func(value interface{}) error {
		s, _ := value.(string)
		if s != str {
			return errors.New("unexpected string")
		}
		return nil
	}
}
