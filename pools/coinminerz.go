package pools

import (
	"fmt"
	"time"

	"github.com/vertcoin-project/one-click-miner-vnext/util"
)

var _ Pool = &CoinMinerz{}

type CoinMinerz struct {
	Address           string
	LastFetchedPayout time.Time
	LastPayout        uint64
}

func NewCoinMinerz(addr string) *CoinMinerz {
	return &CoinMinerz{Address: addr}
}

func (p *CoinMinerz) GetPendingPayout() uint64 {
	jsonPayload := map[string]interface{}{}
	err := util.GetJson(fmt.Sprintf("http://badhasher.com/api/v1/Vertcoin/miner?address=%s", p.Address), &jsonPayload)
	if err != nil {
		return 0
	}

	el, ok := jsonPayload["payments"].(map[string]interface{})
	if !ok {
		return 0
	}

	balance, ok := el["balance"].(float64)
	if !ok {
		return 0
	}

	immature, ok := el["immature"].(float64)
	if !ok {
		return 0
	}

	generate, ok := el["generate"].(float64)
	if !ok {
		return 0
	}

	vtc := balance + immature + generate
	vtc *= 100000000
	return uint64(vtc)
}

func (p *CoinMinerz) GetStratumUrl() string {
	return "stratum+tcp://stratum.coinminerz.com:3517"
}

func (p *CoinMinerz) GetUsername() string {
	return p.Address
}

func (p *CoinMinerz) GetPassword() string {
	return "x"
}

func (p *CoinMinerz) GetID() int {
	return 10
}

func (p *CoinMinerz) GetName() string {
	return "Coinminerz.com"
}

func (p *CoinMinerz) GetFee() float64 {
	return 0.5
}

func (p *CoinMinerz) OpenBrowserPayoutInfo(addr string) {
	util.OpenBrowser(fmt.Sprintf("https://badhasher.com/miner/Vertcoin/%s", addr))
}
