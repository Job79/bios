package form

import (
	"net/url"
	"testing"
)

type test struct {
	String string   `form:"string"`
	Int    int      `form:"Int"`
	Array  []string `form:"Array"`
}

// TestUnmarshal tests whether the unmarshal function works correctly
func TestUnmarshal(t *testing.T) {
	values := url.Values{
		"string": []string{"string"},
		"Int":    []string{"1"},
		"Array":  []string{"a", "b"},
	}

	var test test
	Unmarshal(values, &test)

	if test.String != "string" {
		t.Errorf("string = %s, want string", test.String)
	} else if test.Int != 1 {
		t.Errorf("Int = %d, want 1", test.Int)
	} else if len(test.Array) != 2 {
		t.Errorf("Array = %v, want [a b]", test.Array)
	}
}
