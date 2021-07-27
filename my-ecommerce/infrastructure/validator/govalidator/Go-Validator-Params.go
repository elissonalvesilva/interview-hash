package govalidator

import "github.com/asaskevich/govalidator"

type Validator struct {}

func NewValidator() *Validator {
	return &Validator{}
}

func (validator *Validator) ValidateRequestParams(params interface{}) error {
	_, errorValidationParam := govalidator.ValidateStruct(params)
	return errorValidationParam
}
