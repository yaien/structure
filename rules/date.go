package rules

import (
	"fmt"
	"reflect"
	"time"

	"github.com/yaien/structure/core"
)

// Date validate if a string is a correctly formated date string
func Date(params []string) (core.ValidateFunc, error) {
	layout := params[0]
	return func(item *core.Item) error {
		if item.Value.Kind() != reflect.String {
			return fmt.Errorf("%v is not a valid date string", item.Value)
		}

		if _, err := time.Parse(layout, item.Value.String()); err != nil {
			return fmt.Errorf("%v is not a valid date string in layout %s", item.Value, layout)
		}

		return nil

	}, nil
}
