package library

type Usage struct {
	UsageID    uint   `gorm:"primary_key"`
	LibBookID  uint   `json:"lib_book_id"`
	UserID     uint   `json:"user_id"`
	DateIssue  string `json:"date_issue"`
	DateReturn string `json:"date_return"`
}
