package controllers

import (
	"net/http"
	"pbi-final-task-go-api/database"
	"pbi-final-task-go-api/models"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func AddPhoto(c *gin.Context) {
	var body struct {
		Title    string
		Caption  string
		PhotoURL string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to read body"})
		return
	}

	if !govalidator.IsURL(body.PhotoURL) {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid URL"})
		return
	}

	currentUser, _ := c.Get("user")

	newPhoto := models.Photo{Title: body.Title, Caption: body.Caption, PhotoURL: body.PhotoURL, UserID: currentUser.(models.User).ID}
	result := database.DB.Create(&newPhoto)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to add photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": "Photo added successfully"})
}

func GetPhotos(c *gin.Context) {
	var photos []models.Photo
	type getPhotoStruct struct {
		ID       uint
		Title    string
		Caption  string
		PhotoURL string
		UserID   uint
	}

	var getPhotos []getPhotoStruct

	currentUser, _ := c.Get("user")

	result := database.DB.Model(&photos).Find(&getPhotos, "user_id = ?", currentUser.(models.User).ID)

	if c.Query("title") != "" {
		result = database.DB.Model(&photos).Raw("SELECT * FROM photos WHERE title ILIKE ?", "%"+c.Query("title")+"%").Scan(&getPhotos)
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to fetch photos"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": "Photos fetched successfully", "photos": getPhotos})
}

func EditPhoto(c *gin.Context) {
	var body struct {
		Title    string
		Caption  string
		PhotoURL string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Failed to read body"})
		return
	}

	if !govalidator.IsURL(body.PhotoURL) {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid URL"})
		return
	}

	updateID, _ := strconv.ParseUint(c.Param("photoId"), 10, 64)

	var photo models.Photo
	database.DB.First(&photo, "id = ?", updateID)

	if photo.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Photo with specified ID not found"})
		return
	}

	currentUser, _ := c.Get("user")

	if photo.UserID != currentUser.(models.User).ID {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Cannot edit photo of other users"})
		return
	}

	result := database.DB.Model(&photo).Updates(map[string]interface{}{"title": body.Title, "caption": body.Caption, "photo_url": body.PhotoURL})

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to update photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": "Photo editted successfully"})
}

func DeletePhoto(c *gin.Context) {
	deleteID, _ := strconv.ParseUint(c.Param("photoId"), 10, 64)

	var photo models.Photo
	database.DB.First(&photo, "id = ?", deleteID)

	if photo.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Photo with specified ID not found"})
		return
	}

	currentUser, _ := c.Get("user")

	if photo.UserID != currentUser.(models.User).ID {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Cannot delete photo of other users"})
		return
	}

	result := database.DB.Delete(&photo)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Failed to delete photo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success": "Photo deleted successfully"})
}
