package model

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"testing"

	"github.com/golang-module/carbon"
	"github.com/stretchr/testify/assert"
)

func TestGenerateHash(t *testing.T) {
	var block Block
	block.GenerateHash()
	hashLen := 32

	fmt.Println(hex.EncodeToString(block.Hash))
	assert.Equal(t, hashLen, len(block.Hash))
}

func TestValidateHash(t *testing.T) {
	var block Block
	hashStatus := block.ValidateHash([]byte("aa2db984d7a871233d4749b7beec401c52f62a2350a9b012159e0807b9eb9367"))
	fmt.Println(hashStatus)
	assert.True(t, hashStatus)
}

func TestValidateBlock(t *testing.T) {
	blockchain := InitBlockchain()
	prevBlock := blockchain.blocks[len(blockchain.blocks)-1]
	block := &Block{
		Data:      Student{},
		PrevHash:  prevBlock.Hash,
		Position:  prevBlock.Position + 1,
		Timestamp: carbon.Now().ToDateTimeString(),
	}
	block.GenerateHash()
	blockStatus := ValidateBlock(prevBlock, block)
	assert.True(t, blockStatus)
}

func TestAddBlock(t *testing.T) {
	var blockchain = InitBlockchain()
	student := Student{
		Name:       "Lawrence Segun",
		Age:        19,
		Class:      "400L",
		Department: "Computer Science",
	}

	h := md5.New()
	io.WriteString(h, student.Class+student.Name+student.Department)
	student.StudentID = fmt.Sprintf("%x", h.Sum(nil))

	blockchain.AddBlock(student)
	for _, block := range blockchain.blocks {
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Prev Hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %v\n", block.Data)
	}
}
