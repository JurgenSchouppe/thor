package proto

import (
	"testing"

	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/assert"
	"github.com/vechain/thor/block"
)

func TestGetBlockByIDResult(t *testing.T) {
	var g nilableBlock
	data, err := rlp.EncodeToBytes(&g)
	assert.Nil(t, err)

	g.Block = new(block.Builder).Build()
	assert.Nil(t, rlp.DecodeBytes(data, &g))
	assert.Nil(t, g.Block)

	g.Block = new(block.Builder).Build()
	data, err = rlp.EncodeToBytes(&g)
	assert.Nil(t, err)

	g.Block = nil

	assert.Nil(t, rlp.DecodeBytes(data, &g))
	assert.Equal(t, new(block.Builder).Build().Header().ID(), g.Block.Header().ID())
}