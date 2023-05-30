package apiv1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/karsteneugene/absen-backend/models"
	"github.com/karsteneugene/absen-backend/settings"
)

func GetUsers(c *gin.Context) {
	var users []models.User

	settings.DB.Find(&users)

	c.IndentedJSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetUserByID(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	var user models.User
	settings.DB.First(&user, id)

	if user.ID != 0 {
		c.IndentedJSON(http.StatusOK, gin.H{
			"user": user,
		})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"Error": "User not found",
		})
	}
}

func PostUser(c *gin.Context) {
	var user models.User

	c.Bind(&user)

	if user.Username == "" || user.Password == "" || user.FirstName == "" || user.LastName == "" {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{
			"Error": "Empty field(s)",
		})
	} else {
		user.Password, _ = HashPassword(user.Password)
		settings.DB.Create(&user)
		c.IndentedJSON(http.StatusCreated, gin.H{
			"user": user,
		})
	}
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	var user models.User

	settings.DB.First(&user, id)

	if user.FirstName == "" || user.LastName == "" {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{
			"Error": "Empty field(s)",
		})
	} else {
		if user.ID == 0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"Error": "User not found",
			})
		} else {
			var newUser models.User
			c.Bind(&newUser)

			newUser.Password, _ = HashPassword(newUser.Password)

			result := models.User{
				ID:        user.ID,
				Username:  newUser.Username,
				Password:  newUser.Password,
				FirstName: newUser.FirstName,
				LastName:  newUser.LastName,
			}

			settings.DB.Save(&result)

			c.IndentedJSON(http.StatusOK, gin.H{
				"user": result,
			})
		}
	}
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	var user models.User

	settings.DB.First(&user, id)

	if user.ID != 0 {
		settings.DB.Delete(&user)
		c.IndentedJSON(http.StatusOK, gin.H{
			"Success": "User deleted",
		})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"Error": "User not found",
		})
	}
}
