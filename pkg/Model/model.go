package Model

// Item to be checked out
type Book struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishDate string `json:"publish_date"`
	ISBN        string `json:"isbn"`
}

// This respresents the data of book to be checked out
type BookCheckOut struct {
	BookId       string `json:"book_id"`
	User         string `json:"user"`
	CheckoutDate string `json:"checkout_date"`
	IsGenesis    bool   `json:"is_genesis"`
}

// Block represent the individual block on the blockchain
type Block struct {
	Pos       int
	Data      BookCheckOut
	Timestamp string
	Hash      string
	PrevHash  string
}

// This represnt the core blockchain implementation
type BlockChain struct {
	Blocks []*Block
}

// For creating an instance of blockchain
func BlocksStruct() *BlockChain {
	return new(BlockChain)
}
