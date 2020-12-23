package structs

// ToMap converts a struct to a map.
// tag is defined in struct's tags, used as map key.
// Panics if s is not a struct.
func ToMap(s interface{}, tag string) (map[string]interface{}, error) {
	m := map[string]interface{}{}

	v := structValue(s)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		tf := t.Field(i)

		tn, to := parseTag(tf.Tag.Get(tag))
		if tn == "" || tn == "-" {
			continue
		}

		vf := v.Field(i)
		if to.Contains("omitempty") && isEmptyValue(vf) {
			continue
		}

		m[tn] = vf.Interface()
	}

	return m, nil
}
