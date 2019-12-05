package rules

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/yaien/structure/core"
)

// Pattern get a pattern validation func
func Pattern(params []string) (core.ValidateFunc, error) {
	if len(params) != 1 {
		return nil, errors.New("pattern validation needs to have 1 param")
	}
	pattern := params[0]
	regex, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return func(item *core.Item) error {
		value := item.Value.String()
		if regex.MatchString(value) {
			return nil
		}
		return fmt.Errorf("'%s' doesn't match the pattern '%s'", value, pattern)
	}, nil
}
