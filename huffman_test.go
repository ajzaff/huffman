package huffman

import (
	"strconv"
	"testing"
)

func TestElemZero_65(t *testing.T) {
	defer func() {
		if recover() != nil {
			t.Fail()
		}
	}()
	t.Log(Elem(0, 65))
}

func TestElem64_65(t *testing.T) {
	defer func() {
		if recover() != nil {
			t.Fail()
		}
	}()
	t.Log(Elem(64, 65))
}

func TestElem4Seq(t *testing.T) {
	for i := range 4 {
		x := Elem(i, 4)
		t.Log(i, strconv.FormatUint(x, 2))
	}
}

func TestElem65Seq(t *testing.T) {
	for i := range 65 {
		x := Elem(i, 65)
		t.Log(i, strconv.FormatUint(x, 2))
	}
}

func TestSeqNZero(t *testing.T) {
	defer func() {
		if recover() != nil {
			t.Fail()
		}
	}()
	for x := range SeqN(0) {
		t.Log(x)
	}
}

func TestSeq4(t *testing.T) {
	var i int
	for x := range SeqN(4) {
		t.Log(i, x, strconv.FormatUint(x, 2))
		i++
	}
}

func TestSeq(t *testing.T) {
	var i int
	for x := range Seq() {
		t.Log(i, x, strconv.FormatUint(x, 2))
		i++
	}
}
