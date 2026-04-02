// Blockchain Block Full Validator
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Index     int
	Timestamp int64
	Data      string
	PrevHash  string
	Hash      string
	Difficulty int
}

func calculateHash(b Block) string {
	record := strconv.Itoa(b.Index) + strconv.FormatInt(b.Timestamp, 10) + b.Data + b.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}

func ValidateBlock(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	prefix := string(make([]byte, newBlock.Difficulty))
	for i := range prefix {
		prefix[i] = '0'
	}
	if newBlock.Hash[:newBlock.Difficulty] != prefix {
		return false
	}
	return true
}

func main() {
	genesis := Block{0, time.Now().Unix(), "Genesis Block", "0", "", 2}
	genesis.Hash = calculateHash(genesis)
	block1 := Block{1, time.Now().Unix(), "Transfer 10 GCT", genesis.Hash, "", 2}
	block1.Hash = calculateHash(block1)
	fmt.Println("Block Valid:", ValidateBlock(block1, genesis))
}
