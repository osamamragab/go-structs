package structs

import "testing"

func TestFieldNames(t *testing.T) {
	s1 := struct {
		unexported string
		Name       string `tag:"name"`
		IsValid    bool   `tag:"is_valid"`
		Score      int
	}{
		"invisible",
		"myname",
		true,
		20,
	}

	want := []string{"name", "is_valid", "Score"}
	got := FieldNames(s1, "tag")

	if len(want) != len(got) {
		t.Errorf("FieldNames: incorrect slice length. want %d, got %d", len(want), len(got))
	}

	for i, v := range want {
		if gv := got[i]; gv != v {
			t.Errorf("FieldNames: (%d) want %v, got %v", i, v, gv)
		}
	}
}
