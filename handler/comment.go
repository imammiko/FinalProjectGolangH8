package handler

import (
	"FinalProjectGolangH8/comment"
	"FinalProjectGolangH8/domain"
	"FinalProjectGolangH8/photo"
	"FinalProjectGolangH8/user"
	"FinalProjectGolangH8/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type commentHandler struct {
	commentService comment.Service
	userService    user.Service
	photoService   photo.Service
}

func NewCommentHandler(commentService comment.Service, userService user.Service, photoService photo.Service) *commentHandler {
	return &commentHandler{
		commentService: commentService,
		userService:    userService,
		photoService:   photoService,
	}
}

// CreateComment godoc
// @Summary Create Comment.
// @Description Create barerToken.
// @Tags Comments
// @Param Body body domain.InputCommentCreate true "the body to Create a new Photo"
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 201 {object} domain.CommentCreate
// @Router /comments [post]
func (h *commentHandler) CreateComment(C *gin.Context) {
	var input domain.InputCommentCreate
	if err := C.ShouldBindJSON(&input); err != nil {
		errors := utils.FormatVlidationError(err)
		errorMessage := gin.H{"errors": errors}
		C.JSON(http.StatusUnprocessableEntity, errorMessage)
	}
	_, err := h.photoService.GetPhotoByID(input.PhotoID)
	if err != nil {
		C.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	comment, err := h.commentService.CreateComment(input, C.MustGet("currentUser").(domain.User).ID)
	if err != nil {
		C.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	C.JSON(http.StatusCreated, domain.FormatterCommentOutputCreate(comment))
}

// Get All  godoc
// @Summary Get All Comments.
// @Description Create barerToken.
// @Tags Comments
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} []domain.CommentGet
// @Router /comments [get]
func (h *commentHandler) GetAll(c *gin.Context) {
	comment, err := h.commentService.GetAll()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	var Comments []domain.CommentGet
	for _, v := range comment {
		Comments = append(Comments, domain.FormatterCommentOutputGet(v))
	}
	c.JSON(http.StatusOK, Comments)
}

// UpdatComments godoc
// @Summary Update Comments.
// @Description Update Comments by id.
// @Tags Comments
// @Param Body body domain.UpdateComment true "the body to update a new Photo"
// @Produce json
// @Param id path string true "Comment id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Success 200 {object} domain.CommentPut
// @Router /comments/{id} [put]
func (h *commentHandler) PutComment(c *gin.Context) {
	var input domain.UpdateComment
	if err := c.ShouldBindJSON(&input); err != nil {
		errors := utils.FormatVlidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
	}
	intVar, _ := strconv.Atoi(c.Param("id"))
	comment, err := h.commentService.PutPhoto(input, intVar)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, domain.FormatterCommentOutputPut(comment))

}

// DeleteComment godoc
// @Summary Delete one Comment.
// @Description Delete a Comment by id.
// @Tags Comments
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path string true "Comment id"
// @Success 200 {object} domain.MessagePhoto
// @Router /comments/{id} [delete]
func (h *commentHandler) DeleteComment(c *gin.Context) {
	intVar, _ := strconv.Atoi(c.Param("id"))
	_, err := h.commentService.DeleteComment(intVar)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, domain.MessagePhoto{Message: "Your Comment has been successfully deleted"})
}
