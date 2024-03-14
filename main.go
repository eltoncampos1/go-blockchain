package main

import (
    "bytes"
    "crypto/sha256"
    "strconv"
    "time"
    "fmt"
)

type Block struct {
    Timestamp int64
    Data []byte
    PrevBlockHash []byte
    hash []byte
}

type Blockchain struct {
    blocks []*Block
}

func (bc *Blockchain) AddBlock(data string){
    prevBlock := bc.blocks[len(bc.blocks) - 1]
    newBlock := NewBlock(data, prevBlock.hash)
    bc.blocks = append(bc.blocks, newBlock)
}

func (b *Block) Sethash(){
    timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
    headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
    hash := sha256.Sum256(headers)

    b.hash = hash[:]
}

func NewGenesisBlock() *Block{
    return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain{
    return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func NewBlock(data string, prevBlockHash []byte) *Block{
    block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
    block.Sethash()
    return block
}

func main(){
    bc := NewBlockchain()

    bc.AddBlock("Send 1 BTC to Jhon Doe")
    bc.AddBlock("Send 2 more BTC to Jhon Doe")

    for _, block := range bc.blocks{
        fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
        fmt.Printf("Data: %s\n", block.Data)
        fmt.Printf("Hash: %x\n", block.hash)
        fmt.Println()
    }
}
