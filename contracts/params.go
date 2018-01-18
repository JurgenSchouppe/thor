package contracts

import (
	"math/big"

	"github.com/vechain/thor/thor"
)

// keys of governance params.
const (
	ParamRewardPercentage   = "reward-percentage"
	ParamEnergyGrowthFactor = "energy-growth-factor"
)

type params struct {
	contract
}

// PackInitialize packs input data of `Params.initialize` function.
func (p *params) PackInitialize(voting thor.Address) []byte {
	return p.mustPack("initialize", voting)
}

// PackGet packs input data of `Params.get` function.
func (p *params) PackGet(key string) []byte {
	return p.mustPack("get", key)
}

// UnpackGet unpacks output data of `Params.get` function.
func (p *params) UnpackGet(output []byte) *big.Int {
	var value *big.Int
	p.mustUnpack(&value, "get", output)
	return value
}

// PackSet packs input data of `Params.set` function.
func (p *params) PackSet(key string, value *big.Int) []byte {
	return p.mustPack("set", key, value)
}

func (p *params) PackPreset(key string, value *big.Int) []byte {
	return p.mustPack("preset", key, value)
}

// Params binder of `Params` contract.
var Params = params{mustLoad(
	thor.BytesToAddress([]byte("par")),
	"compiled/Params.abi",
	"compiled/Params.bin-runtime")}
