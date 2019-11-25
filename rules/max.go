package rules

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/yaien/structure/core"
)

func Max(params []string) (core.ValidateFunc, error) {
	max, err := strconv.ParseFloat(params[0], 64)
	if err != nil {
		return nil, err
	}
	return func(item *core.Item) error {
		switch item.Value.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			value := item.Value.Int()
			if value > int64(max) {
				return fmt.Errorf("should be greater or equal than %d", int64(max))
			}
		case reflect.Float32, reflect.Float64:
			value := item.Value.Float()
			if value > max {
				return fmt.Errorf("should be greater or equal than %f", max)
			}
		case reflect.String:
			value := item.Value.String()
			maxlength := int(max)
			if len(value) > maxlength {
				return fmt.Errorf("should have a length greater or equal than %d", maxlength)
			}
		}
		return nil
	}, nil
}
