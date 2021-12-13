package benchmark

import (
	"fmt"
	"testing"
)

type slice struct {
	I int32
	F float32
	A [32]int32 //128 136
}

func sliceOfPointers() []*slice {
	var s []*slice
	for i := 0; i < 1_0; i++ {
		s = append(s, &slice{I: 23, F: 3.445, A: [32]int32{2, 3, 0, 4, 4, 54, 5, 5, 5, 5, 6, 6, 6, 0, 0, 76, 7, 6, 7}})
	}
	return s
}

func pointerToSlice() *[]slice {
	var s []slice
	for i := 0; i < 1_0; i++ {
		s = append(s, slice{I: 23, F: 3.445, A: [32]int32{2, 3, 0, 4, 4, 54, 5, 5, 5, 5, 6, 6, 6, 0, 0, 76, 7, 6, 7}})
	}
	return &s

}

func Benchmark_pointerToSlice(b *testing.B) {
	b.ReportAllocs()
	var s *[]slice
	for i := 0; i < b.N; i++ {
		s = pointerToSlice()
	}

	fmt.Println(len(*s), cap(*s), "--", s)
}

func Benchmark_sliceOfPointers(b *testing.B) {
	b.ReportAllocs()
	var s []*slice
	for i := 0; i < b.N; i++ {
		s = sliceOfPointers()
	}
	fmt.Println(len(s), cap(s), "--", s[4])
}
