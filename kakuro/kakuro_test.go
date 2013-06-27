package kakuro

import (
	"os"
	"testing"
)

func benchsolve(name string, b *testing.B) {
	p := os.Getenv("GOPATH") + "/src/github.com/sbalev/gobananas/kakuro/data/" +
		name + ".txt"
	f, err := os.Open(p)
	if err != nil {
		b.Fatal(err)
	}
	g, err := ReadGrid(f)
	f.Close()
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := g.Solve(); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSolve5x5(b *testing.B) {
	benchsolve("5x5", b)
}

func BenchmarkSolve7x9(b *testing.B) {
	benchsolve("7x9", b)
}

func BenchmarkSolve9x11(b *testing.B) {
	benchsolve("9x11", b)
}

func BenchmarkSolve10x10(b *testing.B) {
	benchsolve("10x10", b)
}

func BenchmarkSolve15x15(b *testing.B) {
	benchsolve("15x15", b)
}

func BenchmarkSolve29x29(b *testing.B) {
	benchsolve("29x29", b)
}
