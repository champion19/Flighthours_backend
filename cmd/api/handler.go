package api

import (
	"fmt"
	"html/template"
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
		if email == "" {
			h.HandleError(c, domain.ErrEmployeeCannotFound)
			return
		}

		employee, err := h.service.GetEmployeeByEmail(email)
		if err != nil {
			h.HandleError(c, err)
			return
		}

		response := EmployeeResponse{
			ID:                   employee.ID,
			Name:                 employee.Name,
			Email:                employee.Email,
			Airline:              employee.Airline,
			Emailconfirmed:       employee.Emailconfirmed,
			IdentificationNumber: employee.IdentificationNumber,
			Bp:                   employee.Bp,
			StartDate:            employee.StartDate,
			EndDate:              employee.EndDate,
			Active:               employee.Active,
		}

		c.JSON(http.StatusOK, response)
	}
}

func (h handler) GetEmployeeByID() func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		if id == "" {
			h.HandleError(c, domain.ErrEmployeeCannotFound)
			return
		}

		employee, err := h.service.GetEmployeeByID(id)
		if err != nil {
			h.HandleError(c, err)
			return
		}

		response := EmployeeResponse{
			ID:                   employee.ID,
			Name:                 employee.Name,
			Email:                employee.Email,
			Airline:              employee.Airline,
			Emailconfirmed:       employee.Emailconfirmed,
			IdentificationNumber: employee.IdentificationNumber,
			Bp:                   employee.Bp,
			StartDate:            employee.StartDate,
			EndDate:              employee.EndDate,
			Active:               employee.Active,
		}

		c.JSON(http.StatusOK, response)
	}
}

func (h handler) Save() func(c *gin.Context) {
	return func(c *gin.Context) {
		var employeeRequest EmployeeRequest
		err := c.ShouldBindJSON(&employeeRequest)
		if err != nil {
			h.HandleError(c, ErrUnmarshalBody)
			return
		}

		err = employeeRequest.Validate()
		if err != nil {
			h.HandleError(c, ErrValidationUser)
			return
		}

		domainEmployee, err := employeeRequest.ToDomain()
		if err != nil {
				h.HandleError(c, err)
				return
		}

		employee, err := h.service.Save(domainEmployee)
		if err != nil {
				h.HandleError(c, err)
				return
		}
		
		response := EmployeeResponse{
			ID:                   employee.ID,
			Name:                 employee.Name,
			Email:                employee.Email,
			Airline:              employee.Airline,
			Emailconfirmed:       employee.Emailconfirmed,
			IdentificationNumber: employee.IdentificationNumber,
			Bp:                   employee.Bp,
			StartDate:            employee.StartDate,
			EndDate:              employee.EndDate,
			Active:               employee.Active,
		}

		c.JSON(http.StatusCreated, gin.H{
			"message":  "Employee created successfully. Please check your email to verify your account.",
			"employee": response,
		})
	}
}

func (h handler) VerifyEmail() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.Query("token")
		if token == "" {
			h.HandleError(c, domain.ErrInvalidToken)
			return
		}

		err := h.service.VerifyEmail(token)
		if err != nil {

			h.HandlerErrorVerification(c, err)
			return
		}
		data := ResponseEmail{
			Title:   "Email verificado con exito!",
			Content: template.HTML(" <p>Gracias por verificar tu correo.</p>"),
		}
		c.HTML(http.StatusOK, "response.html", data)
	}
}

func (h handler) EmailStatus() func(c *gin.Context) {
	return func(c *gin.Context) {
		email := c.Query("email")
		if email == "" {
			h.HandleError(c, domain.ErrEmailNotVerified)
			return
		}
	  employee,err:= h.service.GetEmployeeByEmail(email)
		if err !=nil{
			h.HandleError(c, err)
			return
		}
		response:= EmployeeEmailConfirmedResponse{
			Emailconfirmed:employee.Emailconfirmed,
		}
		c.JSON(http.StatusOK, response)
		}

}


func (h handler) HandlerErrorVerification(c *gin.Context, err error) {
	switch err {
	case domain.ErrEmailAlreadyVerified:
		h.renderErrorpage(c, "tu correo ya ha sido verificado", "puedes iniciar sesion")

	case domain.ErrTokenExpired:
		h.renderErrorpage(c, "tu token ha expirado", "puedes volver a enviar el correo de verificacion")

	case domain.ErrTokenNotFound:
		h.renderErrorpage(c, "tu token no fue encontrado", "enlace de verificacion no valido")

	case domain.ErrTokenMalformed:
		h.renderErrorpage(c, "token malformado", "el formato del token es incorrecto")

	case domain.ErrTokenUnexpected:
		h.renderErrorpage(c, "token inválido", "método de firma no esperado")

	case domain.ErrTokenInvalid:
		h.renderErrorpage(c, "token inválido", "el token no es válido o ha sido corrompido")

	case domain.ErrInvalidToken:
		h.renderErrorpage(c, "token no válido", "enlace de verificacion no valido")

	default:
		h.HandleError(c, ErrValidationUser)
	}
}

func (h handler) renderErrorpage(c *gin.Context, title string, content string) {
	data := ResponseEmail{
		Title:   title,
		Content: template.HTML(fmt.Sprintf(" <p>%s</p>", content)),
	}

	c.HTML(http.StatusOK, "response.html", data)

}

func (h handler) Login() func(c *gin.Context) {
	return func(c *gin.Context) {
		var loginRequest LoginRequest
		err := c.ShouldBindJSON(&loginRequest)
		if err != nil {
			h.HandleError(c, ErrUnmarshalBody)
			return
		}

		err = loginRequest.Validate()
		if err != nil {
			h.HandleError(c, ErrValidationUser)
			return
		}

		employee, token, err := h.service.Login(loginRequest.ToDomain())
		if err != nil {
			switch err {
			case domain.ErrEmailNotVerified:
				c.JSON(http.StatusUnauthorized, gin.H{
					"error":   "email_not_verified",
					"message": "Please verify your email before logging in.",
				})
			default:
				c.JSON(http.StatusUnauthorized, gin.H{
					"error":   "invalid_credentials",
					"message": "Invalid email or password.",
				})
			}
			return
		}

		response := LoginResponse{
			ID:                   employee.ID,
			Name:                 employee.Name,
			Email:                employee.Email,
			Airline:              employee.Airline,
			Emailconfirmed:       employee.Emailconfirmed,
			IdentificationNumber: employee.IdentificationNumber,
			Bp:                   employee.Bp,
			StartDate:            employee.StartDate,
			EndDate:              employee.EndDate,
			Active:               employee.Active,
			Token:                token,
		}

		c.JSON(http.StatusOK, response)
	}
}

func (h handler) GetAllAirlines() func(c *gin.Context) {
	return func(c *gin.Context) {
		airlines, err := h.service.GetAllAirlines()
		if err != nil {
			h.HandleError(c, domain.ErrGetAirline)
			return
		}

		var response []AirlineResponse
		for _, a := range airlines {
			response = append(response, AirlineResponse{
				ID:   a.ID,
				Name: a.Name,
			})
		}

		c.JSON(http.StatusOK, response)
	}
}
