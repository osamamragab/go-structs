package structs

import "reflect"

// FieldNames gets field names of tag in struct.
// if tag is empty string (""), field name is used.
// Panics if s is not a struct.
func FieldNames(s interface{}, tag string) []string {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		panic("structs.FieldNames: given value must be a struct")
	}

	t := v.Type()
	n := v.NumField()

	names := make([]string, 0, n)
	for i := 0; i < n; i++ {
		f := t.Field(i)
		// skip unexported fields
		if f.PkgPath != "" {
			continue
		}

		name := f.Name
		if tag != "" {
			tn, _ := parseTag(t.Field(i).Tag.Get(tag))
			if tn == "-" {
				continue
			}
			if tn != "" {
				name = tn
			}
		}
		names = append(names, name)
	}

	return names
}
