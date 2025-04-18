package huffman

import (
	"encoding/binary"
	"strconv"
	"testing"
)

func TestElemZero(t *testing.T) {
	defer func() {
		if recover() != nil {
			t.Fail()
		}
	}()
	t.Log(Elem(0))
}

func TestElem64(t *testing.T) {
	defer func() {
		if recover() != nil {
			t.Fail()
		}
	}()
	t.Log(Elem(64))
}

func TestElem4Seq(t *testing.T) {
	for i := range 4 {
		x := Elem(i)
		t.Log(i, x, strconv.FormatUint(x, 2))
	}
}

func TestElemSeq(t *testing.T) {
	for i := range 65 {
		x := Elem(i)
		t.Log(i, x, strconv.FormatUint(x, 2))
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

func TestLen(t *testing.T) {
	var i int
	for x := range Seq() {
		t.Log(i, Len(x), x, strconv.FormatUint(x, 2))
		i++
	}
}

func TestAppendSeqN_9(t *testing.T) {
	var i int
	for x := range SeqN(9) {
		t.Logf("%d %s %x %x %d %d", i, strconv.FormatUint(x, 2), Append(nil, x), binary.AppendUvarint(nil, x), len(Append(nil, x)), len(binary.AppendUvarint(nil, x)))
		i++
	}
}

func TestAppendSeq(t *testing.T) {
	var i int
	for x := range Seq() {
		t.Logf("%d %s %x %x %d %d", i, strconv.FormatUint(x, 2), Append(nil, x), binary.AppendUvarint(nil, x), len(Append(nil, x)), len(binary.AppendUvarint(nil, x)))
		i++
	}
}
