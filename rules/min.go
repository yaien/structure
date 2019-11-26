package rules

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/yaien/structure/core"
)

// Min validate if a value is greater or equal than a min value, for strings it compares the length
func Min(params []string) (core.ValidateFunc, error) {
	min, err := strconv.ParseFloat(params[0], 64)
	if err != nil {
		return nil, err
	}
	return func(item *core.Item) error {
		switch item.Value.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value := item.Value.Int()
			if value < int64(min) {
				return fmt.Errorf("should be greater or equal than %d", int64(min))
			}
		case reflect.Float32, reflect.Float64:
			value := item.Value.Float()
			if value < min {
				return fmt.Errorf("should be greater or equal than %f", min)
			}
		case reflect.String:
			value := item.Value.String()
			minlength := int(min)
			if len(value) < minlength {
				return fmt.Errorf("should have a length greater or equal than %d", minlength)
			}
		}
		return nil
	}, nil
}
