package controller

import (
	"mygram-finalprojectdts/helper"
	"mygram-finalprojectdts/model"
	"mygram-finalprojectdts/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type photoController struct {
	photoRepository repository.IPhotoRepository
}

func NewPhotoController(photoRepository repository.IPhotoRepository) *photoController {
	return &photoController{
		photoRepository: photoRepository,
	}
}

func (pc *photoController) Create(ctx *gin.Context) {
	claims, exist := ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "unauthorized"))
		return
	}

	sub, err := helper.GetSubClaims(claims)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, err.Error()))
	}

	var uploadPhoto model.Photo
	err = ctx.ShouldBindJSON(&uploadPhoto)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.CreateResponse(false, nil, err.Error()))
		return
	}

	uploadPhoto.UserID = uint(sub.(float64))

	photoCreated, err := pc.photoRepository.Create(uploadPhoto)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helper.CreateResponse(true, gin.H{
		"id":         photoCreated.ID,
		"title":      photoCreated.Title,
		"caption":    photoCreated.Caption,
		"photo_url":  photoCreated.PhotoUrl,
		"user_id":    photoCreated.UserID,
		"created_at": photoCreated.CreatedAt,
	}, ""))
}

func (pc *photoController) GetAll(ctx *gin.Context) {
	claims, exist := ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "unauthorized"))
		return
	}

	sub, err := helper.GetSubClaims(claims)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, err.Error()))
	}

	var photoRequest model.Photo

	photoRequest.UserID = uint(sub.(float64))

	photos, err := pc.photoRepository.GetAll()

	photoResponse := make([]model.ResponsePhoto, 0)
	for _, photo := range photos {
		photoResponse = append(photoResponse, model.ResponsePhoto{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoUrl:  photo.PhotoUrl,
			UserID:    photo.UserID,
			CreatedAt: photo.CreatedAt,
			UpdatedAt: photo.UpdatedAt,
			User: model.UpdatePhotoResponse{
				Email:    photo.User.Email,
				Username: photo.User.UserName,
			},
		})
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
	}

	ctx.JSON(http.StatusOK, helper.CreateResponse(true, photoResponse, ""))
}

func (pc *photoController) GetOne(ctx *gin.Context) {
	claims, exist := ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "unauthorized"))
		return
	}

	sub, err := helper.GetSubClaims(claims)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, err.Error()))
	}

	var photoRequest model.Photo

	socialMediaID, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
	}

	photoRequest.UserID = uint(sub.(float64))
	photoRequest.ID = uint(socialMediaID)

	result, err := pc.photoRepository.GetOne(photoRequest.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
	}

	ctx.JSON(http.StatusOK, helper.CreateResponse(true, result, ""))
}

func (pc *photoController) Delete(ctx *gin.Context) {
	claims, exist := ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "unauthorized"))
		return
	}

	sub, err := helper.GetSubClaims(claims)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, err.Error()))
	}

	photoID, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	var photoRequest model.Photo

	photoRequest.UserID = uint(sub.(float64))
	photoRequest.ID = uint(photoID)

	err = pc.photoRepository.Delete(photoRequest.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your Photo Has Been Successfully Deleted",
	})

}

func (pc *photoController) Update(ctx *gin.Context) {
	claims, exist := ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "unauthorized"))
		return
	}

	sub, err := helper.GetSubClaims(claims)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, err.Error()))
	}

	photoID, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	var photoRequest model.Photo

	err = ctx.ShouldBindJSON(&photoRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
		return
	}

	photoRequest.UserID = uint(sub.(float64))
	photoRequest.ID = uint(photoID)

	updatedPhoto, err := pc.photoRepository.Update(photoRequest.ID, photoRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
	}

	ctx.JSON(http.StatusOK, helper.CreateResponse(true, gin.H{
		"id":        updatedPhoto.ID,
		"title":     updatedPhoto.Title,
		"caption":   updatedPhoto.Caption,
		"photo_url": updatedPhoto.PhotoUrl,
		"user_id":   updatedPhoto.UserID,
		"update_at": updatedPhoto.UpdatedAt,
	}, ""))

}
