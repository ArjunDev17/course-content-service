package postgres

import (
	"context"

	"github.com/ArjunDev17/course-content-service/internal/database"
	"github.com/ArjunDev17/course-content-service/domain"
	courseservice "github.com/ArjunDev17/course-content-service/service/course"
)

type repository struct {
	db *database.PostgreSQL
}

func New(
	db *database.PostgreSQL,
) courseservice.Repository {

	return &repository{
		db: db,
	}
}

func (r *repository) Create(
	ctx context.Context,
	course *domain.Course,
) error {

	query := `
	INSERT INTO courses
	(
		title,
		description,
		category,
		instructor
	)
	VALUES
	(
		$1,
		$2,
		$3,
		$4
	)
	RETURNING
		id,
		created_at,
		updated_at;
	`

	return r.db.Pool.QueryRow(
		ctx,
		query,
		course.Title,
		course.Description,
		course.Category,
		course.Instructor,
	).Scan(
		&course.ID,
		&course.CreatedAt,
		&course.UpdatedAt,
	)
}