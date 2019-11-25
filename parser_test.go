package structure

import (
	"reflect"
	"strings"
	"testing"

	"github.com/yaien/structure/core"
)

func TestParseRule(t *testing.T) {
	testcases := []string{"min:2", "between:2,3", "required"}
	p := &parser{}
	for _, testcase := range testcases {
		args := strings.Split(testcase, ":")
		params := make([]string, len(args)-1)
		rule := p.rule(testcase)

		if len(args) > 1 {
			params = strings.Split(args[1], ",")
		}

		if rule.name != args[0] {
			t.Errorf("expected to receive rule name equal to '%s' for testcase '%s', received: '%s'", args[0], testcase, rule.name)
		}

		if len(params) != len(rule.params) {
			t.Fatalf("expected to receive rule params to '%s' for testcase '%s', received '%s'", params, testcase, rule.params)
		}

		for i := 0; i < len(params); i++ {
			if params[i] != rule.params[i] {
				t.Fatalf("expected to receive rule params to '%s' for testcase '%s', received '%s'", params, testcase, rule.params)
			}
		}
	}
}

func TestParseStructure(t *testing.T) {
	type schema struct {
		name  string `struct:"min:2"`
		email string `struct:"email"`
	}
	s := schema{}
	validations := map[string]core.GetValidateFunc{
		"min": func(params []string) (core.ValidateFunc, error) {
			return func(item *core.Item) error {
				return nil
			}, nil
		},
		"email": func(params []string) (core.ValidateFunc, error) {
			return func(item *core.Item) error {
				return nil
			}, nil
		},
	}
	parser := &parser{validations}
	structure, err := parser.structure(s)

	if err != nil {
		t.Errorf("Failed to build structure %v", err)
	}

	tp := reflect.TypeOf(s)
	for i := 0; i < tp.NumField(); i++ {
		field := tp.Field(i)
		rules, found := structure[field.Name]
		if !found {
			t.Errorf("expected structure to have field '%s', it wasn't found", field.Name)
		}
		tag := field.Tag.Get("struct")
		length := len(strings.Split(tag, "|"))
		if len(rules) != length {
			t.Errorf("expected rules to be %d for tag '%s', received: %d", length, tag, len(rules))
		}

	}

}

func TestFailedParsedStructure(t *testing.T) {
	type schema struct {
		name string `struct:"min:2"`
	}
	parser := &parser{}
	_, err := parser.structure(schema{})
	if err == nil {
		t.Fatal("expected parse structure to fail")
	}

	message := "validation rule 'min' doesn't exists"
	if err.Error() != message {
		t.Errorf("expected error message to be '%s', received: '%s'", message, err.Error())
	}
}
