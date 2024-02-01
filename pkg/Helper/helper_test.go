package helper_test

import (
	"testing"

	helper "github.com/Ibukun-tech/go_chain/pkg/Helper"
	"github.com/Ibukun-tech/go_chain/pkg/Model"
)

func TestGenerateHash(t *testing.T) {
	block := &Model.Block{
		Pos:       1,
		Timestamp: "2022-02-03",
		PrevHash:  "previousFunction",
		Data: Model.BookCheckOut{
			BookId:       "123",
			User:         "Ibukunoluwa",
			CheckoutDate: "2023-04-23",
			IsGenesis:    true,
		},
	}
	helper.GenerateHash(block)

	if block.Hash == "" {
		t.Error("Expected non-empty hash, got an empty string")
	}

}

func TestCreateBlock(t *testing.T) {
	prevBlock := &Model.Block{
		Pos:       1,
		Timestamp: "202-04-23",
		Hash:      "strtsgsgssg",
	}
	data := Model.BookCheckOut{
		BookId:       "123",
		User:         "Ibukunoluwa Oyetunji",
		CheckoutDate: "2024-23-12",
		IsGenesis:    false,
	}
	blck := helper.CreateBlock(prevBlock, data)
	if blck.Pos != prevBlock.Pos+1 {
		t.Errorf("expected  %d and prevBlock %d", prevBlock.Pos+1, blck.Pos)
	}
}
