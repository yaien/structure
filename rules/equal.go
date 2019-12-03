package rules

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/yaien/structure/core"
)

// Equal validate if the field's value is equal to another given field
func Equal(params []string) (core.ValidateFunc, error) {
	if len(params) != 1 {
		return nil, errors.New("equal rule needs to have one param")
	}

	return func(item *core.Item) error {
		field := reflect.ValueOf(item.Source).FieldByName(params[0]).Interface()
		switch item.Value.Kind() {
		case reflect.String:
			value := item.Value.String()
			if value != field.(string) {
				return fmt.Errorf("should be equal to the '%s' field", params[0])
			}
			return nil
		case reflect.Int:
			value := int(item.Value.Int())
			if value != field.(int) {
				return fmt.Errorf("should be equal to the '%s' field", params[0])
			}
			return nil
		case reflect.Bool:
			value := item.Value.Bool()
			if value != field.(bool) {
				return fmt.Errorf("should be equal to the '%s' field", params[0])
			}
			return nil
		}
		return fmt.Errorf("unsupported type %v", item.Value.Type())
	}, nil
}
