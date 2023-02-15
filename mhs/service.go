package mhs

type ServiceInterface interface {
	Create(mhs Mahasiswa) error
	Update(input InputMhsDetail, mhs InputMhs, userID int) error
	Delete(input InputMhsDetail, userID int)
	FindAll() ([]Mahasiswa, error)
	FindByNim(input InputMhsDetail, userID int) (Mahasiswa, error)
}

type service struct {
	repository RepositoryInterface
}

func NewService(repository RepositoryInterface) *service {
	return &service{repository}
}

func (s *service) Create(mhs Mahasiswa) error {
	err := s.repository.Create(mhs)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Update(input InputMhsDetail, mhs InputMhs, userID int) error {

	data, err := s.repository.FindByNim(input.Nim, userID)
	if err != nil {
		return err
	}

	if mhs.Nama == "" {
		mhs.Nama = data.Nama
	}
	if mhs.Email == "" {
		mhs.Email = data.Email
	}
	if mhs.Alamat == "" {
		mhs.Alamat = data.Alamat
	}

	err = s.repository.Update(mhs, input.Nim, userID)
	if err != nil {
		return err
	}

	return err
}

func (s *service) Delete(input InputMhsDetail, userID int) {
	s.repository.Delete(input.Nim, userID)
}

func (s *service) FindAll() ([]Mahasiswa, error) {
	data, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *service) FindByNim(input InputMhsDetail, userID int) (Mahasiswa, error) {
	data, err := s.repository.FindByNim(input.Nim, userID)
	if err != nil {
		return data, err
	}

	return data, err
}
