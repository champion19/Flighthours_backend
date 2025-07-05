package employee

type Repository interface {
	Save(employee Employee) error
	GetEmployeeByID(id string) (*Employee, error)
	GetEmployeeByEmail(email string) (*Employee, error)
}

type Service interface {
	Save(employee Employee) (Employee, error)
	GetEmployeeByID(id string) (*Employee, error)
	GetEmployeeByEmail(email string) (*Employee, error)
}
type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s service) GetEmployeeByID(id string) (*Employee, error) {
	return s.repository.GetEmployeeByID(id)
}

func (s service) GetEmployeeByEmail(email string) (*Employee, error) {
	return s.repository.GetEmployeeByEmail(email)
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
