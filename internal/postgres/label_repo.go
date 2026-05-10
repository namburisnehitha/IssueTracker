package postgres

import (
	"context"
	"database/sql"

	"github.com/namburisnehitha/IssueTracker/domain"
)

type PostgresLabelRepository struct {
	db *sql.DB
}

func NewPostgresLabelRepository(db *sql.DB) *PostgresLabelRepository {
	return &PostgresLabelRepository{
		db: db,
	}
}

func (lr *PostgresLabelRepository) Save(ctx context.Context, label domain.Label) error {
	query := `INSERT into labels(id,label_name,label_description,colour) Values($1,$2,$3,$4)`
	_, err := lr.db.Exec(query, label.Id, label.Name, label.Description, label.Colour)
	return err
}

func (lr *PostgresLabelRepository) GetById(ctx context.Context, id string) (domain.Label, error) {
	var label domain.Label
	query := `SELECT id,label_name,label_description,colour FROM labels WHERE id = $1 `
	err := lr.db.QueryRow(query, id).Scan(&label.Id, &label.Name, &label.Description, &label.Colour)
	return label, err
}

func (lr *PostgresLabelRepository) GetByName(ctx context.Context, name string) (domain.Label, error) {
	var label domain.Label
	query := `SELECT id,label_name,label_description,colour FROM labels WHERE label_name= $1 `
	err := lr.db.QueryRow(query, name).Scan(&label.Id, &label.Name, &label.Description, &label.Colour)
	return label, err
}

func (lr *PostgresLabelRepository) GetByColour(ctx context.Context, colour string) ([]domain.Label, error) {
	var labels []domain.Label
	query := `SELECT id,label_name,label_description,colour FROM labels WHERE colour = $1 `
	rows, err := lr.db.Query(query, colour)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var l domain.Label
		err = rows.Scan(&l.Id, &l.Name, &l.Description, &l.Colour)
		if err != nil {
			return nil, err
		}
		labels = append(labels, l)
	}
	return labels, err
}

func (lr *PostgresLabelRepository) UpdateLabel(ctx context.Context, label domain.Label) error {
	query := `UPDATE labels SET label_name = $1,label_description=$2,colour = $3 WHERE id = $4`
	_, err := lr.db.Exec(query, label.Name, label.Description, label.Colour, label.Id)
	return err
}

func (lr *PostgresLabelRepository) DeleteLabel(ctx context.Context, label domain.Label) error {
	query := `DELETE FROM labels WHERE id = $1 `
	_, err := lr.db.Exec(query, label.Id)
	return err
}

func (lr *PostgresLabelRepository) LabelList(ctx context.Context) ([]domain.Label, error) {
	var labels []domain.Label
	query := `SELECT id,label_name,label_description,colour FROM labels `
	rows, err := lr.db.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var l domain.Label
		err = rows.Scan(&l.Id, &l.Name, &l.Description, &l.Colour)
		if err != nil {
			return nil, err
		}
		labels = append(labels, l)
	}
	return labels, err
}
