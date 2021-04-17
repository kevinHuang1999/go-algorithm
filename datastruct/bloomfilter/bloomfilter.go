package bloomfilter

import (
	"github.com/willf/bitset"
)

const DEFAULT_SIZE = 2 << 24

var seeds = []uint{7, 11, 13, 31, 37, 61}

type SimpleHash struct {
	cap  uint
	seed uint
}

type BloomFilter struct {
	bs    *bitset.BitSet
	funcs [6]SimpleHash
}

func NewBloomFilter() *BloomFilter {
	bf := new(BloomFilter)
	for i := 0; i < len(bf.funcs); i++ {
		bf.funcs[i] = SimpleHash{DEFAULT_SIZE, seeds[i]}
	}
	bf.bs = bitset.New(DEFAULT_SIZE)
	return bf
}

func (s SimpleHash) hash(value string) uint {
	var result uint = 0
	for i := 0; i < len(value); i++ {
		result = result*s.seed + uint(value[i])
	}
	return (s.cap - 1) & result
}

func (bf *BloomFilter) add(value string) {
	for _, f := range bf.funcs {
		bf.bs.Set(f.hash(value))
	}
}

func (bf *BloomFilter) contains(value string) bool {
	if value == "" {
		return false
	}
	ret := true
	for _, f := range bf.funcs {
		ret = ret && bf.bs.Test(f.hash(value))
	}
	return ret
}
