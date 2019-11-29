package rules

import (
	"fmt"
	"regexp"

	"github.com/yaien/structure/core"
)

// Email validation
func Email(params []string) (core.ValidateFunc, error) {
	regex := regexp.MustCompile("[a-zA-Z0-9]+@[a-zA-Z0-9]+\\.[a-zA-Z]+")
	return func(item *core.Item) error {
		str := item.Value.String()
		if regex.MatchString(str) {
			return nil
		}
		return fmt.Errorf("'%s' is not a valid email", str)
	}, nil
}
