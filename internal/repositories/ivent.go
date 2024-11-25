package repository

import (
	"context"

	"github.com/Orken1119/Register/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type IventRepository struct {
	db *pgxpool.Pool
}

func NewIventRepository(db *pgxpool.Pool) models.IventRepository {
	return &IventRepository{db: db}
}

func (iv *IventRepository) CreateIvent(c context.Context, ivent *models.MainIvent) (*models.Ivent, error) {
	var ivent1 models.Ivent
	userQuery := `
	INSERT INTO ivents (
		ivent_name, information, organization, poster_url, preview_url, skill_direction
	) 
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id;`

	var id int

	err := iv.db.QueryRow(c, userQuery, ivent.Name, ivent.Information, ivent.Organization, ivent.PosterUrl, ivent.PreviewUrl, ivent.SkillsDirection).Scan(&id)
	if err != nil {
		return nil, err
	}

	ivent1.ID = id
	return &ivent1, nil
}

func (iv *IventRepository) DeleteIvent(c context.Context, id int) error {
	query := `
	DELETE FROM ivents where id = $1;
	`

	_, err := iv.db.Exec(c, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (iv *IventRepository) UpdateIvent(c context.Context, ivent *models.Ivent) error {
	query := `
		UPDATE ivents
	SET ivent_name = $1,
    	information = $2,
    	organization = $3,
    	poster_url = $4,
    	preview_url = $5,
    	skill_direction = $6
	WHERE id = $7;
	`

	_, err := iv.db.Exec(c, query, ivent.Name, ivent.Information, ivent.Organization, ivent.PosterUrl, ivent.PreviewUrl, ivent.SkillsDirection, ivent.ID)
	if err != nil {
		return err
	}

	return nil
}

func (iv *IventRepository) GetAllIvent(c context.Context) (*[]models.Ivent, error) {
	var ivents []models.Ivent

	query := `
	SELECT * FROM ivents`

	rows, err := iv.db.Query(c, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var ivent models.Ivent
		if err = rows.Scan(
			&ivent.ID,
			&ivent.Name,
			&ivent.Organization,
			&ivent.PosterUrl,
			&ivent.PreviewUrl,
			&ivent.SkillsDirection,
		); err != nil {
			return nil, err
		}
		ivents = append(ivents, ivent)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &ivents, nil

}

func (iv *IventRepository) GetIventById(c context.Context, id int) (*models.Ivent, error) {
	query := `
	SELECT id, ivent_name, information, organization, poster_url, preview_url, skill_direction 
	FROM ivents 
	WHERE id = $1`

	var ivent models.Ivent

	err := iv.db.QueryRow(c, query, id).Scan(
		&ivent.ID,
		&ivent.Name,
		&ivent.Information,
		&ivent.Organization,
		&ivent.PosterUrl,
		&ivent.PreviewUrl,
		&ivent.SkillsDirection,
	)

	if err != nil {
		return nil, err
	}

	return &ivent, nil
}
