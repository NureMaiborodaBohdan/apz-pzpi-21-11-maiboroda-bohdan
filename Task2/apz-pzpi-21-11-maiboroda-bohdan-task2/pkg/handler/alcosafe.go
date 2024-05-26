package handler

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handlers) getCompanies(c *gin.Context) {
	companies, err := h.service.Company.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, companies)
}

func (h *Handlers) getCompanyByID(c *gin.Context) {
	companyID, err := strconv.Atoi(c.Param("companyID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid company ID"})
		return
	}
	company, err := h.service.Company.GetByID(companyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, company)
}

func (h *Handlers) createCompany(c *gin.Context) {
	var company AlcoSafe.Company
	if err := c.BindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}
	id, err := h.service.Company.Create(company)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"CompanyID": id,
	})
}

func (h *Handlers) deleteCompany(c *gin.Context) {
	companyID, err := strconv.Atoi(c.Param("companyID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid company ID"})
		return
	}
	if err := h.service.Company.Delete(companyID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Company deleted"})
}

func (h *Handlers) updateCompany(c *gin.Context) {
	c.JSON(200, gin.H{"message": "updateCompany endpoint"})
}

func (h *Handlers) getLocationByID(c *gin.Context) {
	locationID, err := strconv.Atoi(c.Param("locationID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid location ID"})
		return
	}
	location, err := h.service.Location.GetByID(locationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, location)
}

func (h *Handlers) createLocation(c *gin.Context) {
	var location AlcoSafe.Location
	if err := c.BindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}
	id, err := h.service.Location.Create(location)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handlers) deleteLocation(c *gin.Context) {
	locationID, err := strconv.Atoi(c.Param("locationID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid location ID"})
		return
	}
	if err := h.service.Location.Delete(locationID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Location deleted"})
}

func (h *Handlers) updateLocation(c *gin.Context) {
	c.JSON(200, gin.H{"message": "updateLocation endpoint"})
}

func (h *Handlers) getNotificationByID(c *gin.Context) {
	c.JSON(200, gin.H{"message": "getNotificationByID endpoint"})
}

func (h *Handlers) createNotification(c *gin.Context) {
	c.JSON(200, gin.H{"message": "createNotification endpoint"})
}

func (h *Handlers) updateNotification(c *gin.Context) {
	c.JSON(200, gin.H{"message": "updateNotification endpoint"})
}

func (h *Handlers) deleteNotification(c *gin.Context) {
	c.JSON(200, gin.H{"message": "deleteNotification endpoint"})
}

func (h *Handlers) getAccessControlByID(c *gin.Context) {
	c.JSON(200, gin.H{"message": "getAccessControlByID endpoint"})
}

func (h *Handlers) createAccessControl(c *gin.Context) {
	c.JSON(200, gin.H{"message": "createAccessControl endpoint"})
}

func (h *Handlers) updateAccessControl(c *gin.Context) {
	c.JSON(200, gin.H{"message": "updateAccessControl endpoint"})
}

func (h *Handlers) deleteAccessControl(c *gin.Context) {
	c.JSON(200, gin.H{"message": "deleteAccessControl endpoint"})
}
func (h *Handlers) getTestResult(c *gin.Context) {
	c.JSON(200, gin.H{"message": "getTestResult endpoint"})
}
func (h *Handlers) getTestResultByID(c *gin.Context) {
	c.JSON(200, gin.H{"message": "getTestResultByID endpoint"})
}

func (h *Handlers) createTestResult(c *gin.Context) {
	c.JSON(200, gin.H{"message": "createTestResult endpoint"})
}

func (h *Handlers) updateTestResult(c *gin.Context) {
	c.JSON(200, gin.H{"message": "updateTestResult endpoint"})
}

func (h *Handlers) deleteTestResult(c *gin.Context) {
	c.JSON(200, gin.H{"message": "deleteTestResult endpoint"})
}

func (h *Handlers) getUserByID(c *gin.Context) {
	c.JSON(200, gin.H{"message": "getUserByID endpoint"})
}

func (h *Handlers) getUsers(c *gin.Context) {
	c.JSON(200, gin.H{"message": "getUsers endpoint"})
}

func (h *Handlers) createUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "createUser endpoint"})
}

func (h *Handlers) updateUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "updateUser endpoint"})
}

func (h *Handlers) deleteUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "deleteUser endpoint"})
}

func (h *Handlers) getCurrentUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "getCurrentUser endpoint"})
}

func (h *Handlers) updateCurrentUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "updateCurrentUser endpoint"})
}

func (h *Handlers) getUserTestResults(c *gin.Context) {
	c.JSON(200, gin.H{"message": "getUserTestResults endpoint"})
}

func (h *Handlers) getUserTestResultByID(c *gin.Context) {
	c.JSON(200, gin.H{"message": "getUserTestResultByID endpoint"})
}

func (h *Handlers) createUserTestResult(c *gin.Context) {
	c.JSON(200, gin.H{"message": "createUserTestResult endpoint"})
}

func (h *Handlers) updateUserTestResult(c *gin.Context) {
	c.JSON(200, gin.H{"message": "updateUserTestResult endpoint"})
}

func (h *Handlers) deleteUserTestResult(c *gin.Context) {
	c.JSON(200, gin.H{"message": "deleteUserTestResult endpoint"})
}

func (h *Handlers) getUserNotifications(c *gin.Context) {
	c.JSON(200, gin.H{"message": "getUserNotifications endpoint"})
}

func (h *Handlers) getUserNotificationByID(c *gin.Context) {
	c.JSON(200, gin.H{"message": "getUserNotificationByID endpoint"})
}

func (h *Handlers) getUserAccessControls(c *gin.Context) {
	c.JSON(200, gin.H{"message": "getUserAccessControls endpoint"})
}

func (h *Handlers) getUserAccessControlByID(c *gin.Context) {
	c.JSON(200, gin.H{"message": "getUserAccessControlByID endpoint"})
}

func (h *Handlers) backupData(c *gin.Context) {
	c.JSON(200, gin.H{"message": "backupData endpoint"})
}

func (h *Handlers) restoreData(c *gin.Context) {
	c.JSON(200, gin.H{"message": "restoreData endpoint"})
}

func (h *Handlers) exportData(c *gin.Context) {
	c.JSON(200, gin.H{"message": "exportData endpoint"})
}

func (h *Handlers) importData(c *gin.Context) {
	c.JSON(200, gin.H{"message": "importData endpoint"})
}
