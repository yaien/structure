package rules

import (
	"fmt"
	"regexp"

	"github.com/yaien/structure/core"
)

/*AlphaNum return alpha's validation func that validate if a
string only has letters and numbers
*/
func AlphaNum(params []string) (core.ValidateFunc, error) {
	regex := regexp.MustCompile("(?i)^[a-z0-9]+$")
	return func(item *core.Item) error {
		value := item.Value.String()
		if regex.MatchString(value) {
			return nil
		}
		return fmt.Errorf("%s isn't a valid alpha string", value)
	}, nil
}
