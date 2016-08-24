package scratch

import (
	"testing"
	"runtime"
	"fmt"
)

const arrlen = 1 << 10
var TarrGlobal [arrlen]int
var TarrSliceGlobal = TarrGlobal[:]

func init() {
	fmt.Printf("==== arrlen: %v ====\n", int(arrlen))
}	

type TArr struct {
	arr [arrlen]int
	arrslice []int
}

func TestArrayVsSliceNoop(t *testing.T) { }

func BenchmarkArrayGlobal(b *testing.B) {
	runtime.GC()
	b.ResetTimer()	
	for i := 0; i < b.N; i++ {
		for j := 0; j < arrlen; j++ {
			TarrGlobal[j] = j
			// _ = arr[j] == -1
		}
	}
}

func BenchmarkArrayGlobalPtr(b *testing.B) {
	arrp := &TarrGlobal
	runtime.GC()
	b.ResetTimer()	
	for i := 0; i < b.N; i++ {
		for j := 0; j < arrlen; j++ {
			arrp[j] = j
			// _ = arr[j] == -1
		}
	}
}

func BenchmarkSliceGlobal(b *testing.B) {
	runtime.GC()
	b.ResetTimer()	
	for i := 0; i < b.N; i++ {
		for j := 0; j < arrlen; j++ {
			TarrSliceGlobal[j] = j
			// _ = arrslice[j] == -1
		}
	}
}

func BenchmarkArrayLocal(b *testing.B) {
	var arr [arrlen]int
	runtime.GC()
	b.ResetTimer()	
	for i := 0; i < b.N; i++ {
		for j := 0; j < arrlen; j++ {
			arr[j] = j
			// _ = arr[j] == -1
		}
	}
}

func BenchmarkSliceLocal(b *testing.B) {
	var arr [arrlen]int
	arrslice := arr[:]
	runtime.GC()
	b.ResetTimer()	
	for i := 0; i < b.N; i++ {
		for j := 0; j < arrlen; j++ {
			arrslice[j] = j
			// _ = arrslice[j] == -1
		}
	}
}

func BenchmarkArrayStructField(b *testing.B) {
	x := TArr{}
	x.arrslice = x.arr[:]
	runtime.GC()
	b.ResetTimer()	
	for i := 0; i < b.N; i++ {
		for j := 0; j < arrlen; j++ {
			x.arr[j] = j
			// _ = arr[j] == -1
		}
	}
}

func BenchmarkSliceStructField(b *testing.B) {
	x := TArr{}
	x.arrslice = x.arr[:]
	runtime.GC()
	b.ResetTimer()	
	for i := 0; i < b.N; i++ {
		for j := 0; j < arrlen; j++ {
			x.arrslice[j] = j
			// _ = arrslice[j] == -1
		}
	}
}
