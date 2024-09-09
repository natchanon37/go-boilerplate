package utils

import (
	"log"

	"github.com/go-playground/validator/v10"
)

var v *validator.Validate

func decimal2(fl validator.FieldLevel) bool {
	// Extract the float value from the field
	myFloat := fl.Field().Float()

	// Check if the float has exactly two decimal places
	return myFloat == ConvertTo2Decimal(myFloat)
}
func init() {
	v = validator.New()
	// for custom validator
	if err := v.RegisterValidation("decimal2", decimal2); err != nil {
		log.Fatalf("Error registering custom validator: %v", err)
	}
}

func ValidateStruct(s interface{}) error {
	return v.Struct(s)
}
