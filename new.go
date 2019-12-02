package structure

import (
	"github.com/yaien/structure/core"
	"github.com/yaien/structure/rules"
)

// New return the default validation parser
func New() Parser {
	return &parser{validations: map[string]core.GetValidateFunc{
		"min":      rules.Min,
		"max":      rules.Max,
		"between":  rules.Between,
		"date":     rules.Date,
		"required": rules.Required,
		"email":    rules.Email,
		"in":       rules.In,
	}}
}
