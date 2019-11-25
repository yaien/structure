package core

import "reflect"

// Item validation data for the current field
type Item struct {
	Source  interface{}
	Value   reflect.Value
	Field   string
	Message string
}

// GetValidateFunc returns a validation function
type GetValidateFunc = func(params []string) (ValidateFunc, error)

// ValidateFunc validate an item
type ValidateFunc = func(item *Item) error

// Result of validation function
type Result struct {
	Failed bool
	Errors map[string][]string
}
