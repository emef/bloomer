package bloomer

import (
	"hash"
	"math"
	"crypto/sha1"
	"encoding/binary"
	"github.com/emef/bitfield"
)

type Bloomer struct {
	field bitfield.BitField
	size int
	k int
	sha hash.Hash
}

func New(size, k int) *Bloomer {
	return &Bloomer{
		bitfield.New(size),
		size,
		k,
		sha1.New()}
}

func NewSuggested(n int, p float64) *Bloomer {
	m := -(float64(n) * math.Log(p)) / math.Pow(math.Log(2), 2)
	k := (m / float64(n)) * math.Log(2)
	return New(int(m), int(math.Ceil(k)))
}

func (b Bloomer) Add(value []byte) {
	keys := b.getHashKeys(value)
	for _, key := range keys {
		b.field.Set(key)
	}
}

func (b Bloomer) Test(value []byte) bool {
	keys := b.getHashKeys(value)
	for _, key := range keys {
		if !b.field.Test(key) {
			return false
		}
	}
	return true
}

func (b Bloomer) TestAndSet(value []byte) bool {
	found := true
	keys := b.getHashKeys(value)
	for _, key := range keys {
		if !b.field.Test(key) {
			found = false
			b.Set(key)
		}
	}
	return found
}

func (b Bloomer) getHashKeys(value []byte) []uint32 {
	keys := make([]uint32, b.k)
	hashBytes := b.sha.Sum(value)
	hash0:= binary.BigEndian.Uint64(hashBytes[:8])
	hash1 := binary.BigEndian.Uint64(hashBytes[8:16])
	for i := 0; i < b.k; i++ {
		keys[i] = uint32((hash0 + uint64(i) * hash1) % uint64(b.size))
	}
	return keys
}