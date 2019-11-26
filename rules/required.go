package rules

import (
	"fmt"
	"reflect"

	"github.com/yaien/structure/core"
)

// Required validate if the value is a thruthy value
func Required(params []string) (core.ValidateFunc, error) {
	return func(item *core.Item) error {
		switch item.Value.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if item.Value.Int() == 0 {
				return fmt.Errorf("%s is required", item.Field)
			}
		case reflect.String:
			if len(item.Value.String()) == 0 {
				return fmt.Errorf("%s is required", item.Field)
			}
		case reflect.Float32, reflect.Float64:
			if item.Value.Float() == float64(0) {
				return fmt.Errorf("%s is required", item.Field)
			}
		}

		if !item.Value.IsValid() {
			return fmt.Errorf("%s is required", item.Field)
		}
		return nil
	}, nil
}
