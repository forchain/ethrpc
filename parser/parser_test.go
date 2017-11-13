package parser

import (
	"github.com/forchain/ethrpc"
	"testing"
	"sync"
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
