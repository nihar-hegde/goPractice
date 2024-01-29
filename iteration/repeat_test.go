package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeate("a", 4)
	expected := "aaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeate("a", 4)
	}
}
