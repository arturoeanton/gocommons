package bloom

import (
	"fmt"

	"github.com/arturoeanton/gocommons/hash"
	"github.com/arturoeanton/gocommons/memset"
)

type Bloom struct {
	memset *memset.MemSet
	repeat map[uint64]uint64
	fxHash func(string) uint64
}

func NewBloom() *Bloom {
	return &Bloom{memset: memset.NewMemSet(), fxHash: hash.HashStringUint64, repeat: make(map[uint64]uint64)}
}

func (b *Bloom) Add(param interface{}) {
	//fmt.Println("Add", param)
	value := fmt.Sprint(param)
	hashValue := b.fxHash(value)
	if b.memset.Contains(hashValue) {
		if _, ok := b.repeat[hashValue]; !ok {
			b.repeat[hashValue] = 1
			return
		}
		b.repeat[hashValue]++
		return
	}
	b.memset.Add(hashValue)
}

func (b *Bloom) Remove(param interface{}) {
	//fmt.Println("Remove", param)
	value := fmt.Sprint(param)
	hashValue := b.fxHash(value)
	if b.memset.Contains(hashValue) {
		if _, ok := b.repeat[hashValue]; ok {
			b.repeat[hashValue]--
			if b.repeat[hashValue] <= 0 {
				delete(b.repeat, hashValue)
			}
			return
		}
	}
	b.memset.Remove(hashValue)
}

func (b *Bloom) Contains(param interface{}) bool {
	value := fmt.Sprint(param)
	return b.memset.Contains(b.fxHash(value))
}
