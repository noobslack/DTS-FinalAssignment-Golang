package middleware

import (
	"mygram-finalprojectdts/helper"
	"mygram-finalprojectdts/model"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthMiddleware(ctx *gin.Context) {
	authorizationValue := ctx.GetHeader("Authorization")

	splittedValue := strings.Split(authorizationValue, "Bearer ")

	if len(splittedValue) <= 1 {
		var r model.Response = model.Response{
			Success: false,
			Error:   "unauthorized",
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, r)
		return
	}

	jwtToken := splittedValue[1]
	claims, err := helper.GetJWTClaims(jwtToken)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.Set("claims", claims)

	ctx.Next()
}

func PhotoAuth(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		photoID, err := strconv.Atoi(ctx.Param("photoId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
			return
		}

		claims, exist := ctx.Get("claims")
		if !exist {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "unauthorized"))
			return
		}

		userData, ok := claims.(map[string]interface{})["sub"].(float64)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "User ID not found in claims"))
			return
		}

		var photo model.Photo
		if err := db.First(&photo, photoID).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "Internal Server Error",
				"message": err.Error(),
			})
			return
		}

		userIDInt := int(userData)
		if userIDInt != int(photo.UserID) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
			return
		}

		ctx.Next()
	}
}

func CommentAuth(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		commentID, err := strconv.Atoi(ctx.Param("commentId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
			return
		}

		claims, exist := ctx.Get("claims")
		if !exist {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "unauthorized"))
			return
		}

		userData, ok := claims.(map[string]interface{})["sub"].(float64)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "User ID not found in claims"))
			return
		}

		var comment model.Comment
		if err := db.First(&comment, commentID).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "Internal Server Error",
				"message": err.Error(),
			})
			return
		}

		userIDInt := int(userData)
		if userIDInt != int(comment.UserID) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
			return
		}

		ctx.Next()
	}
}

func SocialMediaAuth(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		socialMediaID, err := strconv.Atoi(ctx.Param("socialmediaId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
			return
		}

		claims, exist := ctx.Get("claims")
		if !exist {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "unauthorized"))
			return
		}

		userData, ok := claims.(map[string]interface{})["sub"].(float64)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "User ID not found in claims"))
			return
		}

		var socialmedias model.SocialMedia
		if err := db.First(&socialmedias, socialMediaID).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "Internal Server Error",
				"message": err.Error(),
			})
			return
		}

		userIDInt := int(userData)
		if userIDInt != int(socialmedias.UserID) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
			return
		}

		ctx.Next()
	}
}
