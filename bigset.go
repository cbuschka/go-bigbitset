package bigbitset

// BitSet is a set of bits. The bit position is counted from zero. Bits in
// this set can be flagged as set, but cannot be cleared afterwards. I.e. the fill
// grade converges into the direction of completeness only. Implementations of BitSet
// are considered to manage large counts of bits efficiently.
type BitSet interface {
	// Set the bit at position pos, counted from zero.
	Set(pos uint64)
	// Test if the bit at position pos is set, pos counted from zero.
	IsSet(pos uint64) bool
	IsEmpty() bool
	IsComplete() bool
}
