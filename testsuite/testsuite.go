package testsuite

import (
	"github.com/cbuschka/go-bigbitset"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type Testsuite struct {
	suite.Suite
	createBitSet func(uint64) bigbitset.BitSet
}

func CreateTestsuite(createBitSetFunc func(uint64) bigbitset.BitSet) *Testsuite {
	return &Testsuite{createBitSet: createBitSetFunc}
}

func (ts *Testsuite) TestZeroSizedBitSetIsComplete() {
	bs := ts.createBitSet(0)
	assert.True(ts.T(), bs.IsComplete())
}

func (ts *Testsuite) TestSetAndTestSingleBitBitSet() {
	bs := ts.createBitSet(1)
	assert.False(ts.T(), bs.IsComplete())
	bs.Set(0)
	assert.True(ts.T(), bs.IsSet(0))
	assert.True(ts.T(), bs.IsComplete())
}
