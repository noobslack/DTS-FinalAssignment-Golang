package controller

import (
	"mygram-finalprojectdts/helper"
	"mygram-finalprojectdts/model"
	"mygram-finalprojectdts/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userRepository repository.IUserRepository
}

func NewUserController(userRepository repository.IUserRepository) *userController {
	return &userController{
		userRepository: userRepository,
	}
}

func (uc *userController) Register(ctx *gin.Context) {
	var newUser model.User
	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
		return
	}

	hashPassword, err := helper.Hash([]byte(newUser.Password))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
		return
	}

	newUser.Password = string(hashPassword)

	createdUser, err := uc.userRepository.Create(newUser)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.CreateResponse(true, gin.H{
		"age":      createdUser.Age,
		"email":    createdUser.Email,
		"id":       createdUser.ID,
		"username": createdUser.UserName,
	}, ""))
}

func (uc *userController) Login(ctx *gin.Context) {
	var requestedUser model.User

	err := ctx.ShouldBindJSON(&requestedUser)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
		return
	}

	user, err := uc.userRepository.GetByUsername(requestedUser.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
		return
	}
	if !helper.HashMatched([]byte(user.Password), []byte(requestedUser.Password)) {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, "invalid email/password"))
		return
	}

	token, err := helper.GenerateJWTToken(user.Email, user.ID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "UNAUTHORIZED"))
		return
	}

	ctx.JSON(http.StatusOK, helper.CreateResponse(true, gin.H{
		"token": token,
	}, ""))

}

func (uc *userController) Update(ctx *gin.Context) {

	claims, exist := ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "unauthorized"))
		return
	}

	sub, err := helper.GetSubClaims(claims)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, err.Error()))
		return
	}

	var userRequest model.User
	err = ctx.ShouldBindJSON(&userRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
		return
	}

	userRequest.ID = uint(sub.(float64))

	updatedUser, err := uc.userRepository.Update(userRequest.ID, userRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
		return
	}

	updateResponse := make([]model.UpdateUserResponse, 0)

	updateResponse = append(updateResponse, model.UpdateUserResponse{
		ID:        updatedUser.ID,
		Username:  updatedUser.UserName,
		Age:       updatedUser.Age,
		Email:     updatedUser.Email,
		UpdatedAt: updatedUser.UpdatedAt,
	})

	ctx.JSON(http.StatusOK, helper.CreateResponse(true, updateResponse, ""))
}

func (uc *userController) Delete(ctx *gin.Context) {
	claims, exist := ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "unauthorized"))
		return
	}
	sub, err := helper.GetSubClaims(claims)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, err.Error()))
		return
	}

	userID := sub.(float64)

	err = uc.userRepository.Delete(uint(userID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.CreateResponse(false, nil, err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, helper.CreateResponse(true, gin.H{
		"message": "your account has been sucesfully deleted",
	}, ""))

}
