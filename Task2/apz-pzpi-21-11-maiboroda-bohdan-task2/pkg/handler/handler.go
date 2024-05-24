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
		auth.POST("/registerAdmin", h.registerAdmin)
		auth.POST("/login", h.login)
	}

	admin := router.Group("/admin", h.adminRequired)
	{
		admin.GET("/companies/:companyID", h.getCompanyByID)
		admin.POST("/companies", h.createCompany)
		admin.PUT("/companies/:companyID", h.updateCompany)
		admin.DELETE("/companies/:companyID", h.deleteCompany)

		company := admin.Group("/companies/:companyID")
		{
			company.GET("/locations/:locationID", h.getLocationByID)
			company.POST("/locations", h.createLocation)
			company.PUT("/locations/:locationID", h.updateLocation)
			company.DELETE("/locations/:locationID", h.deleteLocation)

			company.GET("/thresholds/:thresholdID", h.getThresholdByID)
			company.POST("/thresholds", h.createThreshold)
			company.PUT("/thresholds/:thresholdID", h.updateThreshold)
			company.DELETE("/thresholds/:thresholdID", h.deleteThreshold)
		}

		users := admin.Group("/users")
		{
			users.GET("/", h.getUsers)
			users.POST("/", h.createUser)
			users.GET("/:userID", h.getUserByID)
			users.PUT("/:userID", h.updateUser)
			users.DELETE("/:userID", h.deleteUser)
		}

		admin.GET("/testresults/:testID", h.getTestResultByID)
		admin.POST("/testresults", h.createTestResult)
		admin.PUT("/testresults/:testID", h.updateTestResult)
		admin.DELETE("/testresults/:testID", h.deleteTestResult)

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

	api := router.Group("/api", h.userIndetity)
	{
		user := api.Group("/user")
		{
			user.GET("/:userID/testresults", h.getUserTestResults)
			user.GET("/:userID/testresults/:testID", h.getUserTestResultByID)
			user.POST("/:userID/testresults", h.createUserTestResult)
			user.PUT("/:userID/testresults/:testID", h.updateUserTestResult)
			user.DELETE("/:userID/testresults/:testID", h.deleteUserTestResult)

			user.GET("/:userID/notifications", h.getUserNotifications)
			user.GET("/:userID/notifications/:notificationID", h.getUserNotificationByID)

			user.GET("/:userID/accesscontrols", h.getUserAccessControls)
			user.GET("/:userID/accesscontrols/:accessID", h.getUserAccessControlByID)
		}
	}

	return router
}
