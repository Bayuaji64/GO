package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data.", "error": err.Error()})
		return
	}
	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user", "error": err.Error()})
		return
	}

	// context.JSON(http.StatusCreated, gin.H{"message": "User created succesfully", "user_detail": user})

	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user_detail": gin.H{
			"id":       user.ID,
			"email":    user.Email,
			"password": user.Password, // Catatan: Mengembalikan password hashed
		},
	})

}

func login(context *gin.Context) {

}
