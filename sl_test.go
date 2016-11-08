package xxx_test

import "testing"

const sz = 200000000

func BenchmarkInd(b *testing.B) {
	for bc := 0; bc < b.N; bc++ {
		x := make([]int, sz)
		for i := 0; i < sz; i++ {
			x[i] = i
		}
		b.SetBytes(sz)
	}
}

func BenchmarkRInd(b *testing.B) {
	for bc := 0; bc < b.N; bc++ {
		x := make([]int, sz)
		for i := range x {
			x[i] = i
		}
		b.SetBytes(sz)
	}
}

func BenchmarkApp(b *testing.B) {
	for bc := 0; bc < b.N; bc++ {
		x := make([]int, 0, sz)
		for i := 0; i < sz; i++ {
			x = append(x, i)
		}
		b.SetBytes(sz)
	}
}

func BenchmarkManual(b *testing.B) {
	for bc := 0; bc < b.N; bc++ {
		x := make([]int, 0, sz)
		for i := 0; i < sz; i++ {
			x = x[:i+1]
			x[i] = i
		}
		b.SetBytes(sz)
	}
}

func BenchmarkArr(b *testing.B) {
	for bc := 0; bc < b.N; bc++ {
		x := [sz]int{}
		for i := 0; i < sz; i++ {
			x[i] = i
		}
		b.SetBytes(sz)
	}
}

func BenchmarkRArr(b *testing.B) {
	for bc := 0; bc < b.N; bc++ {
		x := [sz]int{}
		for i := range x {
			x[i] = i
		}
		b.SetBytes(sz)
	}
}

var back = [sz]int{}

func BenchmarkBInd(b *testing.B) {
	for bc := 0; bc < b.N; bc++ {
		x := back[:]
		for i := 0; i < sz; i++ {
			x[i] = i
		}
		b.SetBytes(sz)
	}
}

func BenchmarkBRInd(b *testing.B) {
	for bc := 0; bc < b.N; bc++ {
		x := back[:]
		for i := range x {
			x[i] = i
		}
		b.SetBytes(sz)
	}
}

func BenchmarkBApp(b *testing.B) {
	for bc := 0; bc < b.N; bc++ {
		x := back[:0]
		for i := 0; i < sz; i++ {
			x = append(x, i)
		}
		b.SetBytes(sz)
	}
}

func BenchmarkBManual(b *testing.B) {
	for bc := 0; bc < b.N; bc++ {
		x := back[:0]
		for i := 0; i < sz; i++ {
			x = x[:i+1]
			x[i] = i
		}
		b.SetBytes(sz)
	}
}
