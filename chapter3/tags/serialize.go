package tags

import "reflect"

// SerializeStructStrings converts a struct
// to our custom serialization format
// it honors serialize struct tags for string types
func SerializeStructStrings(s interface{}) (string, error) {
	result := ""

	// reflect the interface into
	// a type
	r := reflect.TypeOf(s)
	value := reflect.ValueOf(s)

	// if a pointer to a struct is passed
	// in, handle it appropriately
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
		value = value.Elem()
	}

	// loop over all of the fields
	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)
		// struct tag found
		key := field.Name
		if serialize, ok := field.Tag.Lookup("serialize"); ok {
			// ignore "-" otherwise that whole value
			// becomes the serialize 'key'
			if serialize == "-" {
				continue
			}
			key = serialize
		}

		switch value.Field(i).Kind() {
		// this recipe only supports strings!
		case reflect.String:
			result += key + ":" + value.Field(i).String() + ";"
			// by default skip it
		default:
			continue
		}
	}
	return result, nil
}
