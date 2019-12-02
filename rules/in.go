package rules

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/yaien/structure/core"
)

// In returns a validate function that checks if a given value is in the given options
func In(params []string) (core.ValidateFunc, error) {

	in := func(value string) bool {
		for _, option := range params {
			if option == value {
				return true
			}
		}
		return false
	}

	return func(item *core.Item) error {
		switch item.Value.Kind() {
		case reflect.String:
			value := item.Value.String()
			if in(value) {
				return nil
			}
			return fmt.Errorf("'%s' is not a valid option, value must be in %v", value, params)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value := strconv.FormatInt(item.Value.Int(), 10)
			if in(value) {
				return nil
			}
			return fmt.Errorf("'%s' is not a valid option, value must be in %v", value, params)
		}
		return nil
	}, nil
}
