package eulergo

import (
	"fmt"
	"testing"
)

func euler1TestError(t *testing.T, input uint64, result uint64, expected uint64) {
	if result != expected {
		ErrorString := fmt.Sprintf("Euler1(%d) == %d != %d", input, result, expected)
		t.Error(ErrorString)
	}
}

func TestEuler1(t *testing.T) {
	var result uint64 = 0
	result = Euler1(1000)
	euler1TestError(t, 1000, result, 233168)
	result = Euler1(0)
	euler1TestError(t, 0, result, 0)
	return
}

func BenchmarkEuler1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Euler1(100000)
	}
}
