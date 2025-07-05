package api

import (
	domain "github.com/champion19/Flighthours_backend/internal/domain/employee"

	"github.com/gin-gonic/gin"
)




func routing(app *gin.Engine, dependencies *Dependencies) {
	employeeService := domain.NewService(dependencies.employee)
	handler := New(employeeService)

	app.POST("/v1/employee", handler.Save())
}

func Boostrap(app *gin.Engine) {
	dependencies := initDependencies()
	routing(app, dependencies)
}
