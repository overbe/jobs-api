package validate

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	*validator.Validate
}

func Init() *Validator {
	v := validator.New()

	//TODO: We can add any validating rules

	return &Validator{v}
}
