package form

/**
 * Small helper that converts url.Values to a struct
 */

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
)

// Unmarshal parses the given url.Values into a struct
func Unmarshal(values url.Values, d any) {
	// Get the type and values from the struct
	dType := reflect.TypeOf(d).Elem()
	dVal := reflect.ValueOf(d)

	// Check whether d really is a struct
	if dType.Kind() != reflect.Struct {
		panic(fmt.Errorf("unmarshal: given interface must be a struct"))
	}

	for i := 0; i < dType.NumField(); i++ {
		// Get the field and check whether it has the form tag
		field := dType.Field(i)
		key, ok := field.Tag.Lookup("form")
		if !ok {
			continue
		}

		// Get the field value and check if we can interface with it
		// This is false when the field or struct is private
		result := dVal.Elem().Field(i)
		if result.CanInterface() {

			// Determine the type and set the value
			switch result.Interface().(type) {
			case string:
				result.SetString(values.Get(key))
			case int:
				if i, err := strconv.ParseInt(values.Get(key), 10, 64); err == nil {
					result.SetInt(i)
				}
			case []string:
				result.Set(reflect.ValueOf(values[key]))
			default:
				// Because this is an error that doesn't depend on user input and
				// occurs every time it is run on the given struct, we just panic.
				panic(fmt.Errorf("unmarshal: type of field '%s' is not supported", key))
			}
		}
	}
}
