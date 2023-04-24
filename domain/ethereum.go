package domain

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/forta-network/forta-core-go/utils"
)

// Block is the intersection between parity and go-ethereum block
type Block struct {
	BaseFeePerGas    string        `json:"baseFeePerGas"`
	Difficulty       *string       `json:"difficulty"`
	ExtraData        *string       `json:"extraData"`
	GasLimit         *string       `json:"gasLimit"`
	GasUsed          *string       `json:"gasUsed"`
	Hash             string        `json:"hash"`
	LogsBloom        *string       `json:"logsBloom"`
	Miner            *string       `json:"miner"`
	MixHash          *string       `json:"mixHash"`
	Nonce            *string       `json:"nonce"`
	Number           string        `json:"number"`
	ParentHash       string        `json:"parentHash"`
	ReceiptsRoot     *string       `json:"receiptsRoot"`
	Sha3Uncles       *string       `json:"sha3Uncles"`
	Size             *string       `json:"size"`
	StateRoot        *string       `json:"stateRoot"`
	Timestamp        string        `json:"timestamp"`
	TotalDifficulty  *string       `json:"totalDifficulty"`
	Transactions     []Transaction `json:"transactions"`
	TransactionsRoot *string       `json:"transactionsRoot"`
	Uncles           []*string     `json:"uncles"`
}

func (b *Block) Age() (*time.Duration, error) {
	ts, err := b.GetTimestamp()
	if err != nil {
		return nil, err
	}
	age := time.Since(*ts)
	return &age, nil
}

func (b *Block) GetTimestamp() (*time.Time, error) {
	ts, err := utils.HexToBigInt(b.Timestamp)
	if err != nil {
		return nil, err
	}
	blockTsMs := ts.Mul(ts, big.NewInt(1000))
	result := time.Unix(0, int64(blockTsMs.Uint64())*int64(time.Millisecond))
	return &result, nil
}

// Transaction is the intersection between parity and go-ethereum transactions
type Transaction struct {
	BlockHash            string  `json:"blockHash"`
	BlockNumber          string  `json:"blockNumber"`
	From                 string  `json:"from"`
	Gas                  string  `json:"gas"`
	GasPrice             string  `json:"gasPrice"`
	Hash                 string  `json:"hash"`
	Input                *string `json:"input"`
	Nonce                string  `json:"nonce"`
	To                   *string `json:"to"`
	TransactionIndex     string  `json:"transactionIndex"`
	Value                *string `json:"value"`
	V                    string  `json:"v"`
	R                    string  `json:"r"`
	S                    string  `json:"s"`
	MaxFeePerGas         string  `json:"maxFeePerGas"`
	MaxPriorityFeePerGas string  `json:"maxPriorityFeePerGas"`
}

// LogEntry is a log item inside a receipt
type LogEntry struct {
	Address          *string   `json:"address"`
	BlockHash        *string   `json:"blockHash"`
	BlockNumber      *string   `json:"blockNumber"`
	Data             *string   `json:"data"`
	LogIndex         *string   `json:"logIndex"`
	Removed          *bool     `json:"removed"`
	Topics           []*string `json:"topics"`
	TransactionHash  *string   `json:"transactionHash"`
	TransactionIndex *string   `json:"transactionIndex"`
}

// ToTypesLog converts our type to go-ethereum type.
func (le LogEntry) ToTypesLog() (log types.Log) {
	if le.Address != nil {
		log.Address = common.HexToAddress(*le.Address)
	}
	if le.BlockHash != nil {
		log.BlockHash = common.HexToHash(*le.BlockHash)
	}
	if le.BlockNumber != nil {
		num, _ := hexutil.DecodeBig(*le.BlockNumber)
		log.BlockNumber = num.Uint64()
	}
	if le.Data != nil {
		log.Data = []byte(*le.Data)
	}
	if le.LogIndex != nil {
		num, _ := hexutil.DecodeBig(*le.LogIndex)
		log.Index = uint(num.Uint64())
	}
	if le.Removed != nil {
		log.Removed = *le.Removed
	}
	for _, topic := range le.Topics {
		log.Topics = append(log.Topics, common.HexToHash(*topic))
	}
	if le.TransactionHash != nil {
		log.TxHash = common.HexToHash(*le.TransactionHash)
	}
	if le.TransactionIndex != nil {
		num, _ := hexutil.DecodeBig(*le.TransactionIndex)
		log.TxIndex = uint(num.Uint64())
	}
	return
}

// TransactionReceipt is a result of a eth_getTransactionReceipt call
type TransactionReceipt struct {
	BlockHash         *string    `json:"blockHash"`
	BlockNumber       *string    `json:"blockNumber"`
	ContractAddress   *string    `json:"contractAddress"`
	CumulativeGasUsed *string    `json:"cumulativeGasUsed"`
	From              *string    `json:"from"`
	GasUsed           *string    `json:"gasUsed"`
	Logs              []LogEntry `json:"logs"`
	LogsBloom         *string    `json:"logsBloom"`
	Status            *string    `json:"status"`
	To                *string    `json:"to"`
	TransactionHash   *string    `json:"transactionHash"`
	TransactionIndex  *string    `json:"transactionIndex"`
}

// TraceAction is an element of a trace_block Trace response
type TraceAction struct {
	CallType      *string `json:"callType"`
	To            *string `json:"to"`
	Input         *string `json:"input"`
	From          *string `json:"from"`
	Gas           *string `json:"gas"`
	Value         *string `json:"value"`
	Init          *string `json:"init"`
	Address       *string `json:"address"`
	Balance       *string `json:"balance"`
	RefundAddress *string `json:"refundAddress"`
}

// TraceResult is a result element of a trace_block Trace response
type TraceResult struct {
	Output  *string `json:"output"`
	GasUsed *string `json:"gasUsed"`
	Address *string `json:"address"`
	Code    *string `json:"code"`
}

// Trace is a specific traced action in a transaction
type Trace struct {
	Action              TraceAction  `json:"action"`
	BlockHash           *string      `json:"blockHash"`
	BlockNumber         *int         `json:"blockNumber"`
	Result              *TraceResult `json:"result"`
	Subtraces           int          `json:"subtraces"`
	TraceAddress        []int        `json:"traceAddress"`
	TransactionHash     *string      `json:"transactionHash"`
	TransactionPosition *int         `json:"transactionPosition"`
	Type                string       `json:"type"`
	Error               *string      `json:"error"`
}

// HeaderCh provides new block headers.
type HeaderCh <-chan *types.Header

// ClientSubscription abstracts away the subscription implementation.
type ClientSubscription interface {
	Err() <-chan error
	Unsubscribe()
}
