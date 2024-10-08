package routes

import (
	"fmt"
	"net/http"
	"rest-api/models"
	"rest-api/utils"

	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse json data."})
		return
	}
	_, err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not save user: %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"massage": "user register successfuly"})
}
func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse json data."})
		return
	}
	id, err := user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	token, err := utils.GernateToken(user.Email, *id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("error: %v", err.Error())})
		return
	}
	c.JSON(http.StatusOK, gin.H{"massage": "login successfuly", "token": token})

}
