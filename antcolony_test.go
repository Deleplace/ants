package ants

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

var Sink int

func BenchmarkAntSearchByField1(b *testing.B) {
	r := rand.New(rand.NewSource(42))
	ants := generateColony(r, 1_000_000)
	needle := 42
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, ant := range ants {
			if ant.Field1 == needle {
				Sink = i
			}
		}
	}
}

func BenchmarkDOAntSearchByField1(b *testing.B) {
	r := rand.New(rand.NewSource(42))
	ants := generateDataOrientedColony(r, 1_000_000)
	needle := 42
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i, f1 := range ants.Field1 {
			if f1 == needle {
				Sink = i
			}
		}
	}
}

func BenchmarkAntSearchByField2(b *testing.B) {
	r := rand.New(rand.NewSource(42))
	ants := generateColony(r, 1_000_000)
	needle := "field2-42"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, ant := range ants {
			if ant.Field2 == needle {
				Sink = i
			}
		}
	}
}

func BenchmarkDOAntSearchByField2(b *testing.B) {
	r := rand.New(rand.NewSource(42))
	ants := generateDataOrientedColony(r, 1_000_000)
	needle := "field2-42"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i, f2 := range ants.Field2 {
			if f2 == needle {
				Sink = i
			}
		}
	}
}

func BenchmarkAntInspect(b *testing.B) {
	r := rand.New(rand.NewSource(42))
	ants := generateColony(r, 1_000_000)
	b.ResetTimer()
	Sink = 0
	for i := 0; i < b.N; i++ {
		for _, ant := range ants {
			if ant.Field1 < ant.Field3 {
				Sink++
			}
			if ant.Field5 < ant.Field7 {
				Sink++
			}
		}
	}
	// b.Logf("BenchmarkAntInspect\t%d => %d", b.N, Sink)
}

func BenchmarkDOAntInspect(b *testing.B) {
	r := rand.New(rand.NewSource(42))
	const size = 1_000_000
	ants := generateDataOrientedColony(r, size)
	b.ResetTimer()
	Sink = 0
	for i := 0; i < b.N; i++ {
		for j := 0; j < size; j++ {
			if ants.Field1[j] < ants.Field3[j] {
				Sink++
			}
			if ants.Field5[j] < ants.Field7[j] {
				Sink++
			}
		}
	}
	// b.Logf("BenchmarkDOAntInspect\t%d => %d", b.N, Sink)
}

func generateColony(r *rand.Rand, size int) AntColony {
	ants := make(AntColony, size)
	for i := range ants {
		ants[i] = Ant{
			Field1: r.Intn(1000),
			Field2: fmt.Sprintf("field2-%d", r.Intn(1000)),
			Field3: r.Intn(1000),
			Field4: fmt.Sprintf("field4-%d", r.Intn(1000)),
			Field5: r.Intn(1000),
			Field6: fmt.Sprintf("field6-%d", r.Intn(1000)),
			Field7: r.Intn(1000),
			Field8: fmt.Sprintf("fiel8-%d", r.Intn(1000)),
		}
	}
	return ants
}

func generateDataOrientedColony(r *rand.Rand, size int) DataOrientedAntColony {
	ants := DataOrientedAntColony{
		Field1: make([]int, size),
		Field2: make([]string, size),
		Field3: make([]int, size),
		Field4: make([]string, size),
		Field5: make([]int, size),
		Field6: make([]string, size),
		Field7: make([]int, size),
		Field8: make([]string, size),
	}
	for i := 0; i < size; i++ {
		ants.Field1[i] = r.Intn(1000)
		ants.Field2[i] = fmt.Sprintf("field2-%d", r.Intn(1000))
		ants.Field3[i] = r.Intn(1000)
		ants.Field4[i] = fmt.Sprintf("field4-%d", r.Intn(1000))
		ants.Field5[i] = r.Intn(1000)
		ants.Field6[i] = fmt.Sprintf("field6-%d", r.Intn(1000))
		ants.Field7[i] = r.Intn(1000)
		ants.Field8[i] = fmt.Sprintf("fiel8-%d", r.Intn(1000))
	}
	return ants
}

func TestAntStructSize(t *testing.T) {
	// For a 64-bit platform:
	// 4*8 + 4*16 == 96 bytes
	var i int
	var s string
	intSize := reflect.TypeOf(i).Size()
	stringRefSize := reflect.TypeOf(s).Size()
	expected := 4*intSize + 4*stringRefSize
	var ant Ant
	actual := reflect.TypeOf(ant).Size()
	if actual != expected {
		t.Errorf("Expected Ant struct of 4*%d + 4*%d == %d bytes, observed %d", intSize, stringRefSize, expected, actual)
	}
}
