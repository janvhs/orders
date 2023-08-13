package order

type Service interface {
	ListAll() []Order
}

// Implementation
// ------------------------------------------------------------------------

var _ Service = (*service)(nil)

type service struct {
	repo Repo
}

func NewService(repo Repo) *service {
	return &service{
		repo: repo,
	}
}

// Public Methods
// ------------------------------------------------------------------------

func (s *service) ListAll() []Order {
	return s.repo.GetAll()
}
