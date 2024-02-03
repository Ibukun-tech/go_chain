package Handler

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	helper "github.com/Ibukun-tech/go_chain/pkg/Helper"
	"github.com/Ibukun-tech/go_chain/pkg/Model"
)

var BlockChainConnect Model.BlockChain

func NewBook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		book := new(Model.Book)
		err := json.NewDecoder(r.Body).Decode(book)
		fmt.Println(book.Author)
		if err != nil {
			fmt.Println("No one error in connecting to he get request")
		}
		h := md5.New()
		io.WriteString(h, book.ISBN+book.PublishDate)
		book.Id = hex.EncodeToString(h.Sum(nil))

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		type Result struct {
			Status  bool
			Message string
			Data    interface{}
		}
		vals := Result{
			Status:  true,
			Message: "Its was successful",
			Data:    book,
		}
		json.NewEncoder(w).Encode(vals)
	}

	// s := new(Model.Book)
}
func GetBlockchain(w http.ResponseWriter, r *http.Request) {
	jbyte, _ := json.MarshalIndent(BlockChainConnect, "", "")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(string(jbyte))
}
func WriteBlock(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		st := new(Model.BookCheckOut)
		json.NewDecoder(r.Body).Decode(st)
		resp, err := json.MarshalIndent(st, "", "")
		if err != nil {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err)
		}
		helper.AddBlock(&BlockChainConnect, *st)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		type Result struct {
			Status  bool
			Message string
			Data    interface{}
		}
		vals := Result{
			Status:  true,
			Message: "New Block created",
			Data:    string(resp),
		}
		json.NewEncoder(w).Encode(vals)

	}

}
