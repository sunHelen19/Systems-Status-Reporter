package usecase

type UseCase struct {
	repo Infrastructure
}

func New(i Infrastructure) *UseCase {
	return &UseCase{
		repo: i,
	}
}
