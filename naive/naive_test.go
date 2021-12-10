package naive

import (
	"github.com/cbuschka/go-bigbitset"
	"github.com/cbuschka/go-bigbitset/testsuite"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestNaiveBitSet(t *testing.T) {
	ts := testsuite.CreateTestsuite(func(size uint64) bigbitset.BitSet {
		return CreateNaiveBitSet(size)
	})
	suite.Run(t, ts)
}
