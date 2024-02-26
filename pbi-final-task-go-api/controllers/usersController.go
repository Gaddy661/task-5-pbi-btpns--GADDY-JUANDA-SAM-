package controllers

import (
	"net/http"
	"pbi-final-task-go-api/app"
	"pbi-final-task-go-api/database"
	"pbi-final-task-go-api/helpers"
	"pbi-final-task-go-api/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var body struct {
		Username string
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to read body"})
		return
	}

	if !app.ValidateUserInput(body.Username, body.Email, body.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid input"})
		return
	}

	hash, err := helpers.GenerateFromPassword(body.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to hash password"})
		return
	}

	newUser := models.User{Username: body.Username, Email: strings.ToLower(body.Email), Password: string(hash)}
	result := database.DB.Create(&newUser)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": "User created successfully"})
}

func Login(c *gin.Context) {

	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to read body"})
		return
	}

	if !app.ValidateUserLogin(body.Email, body.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid email or password"})
		return
	}

	var user models.User
	database.DB.First(&user, "email = ?", strings.ToLower(body.Email))

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid email or password"})
		return
	}

	err := helpers.CompareHashAndPassword(user.Password, body.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid email or password"})
		return
	}

	tokenString, err := helpers.CreateTokenString(user.ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to create token"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"Success": "User logged in successfully"})
}

func Update(c *gin.Context) {

	var body struct {
		Username string
		Email    string
		Password string
	}

	updateID, _ := strconv.ParseUint(c.Param("userId"), 10, 64)

	currentUser, _ := c.Get("user")

	if updateID != uint64(currentUser.(models.User).ID) {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Forbidden to update other users"})
		return
	}

	var user models.User
	database.DB.First(&user, "id = ?", updateID)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "User with specified ID not found"})
		return
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to read body"})
		return
	}

	if !app.ValidateUserInput(body.Username, body.Email, body.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid input"})
		return
	}

	hash, err := helpers.GenerateFromPassword(body.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to hash password"})
		return
	}

	result := database.DB.Model(&user).Updates(map[string]interface{}{"username": body.Username, "email": strings.ToLower(body.Email), "password": string(hash)})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": "User updated successfully"})
}

func Delete(c *gin.Context) {

	deleteID, _ := strconv.ParseUint(c.Param("userId"), 10, 64)

	currentUser, _ := c.Get("user")

	if deleteID != uint64(currentUser.(models.User).ID) {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Forbidden to delete other users"})
		return
	}

	var user models.User
	database.DB.First(&user, "id = ?", deleteID)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "User with specified ID not found"})
		return
	}

	result := database.DB.Delete(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": "User deleted successfully"})
}
