package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"time"

	"github.com/Ibukun-tech/go_chain/pkg/Model"
)

func GenerateHash(b *Model.Block) {
	bytes, _ := json.Marshal(b.Data)
	data := string(rune(b.Pos)) + b.Timestamp + string(bytes) + b.PrevHash
	hash := sha256.New()
	b.Hash = hex.EncodeToString(hash.Sum([]byte(data)))
}

func CreateBlock(prvBlock *Model.Block, data Model.BookCheckOut) *Model.Block {
	block := new(Model.Block)
	if data.IsGenesis {
		block.Pos = 0
	} else {
		block.Pos = prvBlock.Pos + 1
	}
	block.Timestamp = time.Now().String()
	block.Data = data
	block.PrevHash = prvBlock.PrevHash
	GenerateHash(block)
	// block.Data.IsGenesis
	return block
}
func ValidBlock(block, prvBlock *Model.Block) bool {
	if block.PrevHash != prvBlock.Hash {
		return false
	}
	if !ValidateHash(block.Hash, block) {
		return false
	}
	if prvBlock.Pos+1 != block.Pos {
		return false
	}
	return true
}
func ValidateHash(hash string, b *Model.Block) bool {
	GenerateHash(b)
	return hash == b.Hash
}
func AddBlock(blks *Model.BlockChain, data Model.BookCheckOut) {
	lastBlock := blks.Blocks[len(blks.Blocks)-1]
	block := CreateBlock(lastBlock, data)
	if ValidBlock(block, lastBlock) {
		blks.Blocks = append(blks.Blocks, block)
	}
}
