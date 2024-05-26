package postgres

import (
	"context"
	"fmt"
	"to_do_app/api/models"
	"to_do_app/pkg/helper"
	"to_do_app/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type labelRepo struct {
	DB *pgxpool.Pool
}

func NewLabelRepo(db *pgxpool.Pool) storage.ILabelStorage {
	return &labelRepo{
		DB: db,
	}
}

func (l *labelRepo) Create(ctx context.Context, createLabel models.CreateLabel) (models.Label, error) {
	label := models.Label{}
	uid := uuid.New()

	query := `INSERT INTO labels (id, name, color, user_id)
			values ($1, $2, $3, $4)
		returning id, name, color, user_id, created_at::text `

	err := l.DB.QueryRow(ctx, query, 
		uid,
		createLabel.Name,
		createLabel.Color,
		createLabel.UserID,
	).Scan(
		&label.ID,
		&label.Name,
		&label.Color,
		&label.UserID,
		&label.CreatedAt,
	)
	if err != nil{
		fmt.Println("error while creating label!", err.Error())
		return models.Label{},err
	}

	return label, nil
}

func (l *labelRepo) Get(ctx context.Context, pKey models.PrimaryKey) (models.Label, error) {
	label := models.Label{}

	query := `SELECT id, name, color, user_id, created_at::text, updated_at::text 
			FROM labels where id = $1 `

	err := l.DB.QueryRow(ctx,query, pKey.ID).Scan(
		&label.ID,
		&label.Name,
		&label.Color,
		&label.UserID,
		&label.CreatedAt,
		&label.UpdatedAt,
	)

	if err != nil{
		fmt.Println("error while getting label!", err.Error())
		return models.Label{}, err
	}

	return label, nil
}

func (l *labelRepo) Update(ctx context.Context, updLabel models.UpdateLabel) (models.UpdateResponse, error) {
	var (
		params = make(map[string]interface{})
		filter = ""
		query = `UPDATE labels set `
		label = models.UpdateResponse{}
	)

	params["id"] = updLabel.ID

	if updLabel.Name != ""{
		params["name"] = updLabel.Name
		filter += " name = @name, "
	}

	if updLabel.Color != ""{
		params["color"] = updLabel.Color
		filter += " color = @color, "
	}

	query += filter + ` updated_at = now() WHERE id = @id 
		returning id, updated_at::text `
	
	fullQuery, args := helper.ReplaceQueryParams(query, params)

	err := l.DB.QueryRow(ctx, fullQuery, args...).Scan(
		&label.ID,
		&label.UpdatedAt,
	)
	fmt.Println(fullQuery)
	if err != nil{
		fmt.Println("error while updating labels!", err.Error())
		return models.UpdateResponse{}, err
	}

	return label, nil
}

func (l *labelRepo) Delete(ctx context.Context, pKey models.PrimaryKey) (error) {
	query := `DELETE FROM labels where id = $1 `

	_, err := l.DB.Exec(ctx, query, pKey.ID)
	if err != nil{
		fmt.Println("error while deleting label!", err.Error())
		return err
	}

	return nil
}

func (l *labelRepo) DeleteByUserID(ctx context.Context, pKey models.PrimaryKey) (error) {
	query := `DELETE FROM labels where user_id = $1 `

	_, err := l.DB.Exec(ctx, query, pKey.ID)
	if err != nil{
		fmt.Println("error while deleting labels by user id!", err.Error())
		return err
	}

	return nil
}