package color

import "testing"

func TestC(t *testing.T) {
	expected := baseColors["blue"] + "Hello World!" + end
	if out := C("blue", "Hello ", "World!"); out != expected {
		t.Errorf("expencted %s, got %s", expected, out)
	}
}

func BenchmarkC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		C("blue", "Hello ", "World!")
	}
}
