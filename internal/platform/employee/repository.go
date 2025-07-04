package employee

import (
	"database/sql"
	"strings"

	domain "github.com/champion19/Flighthours_backend/internal/domain/employee"
)
const(
	querysave="Insert into employee(name,email,password,emailconfirmed,identificationnumber,bp,fecha_inicio,fecha_fin,vigente) values(?,?,?,?,?,?,?,?,?)"
	//querygetbyid="Select id,name,email,password,emailconfirmed,identificationnumber,bp,fecha_inicio,fecha_fin,vigente from employee where id=?"
	//querygetbyemail="Select id,name,email,password,emailconfirmed,identificationnumber,bp,fecha_inicio,fecha_fin,vigente from employee where email=?"

)
type repository struct {
	db *sql.DB
}
func NewRepository(db *sql.DB) domain.Repository {
	return &repository{
		db: db,
	}
}
func (r *repository) Save(employee domain.Employee) error {
employeeToSave:=Employee{
		ID: employee.ID,
		Name: employee.Name,
		Email: employee.Email,
		Password: employee.Password,
		Emailconfirmed: employee.Emailconfirmed,
		IdentificationNumber: employee.IdentificationNumber,
		Bp: employee.Bp,
		Fecha_inicio: employee.Fecha_inicio,
		Fecha_fin: employee.Fecha_fin,
		Vigente: employee.Vigente,
	}
	stmt, err := r.db.Prepare(querysave)
	if err != nil {
		return domain.ErrUserCannotSave
	}
	defer stmt.Close()
	_,err=stmt.Exec(
		employeeToSave.ID,
		employeeToSave.Name,
		employeeToSave.Email,
		employeeToSave.Password,
		employeeToSave.Emailconfirmed,
		employeeToSave.IdentificationNumber,
		employeeToSave.Bp,
		employeeToSave.Fecha_inicio,
		employeeToSave.Fecha_fin,
		employeeToSave.Vigente,
	)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "Duplicate"):
			return domain.ErrDuplicateUser
		default:
			return domain.ErrUserCannotSave
		}
	}

	return nil

}


