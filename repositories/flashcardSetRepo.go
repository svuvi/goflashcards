package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/svuvi/goflashcards/models"
)

type FlashcardSetRepo struct {
	db *sqlx.DB
}

func NewFlashcardSetRepo(db *sqlx.DB) *FlashcardSetRepo {
	return &FlashcardSetRepo{
		db: db,
	}
}

func (r *FlashcardSetRepo) Create(set models.FlashcardSet) (int, error) {
	query := `INSERT INTO flashcard_sets (slug, title, description, edit_token) VALUES (?, ?, ?, ?)`
	result, err := r.db.Exec(query, set.Slug, set.Title, set.Description, set.EditToken)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastInsertID), nil
}

func (r *FlashcardSetRepo) Get(id int) (models.FlashcardSet, error) {
	var set models.FlashcardSet
	err := r.db.Get(&set, "SELECT * FROM flashcard_sets WHERE id = ?", id)
	return set, err
}

func (r *FlashcardSetRepo) List() ([]models.FlashcardSet, error) {
	var sets []models.FlashcardSet
	err := r.db.Select(&sets, "SELECT * FROM flashcard_sets")
	return sets, err
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

func (r *FlashcardSetRepo) Update(set models.FlashcardSet) error {
	query := `UPDATE flashcard_sets SET slug = ?, title = ?, description = ?, edit_token = ? WHERE id = ?`
	_, err := r.db.Exec(query, set.Slug, set.Title, set.Description, set.EditToken, set.ID)
	return err
}

func (r *FlashcardSetRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM flashcard_sets WHERE id = ?", id)
	return err
}
