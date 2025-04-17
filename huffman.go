package huffman

import (
	"encoding/binary"
	"fmt"
	"iter"
	"math/bits"
	"slices"
)

// Elem directly computes the k^th element of the Huffman sequence of n elements.
//
// Elem panics if k or n are out of range.
func Elem(k, n int) uint64 {
	k1 := k + 1
	if k < 0 || n < 0 || k1 >= n || n > 65 {
		if k1 == n {
			return 1<<k - 1
		}
		panic(fmt.Errorf("n is out of range: %d", n))
	}
	return 1<<k1 - 2
}

// Seq returns the full Huffman sequence of length 65 elements.
func Seq() iter.Seq[uint64] { return SeqN(65) }

// SeqN returns the Huffman sequence of n elements.
//
// SeqN panics if n is out of range.
func SeqN(n int) iter.Seq[uint64] {
	if n < 0 || n > 65 {
		panic(fmt.Errorf("n is out of range: %d", n))
	}
	return func(yield func(uint64) bool) {
		if n == 0 {
			return
		}
		n--
		var uv uint64
		for range n {
			if !yield(uv) {
				return
			}
			uv |= 1
			uv <<= 1
		}
		yield(1<<n - 1)
	}
}

// Len returns the length in bits of the k^th huffman elem of n elements.
//
// For Len to function correctly, uv must be a valid Huffman code
// like those produced by Elem, Seq, or SeqN.
//
// Len panics if n is out of range.
func Len(uv uint64, n int) int {
	if n < 0 || n > 64 {
		panic(fmt.Errorf("n is out of range: %d", n))
	}
	sz := bits.OnesCount64(uv)
	if sz < n {
		sz++
	}
	return sz
}

// AppendElem appends the exact Huffman element onto b.
//
// For AppendElem to function correctly, uv must be a valid Huffman code
// like those produced by Elem, Seq, or SeqN.
//
// AppendElem panics if n is out of range.
func AppendElem(b []byte, uv uint64, n int) []byte {
	sz := Len(uv, n)
	b = slices.Grow(b, sz)
	return binary.BigEndian.AppendUint64(b, uv)
}
