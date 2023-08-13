// TODO: Currently, having a service layer feels unnecessary.
package order

type Service interface {
	ListAll() ([]Order, error)
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

func (s *service) ListAll() ([]Order, error) {
	return s.repo.GetAll()
}
