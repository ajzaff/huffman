package huffman

import (
	"fmt"
	"iter"
	"math/bits"
)

// Elem directly computes the k^th element of the Huffman sequence of 64 elements.
//
// Elem panics if k or n are out of range.
func Elem(k int) uint64 {
	if k < 0 || k > 64 {
		panic(fmt.Errorf("k is out of range: %d", k))
	}
	return 1<<k - 1
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
		var (
			uv   uint64
			mask uint64 = 1
		)
		for range n {
			if !yield(uv) {
				return
			}
			uv |= mask
			mask <<= 1
		}
		yield(1<<n - 1)
	}
}

// Len returns the length in bits of the k^th huffman elem in uv.
func Len(uv uint64) int { return bits.TrailingZeros64(^uv) }

// Append appends the exact Huffman element onto b.
//
// For Append to function correctly, uv must be a valid Huffman code
// like those produced by Elem, Seq, or SeqN.
//
// Append panics if n is out of range.
func Append(b []byte, uv uint64) []byte {
	for i := 0; i < 64; i += 8 {
		x := byte(uv >> i)
		b = append(b, x)
		if x < 0xff {
			break
		}
	}
	return b
}
