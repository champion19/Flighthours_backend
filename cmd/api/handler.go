package api

import (
	"net/http"

	domain "github.com/champion19/Flighthours_backend/internal/domain/employee"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service domain.Service
}

func New(service domain.Service) *handler {
	return &handler{
		service: service,
	}
}
func (h handler) GetEmployeeByEmail() func(c *gin.Context) {
	return func(c *gin.Context) {
		email := c.Param("email")

		employee, err := h.service.GetEmployeeByEmail(email) // Usa una función del servicio para obtener el usuario por su email
		if err != nil {
			h.HandleError(c, err)
			return
		}

		c.JSON(http.StatusOK, employee)
	}
}

func (h handler) GetEmployeeByID() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")

		employee, err := h.service.GetEmployeeByID(id)
		if err != nil {
			h.HandleError(c, err)
			return
		}

		c.JSON(http.StatusOK, employee)
	}
}

func (h handler) Save() func(c *gin.Context) {
	return func(c *gin.Context) {
		var employeeRequest EmployeeRequest
		err := c.BindJSON(&employeeRequest)
		if err != nil {
			h.HandleError(c, ErrUnmarshalBody)
			return
		}

		err = employeeRequest.Validate()
		if err != nil {
			h.HandleError(c, ErrValidationUser)
			return
		}

		employee,err := h.service.Save(employeeRequest.ToDomain())
		if err != nil {
			h.HandleError(c, err)
			return
		}

		response := EmployeeResponse{
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
		c.JSON(http.StatusOK, response)
	}
}
