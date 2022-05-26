package equal

import "testing"

func TestEqual(t *testing.T) {
	if !StringEqual("ajith", "ajith") {
		t.Errorf("expected true but got false")
	}
}

func FuzzStringEqual(f *testing.F) {
	f.Fuzz(func(t *testing.T, a, b string) {
		StringEqual(a, b)
	})
}
