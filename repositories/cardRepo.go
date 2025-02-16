package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/svuvi/goflashcards/models"
)

type CardRepo struct {
	db *sqlx.DB
}

func NewCardRepo(db *sqlx.DB) *CardRepo {
	return &CardRepo{
		db: db,
	}
}

func (r *CardRepo) Create(card models.Card) (models.Card, error) {
	query := `INSERT INTO cards (term, definition, set_id) VALUES (?, ?, ?)`
	result, err := r.db.Exec(query, card.Term, card.Definition, card.SetID)
	if err != nil {
		return models.Card{}, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return models.Card{}, err
	}

	card.ID = int(lastInsertID)
	return card, nil
}

func (r *CardRepo) List(setID int) ([]models.Card, error) {
	var cards []models.Card
	err := r.db.Select(&cards, "SELECT * FROM cards WHERE set_id = ?", setID)
	return cards, err
}

func (r *CardRepo) CountCardsInSet(setID int) (int, error) {
	var count int
	err := r.db.Get(&count, "SELECT COUNT(*) FROM cards WHERE set_id = ?", setID)
	return count, err
}

func (r *CardRepo) GetNthCard(setID, n int) (models.Card, error) {
	var card models.Card
	err := r.db.Get(&card, `
		SELECT * FROM cards 
		WHERE set_id = ?
		LIMIT 1 OFFSET ?`,
		setID, n-1)
	return card, err
}

func (r *CardRepo) Update(card models.Card) error {
	query := `UPDATE cards SET term = ?, definition = ? WHERE id = ?`
	_, err := r.db.Exec(query, card.Term, card.Definition, card.ID)
	return err
}

func (r *CardRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM cards WHERE id = ?", id)
	return err
}
