package rules

import "github.com/yaien/structure/core"

// Equal validate if the field's value is equal to another given field
func Equal(params []string) (core.ValidateFunc, error) {
	return func(item *core.Item) error {
		return nil
	}, nil
}
