package handler

import (
	"FinalProjectGolangH8/domain"
	"FinalProjectGolangH8/photo"
	"FinalProjectGolangH8/user"
	"FinalProjectGolangH8/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type photoHandler struct {
	photoService photo.Service
	userService  user.Service
}

func NewPhotoHandler(photoService photo.Service, userService user.Service) *photoHandler {
	return &photoHandler{
		photoService: photoService,
		userService:  userService,
	}
}

// Create godoc
// @Summary Create Photo.
// @Description Create barerToken.
// @Tags Photos
// @Param Body body domain.InputPhotos true "the body to Create a new Photo"
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 201 {object} domain.PhotoOutputCreate
// @Router /photos [post]
func (h *photoHandler) CreatePhoto(c *gin.Context) {
	var input domain.InputPhotos
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := utils.FormatVlidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
	}
	user, err := h.userService.GetUserByID(c.MustGet("currentUser").(domain.User).ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	photo, err := h.photoService.CreatePhoto(input, user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusCreated, domain.FormatPhotoOutputCreate(photo))
}

// Get All  godoc
// @Summary Get All photo.
// @Description Create barerToken.
// @Tags Photos
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []domain.PhotoOutputGet
// @Router /photos [get]
func (h *photoHandler) GetAll(c *gin.Context) {
	photos, err := h.photoService.GetAll()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	var photosuser []domain.PhotoOutputGet
	for _, v := range photos {
		photosuser = append(photosuser, domain.FormatPhotoOutputGet(v))
	}
	c.JSON(http.StatusOK, photosuser)
}

// UpdatePhoto godoc
// @Summary Update Photo.
// @Description Update Photo by id.
// @Tags Photos
// @Param Body body domain.InputPhotos true "the body to update a new Photo"
// @Produce json
// @Param id path string true "Photo id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} domain.Photo
// @Router /photos/{id} [put]
func (h *photoHandler) PutPhoto(c *gin.Context) {
	var input domain.InputPhotos
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := utils.FormatVlidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
	}
	intVar, _ := strconv.Atoi(c.Param("id"))
	photo, err := h.photoService.PutPhoto(input, intVar)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, domain.FormatPhotoOutputCreate(photo))
}

// DeletePhoto godoc
// @Summary Delete one Photo.
// @Description Delete a Photo by id.
// @Tags Photos
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path string true "Photo id"
// @Success 200 {object} domain.MessagePhoto
// @Router /photos/{id} [delete]
func (h *photoHandler) DeletePhoto(c *gin.Context) {
	intVar, _ := strconv.Atoi(c.Param("id"))
	_, err := h.photoService.DeletePhoto(intVar)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, domain.MessagePhoto{Message: "Your Photo Has Been Successfully Deleted"})

}
