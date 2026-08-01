package course

import (
	"context"

	"github.com/ArjunDev17/course-content-service/domain"
)

type Repository interface {
	Create(
		ctx context.Context,
		course *model.Course,
	) (*model.Course, error)

	GetByID(
		ctx context.Context,
		id string,
	) (*model.Course, error)

	GetByTitle(
		ctx context.Context,
		title string,
	) (*model.Course, error)

	Update(
		ctx context.Context,
		course *model.Course,
	) (*model.Course, error)

	Delete(
		ctx context.Context,
		id string,
	) error

	List(
		ctx context.Context,
	) ([]*model.Course, error)
}