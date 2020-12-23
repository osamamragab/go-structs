package structs

// ToMap converts a struct to a map.
// tag is defined in struct's tags, used as map key if not empty
// otherwise field name is used.
// Panics if s is not a struct.
func ToMap(s interface{}, tag string) map[string]interface{} {
	m := make(map[string]interface{})

	v := structValue(s)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		tf := t.Field(i)

		key, opts := tf.Name, tagOptions("")
		if tag != "" {
			tn, to := parseTag(tf.Tag.Get(tag))
			if tn == "-" {
				continue
			}
			if tn != "" {
				key = tn
			}
			opts = to
		}

		vf := v.Field(i)
		if opts.Contains("omitempty") && isEmptyValue(vf) {
			continue
		}

		m[key] = vf.Interface()
	}

	return m
}
