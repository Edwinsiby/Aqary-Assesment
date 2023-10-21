package handlers

import (
	usecase "db/Usecase"
	"db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	err error
)

func AddStudentData(c *gin.Context) {
	userID := c.GetString("userId")
	if userID == "" {
		userID = "defaultUserID"
	}

	var input models.StudentData

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error binding json": err.Error()})
	}

	err = usecase.AddStudentData(userID, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student data added succesfully"})
}
