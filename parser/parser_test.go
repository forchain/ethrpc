package parser

import (
	"github.com/forchain/ethrpc"
	"testing"
)

func TestParse(t *testing.T) {
	invalidNum := 9999999

	rpc_ = ethrpc.NewEthRPC(ADDRESS)
	b, err := rpc_.EthGetBlockByNumber(invalidNum, true)
	t.Log(b, err)

}
