package repo

import (
	"UDVtest/config"
	"UDVtest/handlers"
	"UDVtest/repo/migration"
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const pg = "postgres"

type Store struct {
	cfg config.Database
	db  *gorm.DB
}

func (s *Store) GetBookList(page, size int) (handlers.BookResponseList, error) {

	var books handlers.BookResponseList
	var book handlers.BookResponse
	rows, err := s.db.
		Model(&handlers.BookResponse{}).
		Table("books").
		Select("book_id, title, description, authors, year, edition, " +
			"publishings.name as publishing_name").
		Joins("INNER JOIN publishings ON books.publishing_id = publishings.publishing_id").
		Offset(page).
		Limit(size).
		Rows()
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			println(err)
		}
	}(rows)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&book.BookID, &book.Title, &book.Description, &book.Authors, &book.Year, &book.Edition, &book.PublishingName)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil

}

func (s *Store) GetHistory(id int) (handlers.HistoryResponseList, error) {

	var historyList handlers.HistoryResponseList
	var hist handlers.HistoryResponse
	rows, err := s.db.
		Model(&handlers.HistoryResponse{}).
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
		Where("libraries.lib_book_id = ?", id).
		Rows()

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(rows)
	if err != nil {
		return historyList, err
	}
	for rows.Next() {
		err := rows.Scan(&hist.LibBookID, &hist.BookID, &hist.Title, &hist.Name, &hist.DateIssue, &hist.DateReturn)
		if err != nil {
			return nil, err
		}
		historyList = append(historyList, hist)
	}
	return historyList, nil
}

func (s *Store) GetBook(id int) (handlers.BookResponse, error) {
	var book handlers.BookResponse
	if err := s.db.Table("books").
		Select("book_id, title, description, authors, year, edition, "+
			"publishings.name as publishing_name, "+
			"(Select Count(*) FROM libraries Where book_id= ?) AS amount", id).
		Joins("INNER JOIN publishings ON books.publishing_id = publishings.publishing_id").
		Where("books.book_id = ?", id).
		First(&book).Error; err != nil {
		return book, err
	}
	return book, nil
}

func (s *Store) GetItem(id int) (handlers.ItemResponse, error) {
	var item handlers.ItemResponse
	if err := s.db.
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
			"ORDER BY usages.usage_id DESC", id).
		First(&item).
		Error; err != nil {
		return item, err
	}
	return item, nil
}

func New(cfg config.Database) *Store {
	return &Store{
		cfg: cfg,
	}
}

func (s *Store) InitDB() {

	db, err := gorm.Open(pg, s.cfg.DSN())
	if err != nil {
		fmt.Print(err)
	}

	db.Debug().AutoMigrate(&Book{}, &Publishing{}, &Library{}, &Usage{}, &User{})

	s.db = db

}

func (s *Store) Migrate() {
	migration.Run(s.db.DB())
}
