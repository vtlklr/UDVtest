package service

import (
	"UDVtest/handlers"
)

type BookService struct {
	bookStorage BookStore
}
type BookStore interface {
	GetBookList(page, size int) (handlers.BookResponseList, error)
	GetHistory(id int) (handlers.HistoryResponseList, error)
	GetBook(id int) (handlers.BookResponse, error)
	GetItem(id int) (handlers.ItemResponse, error)
}

func New(bookStore BookStore) *BookService {
	return &BookService{bookStorage: bookStore}
}

func (b *BookService) GetItem(id int) (handlers.ItemResponse, error) {

	return b.bookStorage.GetItem(id)
}
func (b *BookService) GetHistory(id int) (handlers.HistoryResponseList, error) {
	return b.bookStorage.GetHistory(id)
}

func (b *BookService) GetBook(id int) (handlers.BookResponse, error) {
	return b.bookStorage.GetBook(id)

}
func (b *BookService) GetBookList(page, size int) (handlers.BookResponseList, error) {
	return b.bookStorage.GetBookList(page*size-size, size)

}
