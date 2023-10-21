package handlers

import (
	"custom/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CustomBinding(c *gin.Context) {
	var input models.CustomInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error binding json": err.Error()})
	}

	output := &models.Output{
		UserName:   "Given user name :" + input.UserName,
		NickName:   "Given nick name :" + *input.NickName,
		Age:        "Given age : " + strconv.Itoa(input.Age),
		Handicaped: "Is user Handicaped:" + strconv.FormatBool(input.Handicaped),
		Tag:        "Given user tag :" + string(input.Tag),
	}

	c.JSON(http.StatusAccepted, gin.H{"User Data": output})

}
