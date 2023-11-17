package model

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"

	"github.com/golang-module/carbon"
)

type Student struct {
	StudentID  string `json:"student_id,omitempty"`
	Name       string `json:"name,omitempty"`
	Age        int    `json:"age,omitempty"`
	Class      string `json:"class,omitempty"`
	Department string `json:"department,omitempty"`
	IsGenesis  bool   `json:"is_genesis"`
}

type Block struct {
	Data      Student
	PrevHash  []byte
	Hash      []byte
	Position  int
	Timestamp string
}

type Blockchain struct {
	blocks []*Block
}

func (b *Block) GenerateHash() {
	bytes, _ := json.Marshal(b.Data)
	data := string(rune(b.Position)) + b.Timestamp + string(bytes) + string(b.PrevHash)
	hasher := sha256.New()
	hasher.Write([]byte(data))
	b.Hash = hasher.Sum(nil)
}

func NewBlock(prevBlock *Block, student Student) *Block {
	block := &Block{
		Data:      student,
		PrevHash:  prevBlock.PrevHash,
		Position:  prevBlock.Position + 1,
		Timestamp: carbon.Now().ToDateTimeString(),
	}

	block.GenerateHash()
	return block
}

func (b *Block) ValidateHash(hash []byte) bool {
	b.GenerateHash()
	// fmt.Println(hex.EncodeToString(b.Hash))
	// fmt.Println(string(hash))
	return hex.EncodeToString(b.Hash) == string(hash)
}

func ValidateBlock(prevBlock, block *Block) bool {
	prevBlockHash := hex.EncodeToString(prevBlock.Hash)
	blockPreviousHash := hex.EncodeToString(block.PrevHash)
	if prevBlockHash != blockPreviousHash {
		return false
	}
	if prevBlock.Position+1 != block.Position {
		return false
	}

	blockByte := []byte(hex.EncodeToString([]byte(block.Hash)))

	return block.ValidateHash(blockByte)
}

func (bc *Blockchain) AddBlock(student Student) {
	// This return the last block added to the blockchain
	prevBlock := bc.blocks[len(bc.blocks)-1]
	block := NewBlock(prevBlock, student)
	if ValidateBlock(prevBlock, block) {
		bc.blocks = append(bc.blocks, block)
	}
}

func InitBlockchain() *Blockchain {
	return &Blockchain{
		[]*Block{GenesisBlock()},
	}
}

func GenesisBlock() *Block {
	return NewBlock(&Block{}, Student{IsGenesis: true})
}
