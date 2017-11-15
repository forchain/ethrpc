package parser

import (
	"github.com/forchain/ethrpc"
	"testing"
	"sync"
	"math/big"
)

const ADDRESS = "http://10.147.18.28:8545"

func TestParse(t *testing.T) {
	invalidNum := 9999999

	rpc_ = ethrpc.NewEthRPC(ADDRESS)
	b, err := rpc_.EthGetBlockByNumber(invalidNum, true)
	t.Log(b, err)

}

func TestSingle(t *testing.T) {

	rpc_ = ethrpc.NewEthRPC(ADDRESS)
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go parseBlock(11, wg, "/tmp")
	go parseBlock(12, wg, "/tmp")
	wg.Wait()
}

func TestFloat(t *testing.T) {
	s := "123456789000000000000"
	i, _ := new(big.Int).SetString(s, 10)

	x := new(big.Float).SetInt(i)
	y := new(big.Float).SetInt(big.NewInt(1000000000000000000))
	z := new(big.Float).Quo(x, y)
	t.Log(z)
}
