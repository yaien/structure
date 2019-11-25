package structure

import (
	"errors"
	"testing"

	"github.com/yaien/structure/core"
)

func TestValidate(t *testing.T) {
	type schema struct {
		name  string `struct:"min:2"`
		email string `struct:"email"`
	}
	message := "min validation failed"
	p := &parser{
		validations: map[string]core.GetValidateFunc{
			"min": func(params []string) (core.ValidateFunc, error) {
				return func(item *core.Item) error {
					return errors.New(message)
				}, nil
			},
			"email": func(params []string) (core.ValidateFunc, error) {
				return func(item *core.Item) error {
					return nil
				}, nil
			},
		},
	}

	validator, err := p.Build(schema{})

	if err != nil {
		t.Fatal(err)
	}

	result := validator.Validate(schema{name: "n", email: "email"})

	if !result.Failed {
		t.Errorf("Expected validation result.Failed to be true, received: '%v'", result.Failed)
	}

	msg := result.Errors["name"][0]
	if msg != message {
		t.Errorf("Expect to receive an error message for name to be '%s', received: '%s'", message, msg)
	}

}
