package course

import (
	"context"

	"github.com/ArjunDev17/course-content-service/domain"
)

type Repository interface {

	Create(
		ctx context.Context,
		course *domain.Course,
	) error
}