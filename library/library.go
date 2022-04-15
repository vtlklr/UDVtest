package library

type Library struct {
	LibBookID uint   `gorm:"primary_key"`
	BookID    uint   `json:"book_id"`
	Placement string `json:"placement"`
}
