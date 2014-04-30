package bloomer

import (
	"math"
	"testing"
	"crypto/rand"
)

func TestBasic(t *testing.T) {
	N := int(1e4)
	E := 1e-3
	b := NewSuggested(N, E)

	// generate and add random keys
	keys := make([][]byte, N)
	for i := range keys {
		keys[i] = make([]byte, 10)
		rand.Read(keys[i])
		b.Add(keys[i])

		if !b.Get(keys[i]) {
			t.Error("uh oh, bloom filter is busted")
			t.FailNow()
		}
	}

	// generate and test missing random keys
	key := make([]byte, 9)
	falsePos := 0.0
	for i := 0; i < int(N); i++ {
		rand.Read(key)
		if b.Get(key) {
			falsePos++
		}
	}

	// memory usage and false positive rate
	sizeBytes := float64(len(b.field)) / math.Pow(2, 20)
	falseRate := falsePos / float64(N)

	if sizeBytes > 2 {
		t.Errorf("Used too much memory: %.2fMB\n", sizeBytes)
	}

	// this is probabilistic... so give it some room
	if falseRate > (E * 1.5) {
		t.Errorf("False positive rate too high: %.5f\n", falseRate)
	}
}

func BenchmarkAdd(b *testing.B) {
	f := NewSuggested(b.N, 1e-3)
	keys := make([][]byte, b.N)
	for i := range keys {
		keys[i] = make([]byte, 10)
		rand.Read(keys[i])
	}

	b.ResetTimer()
	for _, key := range keys {
		f.Add(key)
	}
}