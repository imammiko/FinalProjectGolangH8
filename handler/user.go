package handler

import (
	"FinalProjectGolangH8/auth"
	"FinalProjectGolangH8/user"
	"FinalProjectGolangH8/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{
		userService: userService,
		authService: authService,
	}
}

// Register godoc
// @Summary Register a user.
// @Description registering a user from public access.
// @Tags Auth
// @Param Body body user.RegisterUserInput true "the body to register a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users/register [post]
func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := utils.FormatVlidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusCreated, user.FormatUser(newUser))
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := utils.FormatVlidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
	}
	user, err := h.userService.Login(input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	token, err := h.authService.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	responseSuccess := gin.H{"token": token}
	c.JSON(http.StatusOK, responseSuccess)
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	var input user.UpdateUserInput
	input.ID = c.MustGet("currentUser").(user.User).ID
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := utils.FormatVlidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
	}
	user, err := h.userService.UpdateUser(input)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	var idUser int
	idUser = c.MustGet("currentUser").(user.User).ID
	_, err := h.userService.DeleteUser(idUser)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Your account has been successfully deleted"})
}
