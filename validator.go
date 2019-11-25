package structure

import (
	"reflect"

	"github.com/yaien/structure/core"
)

// Validator interface
type Validator interface {
	Validate(data interface{}) *core.Result
}

type validator struct {
	schema    interface{}
	structure map[string][]core.ValidateFunc
}

func (v *validator) Validate(data interface{}) *core.Result {
	messages := make(map[string][]string)
	failed := false
	schemaType := reflect.TypeOf(v.schema)
	dataType := reflect.TypeOf(data)
	values := reflect.ValueOf(data)

	if !dataType.AssignableTo(schemaType) {
		return &core.Result{Failed: true}
	}

	for i := 0; i < dataType.NumField(); i++ {
		var errors []string
		field := dataType.Field(i)
		item := &core.Item{
			Source: data,
			Value:  values.FieldByName(field.Name),
			Field:  field.Name,
		}
		validateFuncs := v.structure[field.Name]
		for _, validate := range validateFuncs {
			err := validate(item)
			if err != nil {
				failed = true
				errors = append(errors, err.Error())
			}
		}
		messages[field.Name] = errors
	}

	return &core.Result{Failed: failed, Errors: messages}
}
