package Model_test

import (
	"testing"

	"github.com/Ibukun-tech/go_chain/pkg/Model"
)

func TestBook(t *testing.T) {
	Book := Model.Book{
		Id:          "2",
		Title:       "Ibukun TextBook",
		Author:      "Oyetunji",
		PublishDate: "2023-3-4",
		ISBN:        "4443434343434",
	}
	if Book.Id != "2" && Book.Title != "Ibukun TextBook" && Book.Author != "Oyetunji" && Book.PublishDate != "2023-3-4" && Book.ISBN != "4443434343434" {
		t.Error("Something went wrong")
	}
}
