package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
)

const (
	DB_PATH = "/home/vardhan.kottamasu@npci.org.in/.leveldb/demodb"
)

type Block struct {
	blockNumber       int         `json:"blocknumber"`
	transactionList   []txn       `json:transactions`
	status            blockStatus `json:blockstatus`
	previousBlockHash string      `json:previousblockhash`
	blockHash         string      `json:blockhash`
}

type blockStatus string

func (b blockStatus) String() string {
	return string(b)
}

const (
	COMMITTED blockStatus = "committed"
	PENDING   blockStatus = "pending"
)

type txn struct {
	txnID     int
	timeStamp time.Time
	valid     bool		`json:"valid"`
	value     string	`json:"value"`
	version   string	`json:"version"`
}

type iBlock interface {
	pushValidTransaction(transaction txn)
	updateBlock(transaction txn)
}

func isTransactionValid(transaction txn) {
	if transaction.valid == true {
		block := &Block{}
		block.pushTransaction(transaction)
		writeBlockToLedger(*block)
	}
}

func (b Block) pushTransaction(transaction txn) {
	//TODO: This count variable has to be fixed, as it resets everytimr code reruns
	var count int = 0
	if transaction.valid == true {
		db, err := leveldb.OpenFile(DB_PATH, nil)
		if err != nil {
			fmt.Println("Error. DB can't be created")
		} else {
			count++
			db.Put([]byte("key"+string(rune(count))), []byte(transaction.value), nil)
		}
		defer db.Close()
	}
}

func (b Block) updateBlock(block Block) {
	block.status = COMMITTED
}

func readDB(key string) {
	db, err := leveldb.OpenFile(DB_PATH, nil)
	if err == nil {
		data, err := db.Get([]byte(key), nil)
		if err == nil {
			fmt.Printf("%v", data)
		}
	}
}

func writeBlockToLedger(block Block) {
	ledger,err := os.OpenFile("ledger.json",os.O_APPEND|os.O_CREATE,0777)
	if(err!=nil){
		fmt.Println("Error creating ledger file")
	}else{
		block.blockHash = md5.Sum([]byte("sample"))
	}
}

func main() {
// Once errors with all the functions are cleared, you can start implementing the functions here
}
