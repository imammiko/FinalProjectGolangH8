package handler

import (
	"FinalProjectGolangH8/auth"
	"FinalProjectGolangH8/domain"
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
// @Tags Users
// @Param Body body domain.RegisterUserInput true "the body to register a user"
// @Produce json
// @Success 201 {object} domain.UserFormatter
// @Router /users/register [post]
func (h *userHandler) RegisterUser(c *gin.Context) {
	var input domain.RegisterUserInput
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
	c.JSON(http.StatusCreated, domain.FormatUser(newUser))
}

// Login godoc
// @Summary Login a user.
// @Description Login a user from public access.
// @Tags Users
// @Param Body body domain.LoginUserInput true "the body to Login a user"
// @Produce json
// @Success 201 {object} map[string]string{}
// @Router /users/login [post]
func (h *userHandler) Login(c *gin.Context) {
	var input domain.LoginUserInput
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

// Update godoc
// @Summary Update user.
// @Description Update barerToken.
// @Tags Users
// @Param Body body domain.UpdateUserInput true "the body to update a new users"
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} domain.UserUpdateFormatter
// @Router /users [put]
func (h *userHandler) UpdateUser(c *gin.Context) {
	var input domain.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := utils.FormatVlidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
	}
	userUpdate, err := h.userService.UpdateUser(input, c.MustGet("currentUser").(domain.User).ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	response := domain.FormatUserUpdateFormatter(userUpdate)
	c.JSON(http.StatusOK, response)
}

// DeleteUser godoc
// @Summary Delete one User.
// @Description Delete a Product by Token id.
// @Tags Users
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /users [delete]
func (h *userHandler) DeleteUser(c *gin.Context) {
	var idUser int
	idUser = c.MustGet("currentUser").(domain.User).ID
	_, err := h.userService.DeleteUser(idUser)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Your account has been successfully deleted"})
}
