package main

import (
	"fmt"
	"testing"
)

var (
	chaC = []int{
		12, 3433, 23, 23, 23, 23, 23, 32, 324, 435, 5456, 65, 67, 78, 87, 899, 98, 90, 90, 00, 98, 76, 65, 54, 45, 3432, 23, 43, 54, 54, 56, 54, 34, 4343, 34,
		32, 2323, 32, 32, 23, 23, 344, 54, 54, 65, 657, 67, 8, 76, 6, 554, 43, 4, 23, 6576, 543, 45, 34, 32, 4, 556, 65, 45, 43, 23, 34, 4, 545, 56, 567, 4, 45, 34, 42, 342,
		12, 3433, 23, 23, 23, 23, 23, 32, 324, 435, 5456, 65, 67, 78, 87, 899, 98, 90, 90, 00, 98, 76, 65, 54, 45, 3432, 23, 43, 54, 54, 56, 54, 34, 4343, 34,
		32, 2323, 32, 32, 23, 23, 344, 54, 54, 65, 657, 67, 8, 76, 6, 554, 43, 33235434, 23, 6576, 543, 45, 34, 32, 4, 556, 65, 45, 43, 23, 34, 4, 545, 56, 567, 56543, 45, 34, 42, 342,
		12, 3433, 23, 23, 23, 23, 23, 32, 324, 435, 5456, 65, 67, 78, 87, 899, 98, 90, 90, 00, 98, 76, 65, 54, 45, 3432, 23, 43, 54, 54, 56, 54, 34, 4343, 34,
		32, 2323, 32, 32, 23, 23, 344, 54, 54, 65, 657, 67, 8, 76, 6, 554, 43, 876, 23, 6576, 543, 45, 34, 32, 4, 556, 65, 45, 43, 23, 34, 4, 545, 56, 567, 3443, 45, 34, 42, 342, 12, 3433, 23, 23, 23, 23, 23, 32, 324, 435, 5456, 65, 67, 78, 87, 899, 98, 90, 90, 00, 98, 76, 65, 54, 45, 3432, 23, 43, 54, 54, 56, 54, 34, 4343, 34,
		32, 2323, 32, 32, 23, 23, 344, 54, 54, 65, 657, 67, 8, 76, 6, 554, 43, 65, 23, 6576, 543, 45, 34, 32, 4, 556, 65, 45, 43, 23, 34, 4, 545, 56, 567, 3565, 45, 34, 42, 342,
		12, 3433, 23, 23, 23, 23, 23, 32, 324, 435, 5456, 65, 67, 78, 87, 899, 98, 90, 90, 00, 98, 76, 65, 54, 45, 3432, 23, 43, 54, 54, 56, 54, 34, 4343, 34,
		32, 2323, 32, 32, 23, 23, 344, 54, 54, 65, 657, 67, 8, 76, 6, 554, 43, 6534, 23, 6576, 543, 45, 34, 32, 4, 556, 65, 45, 43, 23, 34, 4, 545, 56, 567, 645, 45, 34, 42, 342,
		12, 3433, 23, 23, 23, 23, 23, 32, 324, 435, 5456, 65, 67, 78, 87, 899, 98, 90, 90, 00, 98, 76, 65, 54, 45, 3432, 23, 43, 54, 54, 56, 54, 34, 4343, 34,
		32, 2323, 32, 32, 23, 23, 344, 54, 54, 65, 657, 67, 8, 76, 6, 554, 43, 64, 23, 6576, 543, 45, 34, 32, 4, 556, 65, 45, 43, 23, 34, 4, 545, 56, 567, 543, 45, 34, 42, 342,
		12, 3433, 23, 23, 23, 23, 23, 32, 324, 435, 5456, 65, 67, 78, 87, 899, 98, 90, 90, 00, 98, 76, 65, 54, 45, 3432, 23, 43, 54, 54, 56, 54, 34, 4343, 34,
		32, 2323, 32, 32, 23, 23, 344, 54, 54, 65, 657, 67, 8, 76, 6, 554, 43, 65, 23, 6576, 543, 45, 34, 32, 4, 556, 65, 45, 43, 23, 34, 4, 545, 56, 567, 7654, 45, 34, 42, 342,
		12, 3433, 23, 23, 23, 23, 23, 32, 324, 435, 5456, 65, 67, 78, 87, 899, 98, 90, 90, 00, 98, 76, 65, 54, 45, 3432, 23, 43, 54, 54, 56, 54, 34, 4343, 34,
		32, 2323, 32, 32, 23, 23, 344, 54, 54, 65, 657, 67, 8, 76, 6, 554, 43, 23, 23, 6576, 543, 45, 34, 32, 4, 556, 65, 45, 43, 23, 34, 4, 545, 56, 567, 43, 45, 34, 42, 342,
		12, 3433, 23, 23, 23, 23, 23, 32, 324, 435, 5456, 65, 67, 78, 87, 899, 98, 90, 90, 00, 98, 76, 65, 54, 45, 3432, 23, 43, 54, 54, 56, 54, 34, 4343, 34,
		32, 2323, 32, 32, 23, 23, 344, 54, 54, 65, 657, 67, 8, 76, 6, 554, 43, 54, 23, 6576, 543, 45, 34, 32, 4, 556, 65, 45, 43, 23, 34, 4, 545, 56, 567, 4334, 45, 34, 42, 342,
		12, 3433, 23, 23, 23, 23, 23, 32, 324, 435, 5456, 65, 67, 78, 87, 899, 98, 90, 90, 00, 98, 76, 65, 54, 45, 3432, 23, 43, 54, 54, 56, 54, 34, 4343, 34,
		32, 2323, 32, 32, 23, 23, 344, 54, 54, 65, 657, 67, 8, 76, 6, 554, 43, 9876, 23, 6576, 543, 45, 34, 32, 4, 556, 65, 45, 43, 23, 34, 4, 545, 56, 567, 334, 45, 34, 42, 342,
	}
)

func TestXiEr(t *testing.T) {
	fmt.Println(XiEr(chaC...))
}

//33186 ns/op
func BenchmarkXiEr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		XiEr(chaC...)
	}
}

func BenchmarkCha2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cha(chaC...)
	}
}

func BenchmarkShellSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ShellSort(chaC)
	}
}

func TestMergeSort(t *testing.T) {
	fmt.Println(MergeSort(chaC))
}