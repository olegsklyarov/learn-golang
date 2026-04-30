package main

import (
	"fmt"
	"unsafe"
)

type AnyInteger interface {
	int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64
}

func printRange[T AnyInteger](a, b T) {
	fmt.Printf("%6T%22d...%d\n", a, a, b)
}

func sizeBits[T AnyInteger]() uint {
	var v T
	return uint(unsafe.Sizeof(v)) * 8
}

func bounds[T AnyInteger]() (T, T) {
	var min, max T
	zero, one := T(0), T(1)
	allBitsOnes := ^zero
	isSigned := allBitsOnes < 0
	if isSigned {
		bits := sizeBits[T]()
		max = one<<(bits-1) - 1
		min = -max - 1
	} else {
		min = 0
		max = allBitsOnes
	}
	return min, max
}

func main() {
	printRange(bounds[int8]())
	printRange(bounds[uint8]())

	printRange(bounds[int16]())
	printRange(bounds[uint16]())

	printRange(bounds[int32]())
	printRange(bounds[uint32]())

	printRange(bounds[int64]())
	printRange(bounds[uint64]())
}
