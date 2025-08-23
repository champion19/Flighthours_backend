package api

import (
	"log"
	"path/filepath"

	"github.com/champion19/Flighthours_backend/tools/utils"
	"github.com/gin-gonic/gin"
)

func routing(app *gin.Engine, dependencies *Dependencies) {
	handler := New(dependencies.employeeService)
	moduleRoot, err := utils.FindModuleRoot()
	if err != nil {
		log.Fatalf("Failed to find module root: %v", err)
	}
	templatePath := filepath.Join(moduleRoot, "cmd", "api", "template", "*.html")
	log.Printf("Template path: %s", templatePath)
	app.LoadHTMLGlob(templatePath)

	app.POST("/v1/employees", handler.Save())
	app.GET("/v1/employees/verify-email", handler.VerifyEmail())
	app.GET("/v1/airlines", handler.GetAllAirlines())
	app.GET("/v1/login", handler.Login())
	app.GET("/v1/Flighthours/email/status", handler.EmailStatus())
}

func Boostrap(app *gin.Engine) {
	dependencies := initDependencies()
	routing(app, dependencies)
}
