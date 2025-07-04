package employee

type Repository interface {
	Save(employee Employee) error
}

type Service interface {
	Save(employee Employee) (Employee, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s service) Save(employee Employee) (Employee, error) {
	employee.setID()
	employee.hashPassword()
	err := s.repository.Save(employee)
	if err != nil {
		return Employee{}, err
	}

	return employee, nil
}
