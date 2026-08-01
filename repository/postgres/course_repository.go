package repository

import (
	"context"

	"github.com/ArjunDev17/course-content-service/internal/database"
	"github.com/ArjunDev17/course-content-service/domain"
)

type CourseRepository interface {
	Create(
		ctx context.Context,
		course *model.Course,
	) (*model.Course, error)
}

type courseRepository struct {
	db *database.PostgreSQL
}

func NewCourseRepository(
	db *database.PostgreSQL,
) CourseRepository {

	return &courseRepository{
		db: db,
	}
}

func (r *courseRepository) Create(
	ctx context.Context,
	course *model.Course,
) (*model.Course, error) {

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

	err := r.db.Pool.QueryRow(
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

	if err != nil {
		return nil, err
	}

	return course, nil
}