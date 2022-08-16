package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	book "github.com/darrSonik/crud-go/models"
)

type bookTuple struct {
	dbIndex  int
	bookData book.Book
}

func searchBook(b book.Books, bookid string) (bookTuple, error) {
	for index, item := range b {
		if item.ID == bookid {
			return bookTuple{
				dbIndex:  index,
				bookData: *item,
			}, nil
		}
	}

	return bookTuple{}, errors.New("Not Found")
}

func HandlerBook(b *book.Books, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	url := strings.Split(r.URL.Path, "/")

	if len(url) != 3 {
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(`{"message": "Not found"}`))

		return
	}

	foundBook, notfound := searchBook(*b, url[2])

	if notfound != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Not found"}`))

		return
	}

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(foundBook.bookData)

	case "PUT":
		var bk book.Book
		json.NewDecoder(r.Body).Decode(&bk)
		bk.ID = foundBook.bookData.ID
		b.Update(foundBook.dbIndex, bk)
		json.NewEncoder(w).Encode(bk)

	case "DELETE":
		b.Remove(foundBook.dbIndex)
		json.NewEncoder(w).Encode(b)

	default:
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"message": "Request not accepted"}`))

	}
}

func HandlerBooks(b *book.Books, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(b)

	case "POST":
		var temp book.Book
		json.NewDecoder(r.Body).Decode(&temp)
		b.Add(temp)
		json.NewEncoder(w).Encode(b)

	default:
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"message": "Request not accepted"}`))
	}
}
