package Handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Ibukun-tech/go_chain/pkg/Handler"
)

func TestWriteBlockPost(t *testing.T) {
	reqBody := `{"id":"335355355", "title":"FIRST bOOK", "publish_date":"23-4-25", "isbn":"2343", "author":"ibk"}`
	req, err := http.NewRequest("POST", "/new", strings.NewReader(reqBody))
	if err != nil {
		fmt.Println(err)
	}
	w := httptest.NewRecorder()
	Handler.NewBook(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %d want %d", status, http.StatusOK)
	}
}

func TestWritePost(t *testing.T) {
	postReq := `{"book_id":"345", "user":"Ibukun", "checkout_date":"23-4-23", "is_genesis":false}`
	req, err := http.NewRequest("POST", "/", strings.NewReader(postReq))
	if err != nil {
		t.Log(err)
	}
	w := httptest.NewRecorder()
	Handler.WriteBlock(w, req)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %d want %d", status, http.StatusOK)
	}
}
