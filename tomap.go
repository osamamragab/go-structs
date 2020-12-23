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

		tn, to := tf.Name, tagOptions("")
		if tag != "" {
			n, o := parseTag(tf.Tag.Get(tag))
			if n == "-" {
				continue
			}
			if n != "" {
				tn = n
			}
			to = o
		}

		vf := v.Field(i)
		if to.Contains("omitempty") && isEmptyValue(vf) {
			continue
		}

		m[tn] = vf.Interface()
	}

	return m
}
