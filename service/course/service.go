package course

type Service struct {
	repository Repository
	publisher  EventPublisher
}

func New(
	repository Repository,
	publisher EventPublisher,
) *Service {

	return &Service{
		repository: repository,
		publisher:  publisher,
	}
}