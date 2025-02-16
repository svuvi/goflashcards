package models

type Card struct {
	ID         int    `db:"id"`
	Term       string `db:"term"`
	Definition string `db:"definition"`
	SetID      int    `db:"set_id"`
}

type CardRepository interface {
	Create(Card) (Card, error)
	List(setID int) ([]Card, error)
	Update(Card) error
	Delete(id int) error
}
