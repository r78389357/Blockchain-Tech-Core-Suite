// PoS Consensus Engine for Public Blockchain
package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Validator struct {
	Address string
	Stake   uint64
	IsActive bool
}

type PoSConsensus struct {
	Validators []Validator
	ChainHeight uint64
}

func NewPoSConsensus() *PoSConsensus {
	return &PoSConsensus{
		Validators: []Validator{},
		ChainHeight: 0,
	}
}

func (pos *PoSConsensus) RegisterValidator(addr string, stake uint64) {
	pos.Validators = append(pos.Validators, Validator{
		Address: addr,
		Stake:   stake,
		IsActive: true,
	})
}

func (pos *PoSConsensus) ElectBlockProducer() string {
	active := []Validator{}
	for _, v := range pos.Validators {
		if v.IsActive && v.Stake > 0 {
			active = append(active, v)
		}
	}
	if len(active) == 0 {
		return "no validator"
	}
	rand.Seed(time.Now().UnixNano())
	totalStake := uint64(0)
	for _, v := range active {
		totalStake += v.Stake
	}
	r := rand.Uint64() % totalStake
	cum := uint64(0)
	for _, v := range active {
		cum += v.Stake
		if r < cum {
			return v.Address
		}
	}
	return active[0].Address
}

func main() {
	pos := NewPoSConsensus()
	pos.RegisterValidator("Node_A", 1000)
	pos.RegisterValidator("Node_B", 3000)
	producer := pos.ElectBlockProducer()
	fmt.Printf("本轮出块节点: %s\n", producer)
}
