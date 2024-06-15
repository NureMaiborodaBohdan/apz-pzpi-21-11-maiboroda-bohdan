package handler

import (
	"apz-pzpi-21-11-maiboroda-bohdan-task2/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handlers {
	return &Handlers{service: services}
}

func (h *Handlers) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.register)
		auth.POST("/registeradmin", h.registerAdmin)
		auth.POST("/login", h.login)
	}

	api := router.Group("/api", h.userIndetity)
	{
		user := api.Group("/user/:userID")
		{
			user.GET("/testresults", h.getUserTestResults)
			user.POST("/testresults", h.createUserTestResult)
			user.GET("/notifications", h.getUserNotifications)
			user.GET("/accesscontrols", h.getUserAccessControls)
		}

		admin := api.Group("/admin", h.adminRequired)
		{
			admin.GET("/companies", h.getCompanies)
			admin.GET("/companies/:companyID", h.getCompanyByID)
			admin.POST("/companies", h.createCompany)
			admin.PUT("/companies/:companyID", h.updateCompany)
			admin.DELETE("/companies/:companyID", h.deleteCompany)

			admin.GET("/locations", h.getLocations)
			admin.GET("/locations/:locationID", h.getLocationByID)
			admin.POST("/locations", h.createLocation)
			admin.PUT("/locations/:locationID", h.updateLocation)
			admin.DELETE("/locations/:locationID", h.deleteLocation)

			users := admin.Group("/users")
			{
				users.GET("/", h.getUsers)
				users.POST("/", h.createUser)
				users.GET("/:userID", h.getUserByID)
				users.PUT("/:userID", h.updateUser)
				users.DELETE("/:userID", h.deleteUser)
			}

			admin.GET("/testresults", h.getTestResult)
			admin.GET("/testresults/:testID", h.getTestResultByID)
			admin.DELETE("/testresults/:testID", h.deleteTestResult)

			admin.GET("/notifications", h.getAllNotification)
			admin.GET("/notifications/:notificationID", h.getNotificationByID)
			admin.POST("/notifications", h.createNotification)
			admin.DELETE("/notifications/:notificationID", h.deleteNotification)

			data := admin.Group("/data")
			{
				data.POST("/backup", h.backupData)
				data.POST("/restore", h.restoreData)
				data.GET("/export", h.exportData)
				data.POST("/import", h.importData)
			}
		}
	}
	return router
}
