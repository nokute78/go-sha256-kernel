package sha256_test

import (
	"bytes"
	sha256_orig "crypto/sha256"
	"encoding/hex"
	"github.com/nokute78/go-sha256-kernel"
	"testing"
)

func TestSum256(t *testing.T) {
	type testcase struct {
		input  []byte
		expect [sha256.Size]byte
	}

	cases := []testcase{
		{input: []byte("hoge")},
	}

	for i, v := range cases {
		cases[i].expect = sha256_orig.Sum256(v.input)

		ret := sha256.Sum256(v.input)
		if 0 != bytes.Compare(ret[:], cases[i].expect[:]) {
			t.Errorf("mismatch!\ngot:   %s\nexpect:%s",
				hex.EncodeToString(ret[:]),
				hex.EncodeToString(cases[i].expect[:]))
			break
		}
	}
}

func TestSumMultiSum(t *testing.T) {
	h := sha256.New()
	h_orig := sha256_orig.New()

	datas := [][]byte{
		[]byte("hoge"),
		[]byte("moge"),
		[]byte("age"),
	}

	for i, v := range datas {
		_, err := h.Write(v)
		if err != nil {
			t.Errorf("%d: sha256.Write err:%s", i, err)
			break
		}
		_, err = h_orig.Write(v)
		if err != nil {
			t.Errorf("%d: sha256_orig.Write err:%s", i, err)
			break
		}

		b := h.Sum([]byte{})
		b_orig := h_orig.Sum([]byte{})

		if 0 != bytes.Compare(b, b_orig) {
			t.Errorf("%d:mismatch!\ngot:   %s\nexpect:%s", i,
				hex.EncodeToString(b),
				hex.EncodeToString(b_orig))
		}
	}

	h.Reset()
	h_orig.Reset()

	for i, v := range datas {
		_, err := h.Write(v)
		if err != nil {
			t.Errorf("%d: sha256.Write err:%s", i, err)
			break
		}
		_, err = h_orig.Write(v)
		if err != nil {
			t.Errorf("%d: sha256_orig.Write err:%s", i, err)
			break
		}

		b := h.Sum([]byte{})
		b_orig := h_orig.Sum([]byte{})

		if 0 != bytes.Compare(b, b_orig) {
			t.Errorf("%d:mismatch!\ngot:   %s\nexpect:%s", i,
				hex.EncodeToString(b),
				hex.EncodeToString(b_orig))
		}
	}
}

func createInputs() [][]byte {
	ret := [][]byte{
		[]byte("hoge"),
		[]byte("moge"),
	}
	return ret
}

func BenchmarkGoSum256(b *testing.B) {
	inputs := createInputs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range inputs {
			sha256_orig.Sum256(v)
		}
	}
}

func BenchmarkKernelSum256(b *testing.B) {
	inputs := createInputs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, v := range inputs {
			sha256.Sum256(v)
		}
	}
}
