package models

import (
	"math/rand"
	"strconv"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Books []*Book

func CreateMockData() Books {
	var b Books

	booknames := [4]string{"A Walk To Remember", "Pride and Prejudice", "Anna Karenina", "To Kill a Mocking Bird"}
	authorsFirst := [4]string{"Nicholas", "Jane", "Leo", "Harper"}
	authorsLast := [4]string{"Sparks", "Austen", "Tolstoy", "Lee"}

	for i := 0; i < 4; i++ {
		randomISBN := rand.Intn(899999) + 100000 // for 6 random six digit integer
		b = append(
			b,
			&Book{
				ID:    strconv.Itoa(i + 1),
				Isbn:  strconv.Itoa(randomISBN),
				Title: booknames[i],
				Author: &Author{
					Firstname: authorsFirst[i],
					Lastname:  authorsLast[i],
				},
			},
		)
	}

	return b
}

func (b *Books) Add(book Book) {
	*b = append(*b, &book)
}

func (b *Books) Remove(index int) {
	btemp := *b
	btemp = append(btemp[:index], btemp[index+1:]...)
	*b = btemp
}

func (bs *Books) Update(index int, b Book) {
	bs.Remove(index)
	bs.Add(b)
}
