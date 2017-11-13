package parser

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/forchain/ethrpc"
	"io/ioutil"
	"log"
	"runtime"
	"sync"
	"time"
)

const BLOCKS_PER_FILE = 10000



var rpc_ *ethrpc.EthRPC
var max_block_ = 4500000

func parseBlock(_file int, _wg *sync.WaitGroup, _outDir string) {
	defer _wg.Done()

	b := new(bytes.Buffer)
	w, err := gzip.NewWriterLevel(b, gzip.BestSpeed)
	if err != nil {
		log.Fatal(err)
	}
	from := _file * BLOCKS_PER_FILE
	to := from + BLOCKS_PER_FILE
	if to > max_block_ {
		to = max_block_
	}
	for i := from; i < to; i++ {
		if b, err := rpc_.EthGetBlockByNumber(i, true); err == nil {
			w.Write([]byte(fmt.Sprintf("<%v> <p> <%v> .\n", b.Hash, b.ParentHash)))
			ts := time.Unix(int64(b.Timestamp), 0).Format(time.RFC3339)
			w.Write([]byte(fmt.Sprintf("<%v> <ts> \"%v\"^^<xs:dateTime> .\n", b.Hash, ts)))
			if len(b.Transactions) > 0 {
				for _, t := range b.Transactions {
					w.Write([]byte(fmt.Sprintf("<%v> <tx> <%v> .\n", t.BlockHash, t.Hash)))
					w.Write([]byte(fmt.Sprintf("<%v> <f> <%v> .\n", t.Hash, t.From)))
					w.Write([]byte(fmt.Sprintf("<%v> <t> <%v> .\n", t.Hash, t.To)))
					w.Write([]byte(fmt.Sprintf("<%v> <v> \"%v\"^^<xs:int> .\n", t.Hash, t.Value.String())))
				}
			}
		} else {
			log.Println(err, i, b)
			time.Sleep(time.Second)
			i--
		}
	}

	w.Close()

	fileName := fmt.Sprintf("%v/%v.rdf.gz", _outDir, _file)
	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
		log.Fatal(err)
	}
	log.Println(fileName)
}

func Parse(_rpc string, _out string) {

	rpc_ = ethrpc.NewEthRPC(_rpc)

	var err error

	cpuNum := runtime.NumCPU()

	wg := new(sync.WaitGroup)
	num := 0

	if num, err = rpc_.EthBlockNumber(); err != nil || num == 0 {
		if s, err := rpc_.EthSyncing(); err == nil && s != nil {
			num = s.CurrentBlock
		}
	}
	if err != nil {
		log.Fatal(err)
	}
	if num == 0 {
		num = max_block_
	} else {
		max_block_ = num
	}
	log.Println(num)

	files := (num-1)/BLOCKS_PER_FILE + 1
	for i := 0; i < files; i++ {
		wg.Add(1)
		go parseBlock(i, wg, _out)
		if (i+1)%cpuNum == 0 {
			wg.Wait()
		}
	}
	wg.Wait()

	if err != nil {
		log.Fatal(err)
	}
}
