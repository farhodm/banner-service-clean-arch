package helpers

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

func ValidationsError(err error, validationErrors map[string]string) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, fieldError := range ve {
			field := fieldError.Field()
			tag := fieldError.Tag()
			validationErrors[field] = fmt.Sprintf("Field '%s' is %s", field, tag)
		}
	}
}
