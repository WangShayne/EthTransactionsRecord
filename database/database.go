package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// type Block struct {
// 	Number           string `json:"number"`
// 	Hash             string `json:"hash"`
// 	ParentHash       string `json:"parentHash"`
// 	Nonce            string `json:"nonce"`
// 	Sha3Uncles       string `json:"sha3Uncles"`
// 	LogsBloom        string `json:"logsBloom"`
// 	TransactionsRoot string `json:"transactionsRoot"`
// 	StateRoot        string `json:"stateRoot"`
// 	ReceiptsRoot     string `json:"receiptsRoot"`
// 	Miner            string `json:"miner"`
// 	Difficulty       string `json:"difficulty"`
// 	TotalDifficulty  string `json:"totalDifficulty"`
// 	ExtraData        string `json:"extraData"`
// 	Size             string `json:"size"`
// 	GasLimit         string `json:"gasLimit"`
// 	GasUsed          string `json:"gasUsed"`
// 	Timestamp        string `json:"timestamp"`
// 	CreateAt         int64  `json:"create_at"`
// }

type Transtion struct {
	Timestamp string `json:"timestamp"`
	From      string `json:"from"`
	To        string `json:"to"`
	Value     string `json:"value"`
}

func OpenSQL() error {
	var err error
	DB, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/ethereum_transtions?parseTime=true")
	if err != nil {
		log.Fatalln(err)
	}

	DB.SetMaxIdleConns(20)
	DB.SetMaxOpenConns(100)

	if err := DB.Ping(); err != nil {
		log.Fatalln(err)
	}
	return nil
}

func InsertData(transtion *Transtion) {

	sqlStr := fmt.Sprintf("INSERT INTO ethereum_transtions( timestamp, from_address, to_address, value_total) VALUES(?, ?, ?, ?)")
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("insert failed,err%v\n", err)
	}
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	result, err := stmt.Exec(
		transtion.Timestamp,
		transtion.From,
		transtion.To,
		transtion.Value,
	)
	if err != nil {
		fmt.Println(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get last insert id failed, err:%v\n", err)
		return
	}
	fmt.Println(id)
}

func QueryAllData() {
	sqlStr := `SELECT * FROM ethereum_block`
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	result, err := stmt.Query()
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Println(result)
}
