package postgres

import (
	"context"
	"database/sql"

	"github.com/namburisnehitha/IssueTracker/domain"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
)

type PostgresLabelRepository struct {
	db     *sql.DB
	tracer trace.Tracer
}

func NewPostgresLabelRepository(db *sql.DB) *PostgresLabelRepository {
	return &PostgresLabelRepository{
		db:     db,
		tracer: otel.Tracer("postgres-label-repo"),
	}
}

func (lr *PostgresLabelRepository) Save(ctx context.Context, label domain.Label) error {

	query := `INSERT into labels(id,label_name,label_description,colour) Values($1,$2,$3,$4)`

	ctx, span := lr.tracer.Start(ctx, "CreateLabel")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	_, err := lr.db.ExecContext(ctx, query, label.Id, label.Name, label.Description, label.Colour)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	return err
}

func (lr *PostgresLabelRepository) GetById(ctx context.Context, id string) (domain.Label, error) {

	var label domain.Label
	query := `SELECT id,label_name,label_description,colour FROM labels WHERE id = $1 `

	ctx, span := lr.tracer.Start(ctx, "GetById")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	err := lr.db.QueryRowContext(ctx, query, id).Scan(&label.Id, &label.Name, &label.Description, &label.Colour)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return domain.Label{}, err
	}
	return label, err
}

func (lr *PostgresLabelRepository) GetByName(ctx context.Context, name string) (domain.Label, error) {

	var label domain.Label
	query := `SELECT id,label_name,label_description,colour FROM labels WHERE label_name= $1 `

	ctx, span := lr.tracer.Start(ctx, "GetByName")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	err := lr.db.QueryRowContext(ctx, query, name).Scan(&label.Id, &label.Name, &label.Description, &label.Colour)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return domain.Label{}, err
	}
	return label, err
}

func (lr *PostgresLabelRepository) GetByColour(ctx context.Context, colour string) ([]domain.Label, error) {

	var labels []domain.Label
	query := `SELECT id,label_name,label_description,colour FROM labels WHERE colour = $1 `

	ctx, span := lr.tracer.Start(ctx, "GetBycolour")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	rows, err := lr.db.QueryContext(ctx, query, colour)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var l domain.Label
		err = rows.Scan(&l.Id, &l.Name, &l.Description, &l.Colour)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}
		labels = append(labels, l)
	}
	if err := rows.Err(); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	return labels, err
}

func (lr *PostgresLabelRepository) UpdateLabel(ctx context.Context, label domain.Label) error {

	query := `UPDATE labels SET label_name = $1,label_description=$2,colour = $3 WHERE id = $4`

	ctx, span := lr.tracer.Start(ctx, "UpdateLabel")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	_, err := lr.db.ExecContext(ctx, query, label.Name, label.Description, label.Colour, label.Id)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	return err
}

func (lr *PostgresLabelRepository) DeleteLabel(ctx context.Context, label domain.Label) error {

	query := `DELETE FROM labels WHERE id = $1 `

	ctx, span := lr.tracer.Start(ctx, "DeleteLabel")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	_, err := lr.db.ExecContext(ctx, query, label.Id)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return err
	}
	return err
}

func (lr *PostgresLabelRepository) LabelList(ctx context.Context) ([]domain.Label, error) {

	var labels []domain.Label
	query := `SELECT id,label_name,label_description,colour FROM labels `

	ctx, span := lr.tracer.Start(ctx, "LabelList")
	span.SetAttributes(semconv.DBQueryTextKey.String(query))
	defer span.End()

	rows, err := lr.db.QueryContext(ctx, query)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var l domain.Label
		err = rows.Scan(&l.Id, &l.Name, &l.Description, &l.Colour)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, err
		}
		labels = append(labels, l)
	}
	if err := rows.Err(); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
	return labels, err
}
