package rules

import (
	"reflect"
	"testing"

	"github.com/yaien/structure/core"
)

func TestDate(t *testing.T) {
	testcases := []struct {
		layout string
		value  interface{}
		pass   bool
	}{
		{layout: "2006-01-02", value: "2012-03-02", pass: true},
		{layout: "15:04:05", value: "20:00:00", pass: true},
		{layout: "2 Jan 2006 15:04", value: "3 Nov 2020", pass: false},
		{layout: "2 Jan 2006 15:04:05", value: "3 Dec 2010 23:04:03", pass: true},
		{layout: "2006-01-02", value: "2019-30-78", pass: false},
	}
	for _, testcase := range testcases {
		validate, err := Date([]string{testcase.layout})
		if err != nil {
			t.Fatal(err)
		}
		item := &core.Item{Value: reflect.ValueOf(testcase.value)}
		err = validate(item)
		pass := err == nil
		if pass != testcase.pass {
			t.Errorf("expected validation pass for layout '%s' with value '%s' result to be '%v'", testcase.layout, testcase.value, testcase.pass)
		}
	}
}
