package collection

import (
	"fmt"
	"strconv"

	"github.com/infoCollection/gpool"

	"github.com/elgs/gojq"
	. "github.com/infoCollection/database"
	curl "github.com/mikemintang/go-curl"
)

type PostData struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      int           `json:"id"`
}

const URL = "http://47.101.47.204:8545"

var HEADERS = map[string]string{
	"content-type": "application/json; charset=UTF-8",
}

// 获取最新区块编号
func getBlockNumber() (num int64) {

	pd := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_blockNumber",
		"params":  make([]interface{}, 0),
		"id":      0,
	}
	req := curl.NewRequest()
	res, err := req.
		SetUrl(URL).
		SetHeaders(HEADERS).
		SetPostData(pd).
		Post()

	if err != nil {
		fmt.Println(err)
	} else {
		if res.IsOk() {
			parser, err := gojq.NewStringQuery(res.Body)
			if err != nil {
				fmt.Println(err)
			}
			result, err := parser.Query("result")
			intResult, err := strconv.ParseInt(result.(string), 0, 64)

			num = intResult

		}
	}
	return num
}

// 根据区块编号获取区块信息 遍历区块交易记录
func getBlockByNumber(num int64, pool *gpool.Pool) {
	stringNum := strconv.FormatInt(num, 16)
	params := make([]interface{}, 2)
	params[0] = "0x" + stringNum
	params[1] = true
	pd := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getBlockByNumber",
		"params":  params,
		"id":      0,
	}
	req := curl.NewRequest()
	res, err := req.
		SetUrl(URL).
		SetHeaders(HEADERS).
		SetPostData(pd).
		Post()
	if err != nil {
		fmt.Println(err)
	} else {
		if res.IsOk() {

			body, err := gojq.NewStringQuery(string(res.Body))
			if err != nil {
				return
			}
			timestamp, err := body.Query("result.timestamp")
			ts, err := body.Query("result.transactions")
			if err != nil {
				fmt.Println(err)
			}

			tsLens := len(ts.([]interface{}))

			var transtion Transtion
			transtion.Timestamp = timestamp.(string)

			for i := 0; i < tsLens; i++ {
				transtion.Timestamp = timestamp.(string)
				from, err := body.Query(fmt.Sprintf("result.transactions.[%d].from", i))
				if err != nil {
					fmt.Println(err)
				}
				transtion.From = reflectTypes(from)
				to, err := body.Query(fmt.Sprintf("result.transactions.[%d].to", i))
				if err != nil {
					fmt.Println(err)
				}
				transtion.To = reflectTypes(to)
				value, err := body.Query(fmt.Sprintf("result.transactions.[%d].value", i))
				if err != nil {
					fmt.Println(err)
				}
				transtion.Value = reflectTypes(value)

				InsertData(&transtion)
			}
			pool.Done()
		}
	}
}

func reflectTypes(items ...interface{}) (str string) {
	for _, v := range items {
		switch v.(type) {
		case string:
			str = v.(string)
		default:
			str = ""
		}
	}
	return str
}

func Collection() {
	var total int64
	total = getBlockNumber()
	startNum := total - 500

	pool := gpool.New(5)
	var i int64
	for i = startNum; i < total; i++ {
		pool.Add(1)
		go getBlockByNumber(i, pool)
	}
	pool.Wait()

}
