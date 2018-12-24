package hello

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	need := "Hello, Saed"
	res := SayHelloFor("Saed")
	if res != need {
		t.Errorf("Bad result, expected '%v', got '%v'", need, res)
	}
}

func BenchmarkSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci(50)
	}
}

//go test -bench . -benchmem -cpuprofile=cpu.out -memprofile=mem.out -memprofilerate=1
//go tool pprof cpu.out
//top
//list fibonacci
//web
//
//alloc_space
