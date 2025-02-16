package models

import "time"

type FlashcardSet struct {
	ID          int       `db:"id"`
	CreatedAt   time.Time `db:"created_at"`
	Slug        string    `db:"slug"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	EditToken   string    `db:"edit_token"`
}

type FlashcardSetRepository interface {
	Create(FlashcardSet) (int, error)
	Get(id int) (FlashcardSet, error)
	List() ([]FlashcardSet, error)
	Update(FlashcardSet) error
	Delete(id int) error
}
