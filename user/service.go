package user

type ServiceInterface interface {
	Register(user User) error
	Login(input UserInput) (User, error)
	ForgotPassword(inputNim UserUri, input UserInput) error
}

type service struct {
	repository RepositoryInterface
}

func NewService(repository RepositoryInterface) *service {
	return &service{repository}
}

func (s *service) Register(user User) error {
	err := s.repository.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Login(input UserInput) (User, error) {
	user, err := s.repository.FindByNim(input.Nim)
	if err != nil {
		return user, err
	}

	return user, nil
}
func (s *service) ForgotPassword(inputNim UserUri, input UserInput) error {
	err := s.repository.UpdatePassword(inputNim.Nim, input.Password)
	if err != nil {
		return err
	}

	return nil
}
