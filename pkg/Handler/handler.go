package Handler

import (
	"crypto/md5"
	"io"
	"net/http"
)

func NewBook(w http.ResponseWriter, r *http.Request) {
	// s := new(Model.Book)
	h := md5.New()
	io.WriteString(h, "hhshs")
}
