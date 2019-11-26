package rules

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"

	"github.com/yaien/structure/core"
)

// Between validate if a value is between a min and max value, for strings it compares the length
func Between(params []string) (core.ValidateFunc, error) {
	if len(params) != 2 {
		return nil, errors.New("between rule should receive 2 params: min,max")
	}
	min, err := strconv.ParseFloat(params[0], 64)
	if err != nil {
		return nil, err
	}
	max, err := strconv.ParseFloat(params[1], 64)
	if err != nil {
		return nil, err
	}
	return func(item *core.Item) error {
		switch item.Value.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value := item.Value.Int()
			if value > int64(max) || value < int64(min) {
				return fmt.Errorf("should be between in %d and %d", int64(min), int64(max))
			}
		case reflect.Float32, reflect.Float64:
			value := item.Value.Float()
			if value > max || value < min {
				return fmt.Errorf("should be between in %f and %f", min, max)
			}
		case reflect.String:
			value := item.Value.String()
			maxlength := int(max)
			minlength := int(min)
			length := len(value)
			if (length > maxlength) || (length < minlength) {
				return fmt.Errorf("should have a length between in %d and %d", minlength, maxlength)
			}
		}
		return nil
	}, nil
}
