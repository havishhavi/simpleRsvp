package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"www.rsvpme.com/handlers"
	"www.rsvpme.com/initializers"
	"www.rsvpme.com/models"
)

func getTotalPersons() (int, error) {
	var total int
	result := initializers.Db.Model(&models.User{}).Select("SUM(persons)").Row()
	err := result.Scan(&total)
	return total, err
}

func CreateRSVP(c *gin.Context) {
	var rsvp models.User
	if err := c.ShouldBindJSON(&rsvp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if initializers.Db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not initialized"})
		return
	}

	if err := initializers.Db.Create(&rsvp).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create RSVP: " + err.Error()})
		return
	}

	if err := handlers.SendEmail(rsvp.Email, "RSVP Confirmation", "Thank you for your RSVP \n Address: 10112 Tidwell Street, Krugerville, Texas, 76227"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email: " + err.Error()})
		return
	}

	totalPersons, err := getTotalPersons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate total persons: " + err.Error()})
		return
	}

	admEmail := os.Getenv("ADMIN_EMAIL")
	adminEmailBody := fmt.Sprintf("New RSVP Details: \n Name: %s\nEmail: %s\nPersons: %d\n\nTotal persons confirmed so far: %d", rsvp.Name, rsvp.Email, rsvp.Persons, totalPersons)
	if err := handlers.SendEmail(admEmail, "New RSVP Received", adminEmailBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send admin email: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, rsvp)
}
