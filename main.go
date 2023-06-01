package main

import (
	"crypto/md5"
	"fmt"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
)

type Block struct {
	blockNumber       int
	transactionList   []txn
	previousBlockHash string
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
	txnID     int         `json:""`
	valid     bool        `json:""`
	timeStamp time.Time   `json:""`
	status    blockStatus `json:""`
	value     string      ``
}

type iBlock interface {
	pushValidTransaction(transaction txn)
	updateBlock(transaction txn)
}

func (b Block) pushValidTransaction(transaction txn) {
	if transaction.status == COMMITTED {
		db, err := leveldb.OpenFile("/home/vardhan.kottamasu@npci.org.in/.leveldb/demodb", nil)
		if err == nil {
			db.Put([]byte("key3"), []byte(transaction), nil)
		}
		defer db.Close()
	}

}

func readDB(key string) {
	db, err := leveldb.OpenFile("/home/vardhan.kottamasu@npci.org.in/.leveldb/demodb", nil)
	if err == nil {
		data, err := db.Get([]byte("key3"), nil)
		if err == nil {
			fmt.Printf("%v", data)
		}
	}
}

func main() {
	s := "strong"
	str := md5.Sum([]byte(s))
	block := Block{}
	block.pushValidTransaction(s)
	fmt.Printf("%x", str)
	readDB("key3")
}
