package rules

import (
	"reflect"
	"testing"

	"github.com/yaien/structure/core"
)

func TestEmail(t *testing.T) {
	testcases := []struct {
		email string
		valid bool
	}{
		{email: "stevensonx", valid: false},
		{email: "stiv@gmail", valid: false},
		{email: "user@live.com", valid: true},
		{email: "1233@live", valid: false},
		{email: "smrx1234@mcondor.net", valid: true},
		{email: "field@$%&$%&.es", valid: false},
		{email: "#@¢¢@kfjdf.es", valid: false},
		{email: "field@field+com", valid: false},
		{email: "sdfkj%live.com", valid: false},
		{email: "home@home.", valid: false},
		{email: "@live.com", valid: false},
		{email: "test@.com", valid: false},
	}
	validate, err := Email([]string{})
	if err != nil {
		t.Fatal(err)
	}
	for _, testcase := range testcases {
		item := &core.Item{Value: reflect.ValueOf(testcase.email)}
		err := validate(item)
		valid := err == nil
		if valid != testcase.valid {
			t.Errorf("Expected email '%s' validation result to be '%v', received: '%v'", testcase.email, testcase.valid, valid)
		}
	}
}
