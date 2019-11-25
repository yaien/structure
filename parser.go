package structure

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/yaien/structure/core"
)

type rule struct {
	name   string
	params []string
}

// Parser a structure validation instance
type Parser interface {
	Build(schema interface{}) (Validator, error)
	Extend(name string, rule core.GetValidateFunc)
}

type parser struct {
	validations map[string]core.GetValidateFunc
}

func (p *parser) rule(validation string) *rule {
	args := strings.Split(validation, ":")
	length := len(args)
	name := args[0]
	params := make([]string, length-1)
	if length > 1 {
		params = strings.Split(args[1], ",")
	}
	return &rule{name, params}
}

func (p *parser) structure(schema interface{}) (map[string][]core.ValidateFunc, error) {
	structure := make(map[string][]core.ValidateFunc)
	schemaType := reflect.TypeOf(schema)
	for i := 0; i < schemaType.NumField(); i++ {
		field := schemaType.Field(i)
		var validationFuncs []core.ValidateFunc
		for _, validation := range strings.Split(field.Tag.Get("struct"), "|") {
			rule := p.rule(validation)
			maker, found := p.validations[rule.name]
			if !found {
				return nil, fmt.Errorf("validation rule '%s' doesn't exists", rule.name)
			}
			validationFunc, err := maker(rule.params)

			if err != nil {
				return nil, err
			}

			validationFuncs = append(validationFuncs, validationFunc)
		}
		structure[field.Name] = validationFuncs
	}

	return structure, nil
}

func (p *parser) Build(schema interface{}) (Validator, error) {
	structure, err := p.structure(schema)
	if err != nil {
		return nil, err
	}
	return &validator{schema, structure}, nil
}

func (p *parser) Extend(name string, builder core.GetValidateFunc) {
	p.validations[name] = builder
}
