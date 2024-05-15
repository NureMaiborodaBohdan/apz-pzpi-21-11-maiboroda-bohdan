package handler

import "github.com/gin-gonic/gin"

type Handlers struct {
}

func (h *Handlers) InitRoutes() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/register", h.register)
		auth.POST("/registerAdmin", h.registerAdmin)
		auth.POST("/login", h.login)
	}

	api := router.Group("/api")
	{
		companies := api.Group("/companies")
		{
			companies.GET("/:companyID", h.getCompanyByID)
			companies.POST("/", h.createCompany)
			companies.PUT("/:companyID", h.updateCompany)
			companies.DELETE("/:companyID", h.deleteCompany)

			companies.GET("/:companyID/locations/:locationID", h.getLocationByID)
			companies.POST("/:companyID/locations", h.createLocation)
			companies.PUT("/:companyID/locations/:locationID", h.updateLocation)
			companies.DELETE("/:companyID/locations/:locationID", h.deleteLocation)

			companies.GET("/:companyID/thresholds/:thresholdID", h.getThresholdByID)
			companies.POST("/:companyID/thresholds", h.createThreshold)
			companies.PUT("/:companyID/thresholds/:thresholdID", h.updateThreshold)
			companies.DELETE("/:companyID/thresholds/:thresholdID", h.deleteThreshold)

			users := companies.Group("/:companyID/users")
			{
				users.GET("/:userID/testresults/:testID", h.getTestResultByID)
				users.POST("/:userID/testresults", h.createTestResult)
				users.PUT("/:userID/testresults/:testID", h.updateTestResult)
				users.DELETE("/:userID/testresults/:testID", h.deleteTestResult)

				users.GET("/:userID/accesscontrols/:accessID", h.getAccessControlByID)
				users.POST("/:userID/accesscontrols", h.createAccessControl)
				users.PUT("/:userID/accesscontrols/:accessID", h.updateAccessControl)
				users.DELETE("/:userID/accesscontrols/:accessID", h.deleteAccessControl)

				users.GET("/:userID/notifications/:notificationID", h.getNotificationByID)
				users.POST("/:userID/notifications", h.createNotification)
				users.PUT("/:userID/notifications/:notificationID", h.updateNotification)
				users.DELETE("/:userID/notifications/:notificationID", h.deleteNotification)
			}
		}
	}

	return router
}
