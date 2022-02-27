package middleware

import (
	"FinalProjectGolangH8/auth"
	"FinalProjectGolangH8/comment"
	"FinalProjectGolangH8/domain"
	"FinalProjectGolangH8/photo"
	socialmedia "FinalProjectGolangH8/socialMedia"
	"FinalProjectGolangH8/user"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token not found"})
			return
		}
		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unathorized"})
			return
		}
		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unathorized"})
			return
		}
		userID := int(claim["user_id"].(float64))
		user, err := userService.GetUserByID(userID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unathorized"})
			return
		}
		c.Set("currentUser", user)
	}
}

func AuthzMiddleware(photoService photo.Service, commentService comment.Service, socialMediaService socialmedia.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.MustGet("currentUser").(domain.User).ID
		intVar, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Id must be number"})
			return
		}
		urlPath := c.Request.URL.Path
		if strings.Contains(urlPath, "photos") {
			photo, err := photoService.GetPhotoByID(intVar)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Photo not found"})
				return
			}
			if photo.User_id != userId {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
		} else if strings.Contains(urlPath, "comments") {
			comment, err := commentService.GetCommentByID(intVar)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Comment not found"})
				return
			}
			if comment.User_id != userId {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
		} else if strings.Contains(urlPath, "socialmedia") {
			socialMedia, err := socialMediaService.GetSocialMediaByID(intVar)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Social Media not found"})
				return
			}

			if socialMedia.UserId != userId {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
				return
			}
		}
	}
}
