package main

import (
	// "crypto/md5"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	// "strconv"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
)

const (
	DB_PATH = "/home/vardhan.kottamasu@npci.org.in/.leveldb/demodb"
)

type Block struct {
	BlockNumber       int         `json:"blocknumber"`
	TransactionList   []txn       `json:"transactions"`
	Status            blockStatus `json:"blockstatus"`
	PreviousBlockHash string      `json:"previousblockhash"`
	BlockHash         string      `json:"blockhash"`
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
	TxnID     int
	TimeStamp time.Time
	Valid     bool		
	Value     string	
	Version   string	
}

type iBlock interface {
	pushValidTransaction(transaction txn)
	updateBlock(transaction txn)
}

func isTransactionValid(transaction txn) {
	if transaction.Valid == true {
		block := &Block{}
		block.pushTransaction(transaction)
		writeBlockToLedger(*block)
	}
}

func (b Block) pushTransaction(transaction txn) {
	//TODO: This count variable has to be fixed, as it resets everytimr code reruns
	var count int = 0
	if transaction.Valid == true {
		db, err := leveldb.OpenFile(DB_PATH, nil)
		if err != nil {
			fmt.Println("Error. DB can't be created")
		} else {
			count++
			db.Put([]byte("key"+string(rune(count))), []byte(transaction.Value), nil)
		}
		defer db.Close()
	}
}

func (b Block) updateBlock(block Block) {
	block.Status = COMMITTED
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
	ledger,err := os.OpenFile("ledger.json",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if(err!=nil){
		fmt.Println("Error creating ledger file")
	}else{
		startTime:=time.Now()
		blockJson,err := json.Marshal(block)
		fmt.Println(blockJson)
		if(err!=nil){
			fmt.Println("Error opening/creating ledger.json file")
		}else{
			ledger.WriteString(string(blockJson)+",")
			fmt.Println("Data written successfully!!")
			endTime := time.Since(startTime)
			fmt.Println("Time taken for block processing is",endTime)					}
		}
	}

	func fetchAllBlocks(){
		ledger,err:=os.Open("ledger.json")
		if err!=nil{
			fmt.Println("Error reading file")
			return
		}
		// defer ledger.Close()
		blockDetails,err := ioutil.ReadAll(ledger)
		
		if(err!=nil){
			fmt.Println("Error reading file")
		}
		var blocks[] Block
		val:=json.Unmarshal(blockDetails, &blocks)
			fmt.Println(val)

		for _, block := range blocks {
			fmt.Println(block.Status)
			fmt.Println(block.BlockNumber)
			fmt.Println(block.BlockHash)
			fmt.Println(block.PreviousBlockHash)
		}
	}

func main() {
// Once errors with all the functions are cleared, you can start implementing the functions here
	// byteArrayForBlock := []byte{}
	block := &Block{}
	block.BlockHash=fmt.Sprintf("%x",md5.Sum([]byte(block.BlockHash)))
	block.PreviousBlockHash=fmt.Sprintf("%x",md5.Sum([]byte(block.PreviousBlockHash)))
	block.Status="valid"
	block.BlockNumber=1
	// byteArrayForBlock = append(byteArrayForBlock,[]byte(block.blockHash)...)
	// byteArrayForBlock = append(byteArrayForBlock,[]byte(block.previousBlockHash)...)
	// byteArrayForBlock = append(byteArrayForBlock,[]byte(strconv.Itoa(block.blockNumber))...)
	// byteArrayForBlock = append(byteArrayForBlock,[]byte(block.status)...)
	// hash:=md5.Sum(byteArrayForBlock)
	// fmt.Printf("%x\n",byteArrayForBlock)

	// fmt.Printf("%x",hash)
	// writeBlockToLedger(*block)
	// fmt.Println(block)
	fetchAllBlocks()
	// fetchBlockByNum(1)

}
