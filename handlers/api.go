package handlers

import (
	"UDVtest/config"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type LibraryApi struct {
	cfg config.Web
	svc LibraryService
}

type LibraryService interface {
	GetBookList(page, size int) (BookResponseList, error)
	GetBook(id int) (BookResponse, error)
	GetItem(id int) (ItemResponse, error)
	GetHistory(id int) (HistoryResponseList, error)
}

type HistoryResponse struct {
	LibBookID  uint   `json:"lib_book_id"`
	BookID     uint   `json:"book_id"`
	Title      string `json:"title"`
	Name       string `json:"name"`
	DateIssue  string `json:"date_issue"`
	DateReturn string `json:"date_return"`
}

type HistoryResponseList []HistoryResponse

type ItemResponse struct {
	LibBookID uint   `json:"lib_book_id"`
	BookID    uint   `json:"book_id"`
	Title     string `json:"title"`
	Placement string `json:"placement"`
}

type BookResponse struct {
	BookID         uint   `json:"book_id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Authors        string `json:"authors"`
	Year           uint   `json:"year"`
	Edition        string `json:"edition"`
	PublishingName string `json:"publishing_name"`
	Amount         uint   `json:"amount"`
}

type BookResponseList []BookResponse

func (l *LibraryApi) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("bad request")
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("bad request ", err)

		return
	}

	book, err := l.svc.GetBook(idInt)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("not found ", err)

		return
	}
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println("error encode ", err)

		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Println("get book done")

}

func (l *LibraryApi) GetBookList(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		fmt.Println("error bad page ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	size, err := strconv.Atoi(r.URL.Query().Get("size"))
	if err != nil {
		fmt.Println("error bad size ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if page <= 0 {
		fmt.Println("error bad page")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if size <= 0 {
		fmt.Println("error bad size")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	books, err := l.svc.GetBookList(page, size)
	if err != nil {
		fmt.Println("could not get data ", err)

		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		fmt.Println("cannot encode ", err)
		return
	}
	fmt.Println("get book list done")
}

func (l *LibraryApi) GetItems(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("wrong id")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("wrong id")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	item, err := l.svc.GetItem(idInt)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("could not get data ", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(item)
	if err != nil {
		fmt.Println("cannot encode", err)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	fmt.Println("get item done")

}

func (l *LibraryApi) GetHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		fmt.Println("wrong id")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println("wrong id")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	HistoryList, err := l.svc.GetHistory(idInt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("cannot get history ", err)
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(HistoryList)
	if err != nil {
		fmt.Println("cannot encode ", err)
		return
	}
	fmt.Println("get history done")
}

func New(cfg config.Web, svc LibraryService) *LibraryApi {
	return &LibraryApi{
		cfg: cfg,
		svc: svc,
	}
}

func (l *LibraryApi) Start() {
	router := SetupRoute(l)

	err := http.ListenAndServe(l.cfg.Port(), router)
	if err != nil {
		fmt.Print(err)
	}

}
