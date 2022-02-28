package handler

import (
	"FinalProjectGolangH8/domain"
	socialmedia "FinalProjectGolangH8/socialMedia"
	"FinalProjectGolangH8/user"
	"FinalProjectGolangH8/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SocialMediaHandler struct {
	socialmediaService socialmedia.Service
	userService        user.Service
}

func NewSocialMediaHandler(socialMediaService socialmedia.Service, userService user.Service) *SocialMediaHandler {
	return &SocialMediaHandler{
		socialmediaService: socialMediaService,
		userService:        userService,
	}
}

// Create SocialMedia
// @Summary Create SocialMedia.
// @Description Create barerToken.
// @Tags SocialMedias
// @Param Body body domain.InputSocialMedia true "the body to Create a new SocialMedia"
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 201 {object} domain.SocialMediaOutputCreate
// @Router /socialmedias [post]
func (h *SocialMediaHandler) CreateSocialMedia(c *gin.Context) {
	var input domain.InputSocialMedia
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
	socialMedia, err := h.socialmediaService.CreateSocialMedia(input, user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusCreated, domain.FormatterSocialMediaOutputCreate(socialMedia))
}

// Get All  godoc
// @Summary Get All SocialMedia.
// @Description Create barerToken.
// @Tags SocialMedias
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []domain.SocialMediaOutputGet
// @Router /socialmedias [get]
func (h *SocialMediaHandler) GetAll(c *gin.Context) {
	socialMedia, photo, err := h.socialmediaService.GetAll()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	var response []domain.SocialMediaOutputGet
	for _, v := range socialMedia {
		var indexphoto int
		for index, photo := range photo {
			if v.UserId == photo.User_id {
				indexphoto = index
			}
		}
		response = append(response, domain.FormatterSocialMediaOutputGet(v, photo[indexphoto]))
	}
	c.JSON(http.StatusOK, response)
}

// UpdateSocialMedia godoc
// @Summary Update SocialMedias.
// @Description Update Photo by id.
// @Tags SocialMedias
// @Param Body body domain.InputSocialMedia true "the body to update a new Photo"
// @Produce json
// @Param id path string true "SocialMedia id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} domain.SocialMediaOutputPut
// @Router /socialmedias/{id} [put]
func (h *SocialMediaHandler) PutSocialMedia(c *gin.Context) {
	var input domain.InputSocialMedia
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := utils.FormatVlidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
	}
	intVar, _ := strconv.Atoi(c.Param("id"))
	socialMedia, err := h.socialmediaService.PutSocialMedia(input, intVar)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, domain.FormatterSocialMediaOutputPUt(socialMedia))
}

// DeleteSocialMedia godoc
// @Summary Delete one SocialMedia.
// @Description Delete a SocialMedia by id.
// @Tags SocialMedias
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path string true "SocialMedia id"
// @Success 200 {object} domain.SocialMediaMessage
// @Router /socialmedias/{id} [delete]
func (h *SocialMediaHandler) DeleteSocialMedia(c *gin.Context) {
	intVar, _ := strconv.Atoi(c.Param("id"))
	_, err := h.socialmediaService.DeleteSocialMedia(intVar)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, domain.FormatterSocialMediaMessage("Your SocialMedia has been successfully deleted"))
}
