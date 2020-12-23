package structs

// FieldNames gets field names of tag in struct.
// if tag is empty string (""), field name is used.
// Panics if s is not a struct.
func FieldNames(s interface{}, tag string) []string {
	v := structValue(s)
	t := v.Type()
	n := v.NumField()

	names := make([]string, 0, n)
	for i := 0; i < n; i++ {
		name := t.Field(i).Name
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
