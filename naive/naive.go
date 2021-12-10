package naive

import (
	"github.com/cbuschka/go-bigbitset"
)

type NaiveBitSet struct {
	blocks []byte
	size   int
}

func CreateNaiveBitSet(size uint64) bigbitset.BitSet {
	sizei := int(size)

	bs := NaiveBitSet{blocks: []byte{}, size: 0}
	bs.ensureCapacity(sizei/8 + 1)
	bs.size = sizei
	return &bs
}

func (bs *NaiveBitSet) IsSet(pos uint64) bool {

	posi := int(pos)

	byteIndex := posi / 8
	bitIndex := posi % 8

	if posi >= bs.size {
		return false
	}

	return bs.blocks[byteIndex]&(0x01<<bitIndex) != 0
}

func (bs *NaiveBitSet) Set(pos uint64) {

	posi := int(pos)

	byteIndex := posi / 8
	bitIndex := posi % 8

	bs.ensureCapacity(posi/8 + 1)

	bs.blocks[byteIndex] = bs.blocks[byteIndex] | 0x01<<bitIndex
	bs.size = posi + 1
}

func (bs *NaiveBitSet) ensureCapacity(requiredBytes int) {
	if len(bs.blocks) > requiredBytes {
		return
	}

	newBlocks := make([]byte, requiredBytes)
	copy(newBlocks, bs.blocks)
	bs.blocks = newBlocks
}

func (bs *NaiveBitSet) IsComplete() bool {

	for i := 0; i < bs.size; i++ {
		if !bs.IsSet(uint64(i)) {
			return false
		}
	}

	return true
}
