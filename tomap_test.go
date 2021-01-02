package structs

import "testing"

func TestToMap(t *testing.T) {
	s1 := struct {
		unexported string
		Name       string `tag:"name"`
		IsValid    bool   `tag:"is_valid"`
	}{
		"invisible",
		"myname",
		true,
	}

	want := map[string]interface{}{
		"name":     "myname",
		"is_valid": true,
	}
	got := ToMap(s1, "tag")

	if len(want) != len(got) {
		t.Errorf("ToMap: incorrect map length. want %d, got %d", len(want), len(got))
	}

	for k, v := range want {
		if gv, exists := got[k]; !exists || gv != v {
			t.Errorf("ToMap: (%s) want %v, got %v", k, v, gv)
		}
	}
}
