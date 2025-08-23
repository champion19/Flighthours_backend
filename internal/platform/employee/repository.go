package employee

import (
	"database/sql"
	"log"
	"strings"

	domain "github.com/champion19/Flighthours_backend/internal/domain/employee"
)

const (
	querySave                 = "Insert into employee(id,name,airline,email,password,email_confirmed,identification_number,bp,start_date,end_date,active) values(?,?,?,?,?,?,?,?,?,?,?)"
	queryGetByID              = "Select id,name,airline,email,password,email_confirmed,identification_number,bp,start_date,end_date,active from employee where id=?"
	QueryByEmail              = "Select id,name,airline,email,password,email_confirmed,identification_number,bp,start_date,end_date,active from employee where email=?"
	queryUpdateEmailConfirmed = "UPDATE employee SET email_confirmed = ? WHERE id = ?"
	queryGetAllAirline        = "Select id,airline_name,airline_code from airline where status='active'"
)

type repository struct {
	db *sql.DB
}

func (r *repository) GetEmployeeByEmail(email string) (*domain.Employee, error) {
	stmt, err := r.db.Prepare(QueryByEmail)
	if err != nil {
		return nil, domain.ErrGettingEmployeeByEmail
	}
	defer func () {
		if err := stmt.Close(); err != nil {
			log.Printf("failed to close statement: %v", err)
		}
	}()

	var employee Employee
	err = stmt.QueryRow(email).Scan(&employee.ID, &employee.Name, &employee.Airline, &employee.Email, &employee.Password, &employee.Emailconfirmed, &employee.IdentificationNumber, &employee.Bp, &employee.StartDate, &employee.EndDate, &employee.Active)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFoundEmployeeByEmail
		}
		return nil, domain.ErrGettingEmployeeByEmail
	}

	employeeDomain := employee.ToDomain()
	return &employeeDomain, nil
}

func (r *repository) GetEmployeeByID(id string) (*domain.Employee, error) {
	stmt, err := r.db.Prepare(queryGetByID)
	if err != nil {
		return nil, domain.ErrGetEmployee
	}
	defer func () {
		if err := stmt.Close(); err != nil {
			log.Printf("failed to close statement: %v", err)
		}
	}()

	var employee Employee
	err = stmt.QueryRow(id).Scan(
		&employee.ID,
		&employee.Name,
		&employee.Airline,
		&employee.Email,
		&employee.Password,
		&employee.Emailconfirmed,
		&employee.IdentificationNumber,
		&employee.Bp,
		&employee.StartDate,
		&employee.EndDate,
		&employee.Active,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrEmployeeCannotFound
		}
		return nil, domain.ErrEmployeeCannotGet
	}
	employeeDomain := employee.ToDomain()
	return &employeeDomain, nil
}

func NewRepository(db *sql.DB) domain.Repository {
	return &repository{
		db: db,
	}
}
func (r *repository) Save(employee domain.Employee) error {
	employeeToSave := Employee{
		ID:                   employee.ID,
		Name:                 employee.Name,
		Airline:              employee.Airline,
		Email:                employee.Email,
		Password:             employee.Password,
		Emailconfirmed:       employee.Emailconfirmed,
		IdentificationNumber: employee.IdentificationNumber,
		Bp:                   employee.Bp,
		StartDate:            employee.StartDate,
		EndDate:              employee.EndDate,
		Active:               employee.Active,
	}
	stmt, err := r.db.Prepare(querySave)
	if err != nil {
		return domain.ErrEmployeeCannotSave
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			log.Printf("failed to close statement: %v", err)
		}
	}()

	_, err = stmt.Exec(
		employeeToSave.ID,
		employeeToSave.Name,
		employeeToSave.Airline,
		employeeToSave.Email,
		employeeToSave.Password,
		employeeToSave.Emailconfirmed,
		employeeToSave.IdentificationNumber,
		employeeToSave.Bp,
		employeeToSave.StartDate,
		employeeToSave.EndDate,
		employeeToSave.Active,
	)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "Duplicate"):
			return domain.ErrDuplicateEmployee
		default:
			return domain.ErrEmployeeCannotSave
		}
	}

	return nil

}

func (r *repository) UpdateEmailConfirmed(id string, confirmed bool) error {
	stmt, err := r.db.Prepare(queryUpdateEmailConfirmed)
	if err != nil {
		return domain.ErrSavingEmployee
	}
	defer func() {
		if cerr := stmt.Close(); cerr != nil {
			log.Printf("failed to close statement: %v", cerr)
		}
	}()

	result, err := stmt.Exec(confirmed, id)
	if err != nil {
		return domain.ErrSavingEmployee
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return domain.ErrSavingEmployee
	}

	if rowsAffected == 0 {
		return domain.ErrEmployeeCannotFound
	}

	return nil
}

func (r *repository) GetAllAirlines() ([]domain.Airline, error) {
	stmt, err := r.db.Prepare(queryGetAllAirline)
	if err != nil {
		log.Printf("Error al preparar la consulta de aerolíneas: %v", err)
		return nil, domain.ErrAirlineCannotFound
	}
	defer func() {
    if cerr := stmt.Close(); cerr != nil {
        log.Printf("Error al cerrar statement de aerolíneas: %v", cerr)
    }
}()

	rows, err := stmt.Query()
	if err != nil {
		log.Printf("Error al ejecutar la consulta de aerolíneas: %v", err)
		return nil, domain.ErrAirlineCannotFound
	}
	defer func() {
		if cerr := rows.Close(); cerr != nil {
			log.Printf("Error al cerrar filas de aerolíneas: %v", cerr)
		}
	}()

	var airlines []domain.Airline

	for rows.Next() {
		var airline domain.Airline
		if err := rows.Scan(&airline.ID, &airline.Name, &airline.Code); err != nil {
			log.Printf("Error al escanear fila de aerolínea: %v", err)
			return nil, domain.ErrAirlineCannotFound
		}
		log.Printf("Aerolínea encontrada: ID=%s, Name=%s, Code=%s", airline.ID, airline.Name, airline.Code)
		airlines = append(airlines, airline)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error al iterar sobre las filas: %v", err)
		return nil, domain.ErrAirlineCannotFound
	}

	log.Printf("Total de aerolíneas encontradas: %d", len(airlines))
	return airlines, nil
}
