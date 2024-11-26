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
		ivent_name, 
		information, 
		organization, 
		poster_url, 
		preview_url, 
		skill_direction, 
		address, 
		time, 
		members_count, 
		estimated_work_hours
	) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
	RETURNING id;`

	var id int

	err := iv.db.QueryRow(
		c,
		userQuery,
		ivent.Name,
		ivent.Information,
		ivent.Organization,
		ivent.PosterUrl,
		ivent.PreviewUrl,
		ivent.SkillsDirection,
		ivent.Address,
		ivent.Time,
		ivent.CountOfPeople,
		ivent.EstimatedWorkHours,
	).Scan(&id)
	if err != nil {
		return nil, err
	}

	ivent1.ID = id
	ivent1.Name = ivent.Name
	ivent1.Information = ivent.Information
	ivent1.Organization = ivent.Organization
	ivent1.PosterUrl = ivent.PosterUrl
	ivent1.PreviewUrl = ivent.PreviewUrl
	ivent1.SkillsDirection = ivent.SkillsDirection
	ivent1.Address = ivent.Address
	ivent1.Time = ivent.Time
	ivent1.CountOfPeople = ivent.CountOfPeople
	ivent1.EstimatedWorkHours = ivent.EstimatedWorkHours

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
    	skill_direction = $6,
		address = $7,
		time = $8,
		people_count = $9,
		estimated_work_hours = $11
	WHERE id = $12;
	`

	_, err := iv.db.Exec(
		c,
		query,
		ivent.Name,
		ivent.Information,
		ivent.Organization,
		ivent.PosterUrl,
		ivent.PreviewUrl,
		ivent.SkillsDirection,
		ivent.Address,
		ivent.Time,
		ivent.CountOfPeople,
		ivent.EstimatedWorkHours,
		ivent.ID)
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
			&ivent.Information,
			&ivent.Organization,
			&ivent.PosterUrl,
			&ivent.PreviewUrl,
			&ivent.SkillsDirection,
			&ivent.Address,
			&ivent.Time,
			&ivent.CountOfPeople,
			&ivent.HowManyPeopleAccepted,
			&ivent.EstimatedWorkHours,
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
	SELECT 
		id, 
		ivent_name, 
		information, 
		organization, 
		poster_url,
		preview_url, 
		skill_direction,
		address,
		time,
		people_count,
		members_count,
		estimated_work_hours 
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
		&ivent.Address,
		&ivent.Time,
		&ivent.CountOfPeople,
		&ivent.HowManyPeopleAccepted,
		&ivent.EstimatedWorkHours,
	)

	if err != nil {
		return nil, err
	}

	return &ivent, nil
}
