package main

import (
	"UDVtest/library"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type getBook struct {
	BookID         uint   `json:"book_id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Authors        string `json:"authors"`
	Year           int    `json:"year"`
	Edition        string `json:"edition"`
	PublishingName string `json:"publishing_name"`
	Amount         uint   `json:"amount"`
}

var GetBook = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var b getBook

	if err1 := library.GetDB().Table("books").
		Select("book_id, title, description, authors, year, edition, "+
			"publishings.name as publishing_name, "+
			"(Select Count(*) FROM libraries Where book_id= ?) AS amount", idInt).
		Joins("INNER JOIN publishings ON books.publishing_id = publishings.publishing_id").
		Where("books.book_id = ?", idInt).
		First(&b).Error; err1 != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err1.Error(), idInt)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(b); err != nil {
		w.WriteHeader(http.StatusForbidden)
	}

}

type getItem struct {
	LibBookID uint   `json:"lib_book_id"`
	BookID    uint   `json:"book_id"`
	Title     string `json:"title"`
	Placement string `json:"placement"`
}

var GetItems = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var itm getItem
	if err1 := library.GetDB().
		Raw("SELECT  libraries.lib_book_id,books.book_id,title, "+
			"case WHEN (date_return IS NULL AND date_issue IS NOT NULL) THEN CONCAT('Книга на на руках у: ', users.name, '. Выдана: ', date_issue) "+
			"WHEN (date_return IS NULL AND date_issue IS NULL) THEN CONCAT('Книга в библиотеке: ',placement) "+
			"WHEN (date_return IS NOT NULL) THEN CONCAT('Книга в библиотеке: ',placement) "+
			"END AS placement "+
			"FROM libraries "+
			"LEFT JOIN books USING(book_id) "+
			"FULL JOIN usages USING(lib_book_id) "+
			"LEFT JOIN users USING(user_id) "+
			"LEFT JOIN publishings USING(publishing_id) "+
			"WHERE libraries.lib_book_id = ?"+
			"ORDER BY usages.usage_id DESC", idInt).
		First(&itm).
		Error; err1 != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err1.Error(), idInt)
		return
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(itm); err != nil {
		w.WriteHeader(http.StatusForbidden)
	}

}

var GetBookList = func(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		fmt.Println("не корректно указана страница " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		fmt.Println("не корректно указан размер " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if page <= 0 {
		fmt.Println("Введите номер страницы больше 0 ")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if size <= 0 {
		fmt.Println("Введите размер страницы больше 0 ")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//var bookList getBookList
	var books []getBook
	var book getBook
	rows, err := library.GetDB().
		Model(&getBook{}).
		Table("books").
		Select("book_id, title, description, authors, year, edition, " +
			"publishings.name as publishing_name").
		Joins("INNER JOIN publishings ON books.publishing_id = publishings.publishing_id").
		Offset(page*size - size).
		Limit(size).
		Rows()
	defer rows.Close()
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println(err.Error())
		return
	}
	for rows.Next() {
		rows.Scan(&book.BookID, &book.Title, &book.Description, &book.Authors, &book.Year, &book.Edition, &book.PublishingName)
		books = append(books, book)
	}
	//fmt.Println(books)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)

}

type getHistory struct {
	LibBookID  uint   `json:"lib_book_id"`
	BookID     uint   `json:"book_id"`
	Title      string `json:"title"`
	Name       string `json:"name"`
	DateIssue  string `json:"date_issue"`
	DateReturn string `json:"date_return"`
}

var GetHistory = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var hist getHistory
	var hists []getHistory
	rows, err := library.GetDB().
		Model(&getHistory{}).
		Table("libraries").
		Select("libraries.lib_book_id, books.book_id, title, "+
			"CASE "+
			"WHEN (date_issue IS NULL) THEN 'Книга не выдавалась' "+
			"WHEN (date_issue IS NOT NULL) THEN users.name "+
			"END as name, "+
			"date_issue, date_return ").
		Joins("LEFT JOIN usages USING(lib_book_id) ").
		Joins("LEFT JOIN users USING(user_id) ").
		Joins("LEFT JOIN books USING(book_id)").
		Where("libraries.lib_book_id = ?", idInt).
		Rows()

	defer rows.Close()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err.Error(), idInt)
		return
	}
	for rows.Next() {
		rows.Scan(&hist.LibBookID, &hist.BookID, &hist.Title, &hist.Name, &hist.DateIssue, &hist.DateReturn)
		hists = append(hists, hist)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(hists)

}
