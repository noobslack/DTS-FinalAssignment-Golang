package controller

import (
	"mygram-finalprojectdts/helper"
	"mygram-finalprojectdts/model"
	"mygram-finalprojectdts/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type socialMediaController struct {
	socialMediaRepository repository.ISocialMediaRepository
}

func NewSocialMediaController(socialMediaRepository repository.ISocialMediaRepository) *socialMediaController {
	return &socialMediaController{
		socialMediaRepository: socialMediaRepository,
	}
}

func (pc *socialMediaController) Create(ctx *gin.Context) {
	claims, exist := ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "unauthorized"))
		return
	}

	sub, err := helper.GetSubClaims(claims)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, err.Error()))
	}

	var socialMediaRequest model.SocialMedia
	err = ctx.ShouldBindJSON(&socialMediaRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.CreateResponse(false, nil, err.Error()))
		return
	}

	socialMediaRequest.UserID = uint(sub.(float64))

	socialMediaCreated, err := pc.socialMediaRepository.Create(socialMediaRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helper.CreateResponse(true, gin.H{
		"id":               socialMediaCreated.ID,
		"name":             socialMediaCreated.Name,
		"social_media_url": socialMediaCreated.SocialMediaURL,
		"user_id":          socialMediaCreated.UserID,
		"created_at":       socialMediaCreated.CreatedAt,
	}, ""))
}

func (pc *socialMediaController) GetAll(ctx *gin.Context) {
	claims, exist := ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "unauthorized"))
		return
	}

	sub, err := helper.GetSubClaims(claims)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, err.Error()))
	}

	var socialMediaRequest model.Photo

	socialMediaRequest.UserID = uint(sub.(float64))

	socmedias, err := pc.socialMediaRepository.GetAll()

	socmedResponse := make([]model.SocialMediaResponse, 0)
	for _, socmed := range socmedias {
		socmedResponse = append(socmedResponse, model.SocialMediaResponse{
			ID:             socmed.ID,
			Name:           socmed.Name,
			SocialMediaURL: socmed.SocialMediaURL,
			UserID:         socmed.UserID,
			CreatedAt:      socmed.CreatedAt,
			UpdatedAt:      socmed.UpdatedAt,
			User: model.UpdateSocMedResponse{
				ID:       socmed.User.ID,
				Username: socmed.User.UserName,
			},
		})
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
	}

	ctx.JSON(http.StatusOK, helper.CreateResponse(true, socmedResponse, ""))
}

func (pc *socialMediaController) GetOne(ctx *gin.Context) {
	claims, exist := ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "unauthorized"))
		return
	}

	sub, err := helper.GetSubClaims(claims)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, err.Error()))
	}

	var socialMediaRequest model.SocialMedia

	socialMediaID, err := strconv.Atoi(ctx.Param("socialmediaId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
	}

	socialMediaRequest.UserID = uint(sub.(float64))
	socialMediaRequest.ID = uint(socialMediaID)

	result, err := pc.socialMediaRepository.GetOne(socialMediaRequest.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
	}

	ctx.JSON(http.StatusOK, helper.CreateResponse(true, result, ""))

}

func (pc *socialMediaController) Delete(ctx *gin.Context) {
	claims, exist := ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "unauthorized"))
		return
	}

	sub, err := helper.GetSubClaims(claims)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, err.Error()))
	}

	socialMediaID, err := strconv.Atoi(ctx.Param("socialmediaId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	var socialMediaRequest model.SocialMedia

	socialMediaRequest.UserID = uint(sub.(float64))
	socialMediaRequest.ID = uint(socialMediaID)

	err = pc.socialMediaRepository.Delete(socialMediaRequest.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your SocialMedia Has Been Successfully Deleted",
	})

}

func (pc *socialMediaController) Update(ctx *gin.Context) {
	claims, exist := ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "unauthorized"))
		return
	}

	sub, err := helper.GetSubClaims(claims)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, err.Error()))
	}

	socialMediaID, err := strconv.Atoi(ctx.Param("socialmediaId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	var socialMediaRequest model.SocialMedia
	err = ctx.ShouldBindJSON(&socialMediaRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
		return
	}

	socialMediaRequest.UserID = uint(sub.(float64))
	socialMediaRequest.ID = uint(socialMediaID)

	updatedSocMed, err := pc.socialMediaRepository.Update(socialMediaRequest.ID, socialMediaRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
	}

	ctx.JSON(http.StatusOK, helper.CreateResponse(true, gin.H{
		"id":               updatedSocMed.ID,
		"name":             updatedSocMed.Name,
		"social_media_url": updatedSocMed.SocialMediaURL,
		"user_id":          updatedSocMed.UserID,
		"created_at":       updatedSocMed.CreatedAt,
	}, ""))

}
