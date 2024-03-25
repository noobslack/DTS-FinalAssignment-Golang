package controller

import (
	"mygram-finalprojectdts/helper"
	"mygram-finalprojectdts/model"
	"mygram-finalprojectdts/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type commentController struct {
	commentRepository repository.ICommentRepository
}

func NewCommentController(commentRepository repository.ICommentRepository) *commentController {
	return &commentController{
		commentRepository: commentRepository,
	}
}

func (pc *commentController) Create(ctx *gin.Context) {
	claims, exist := ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "unauthorized"))
		return
	}

	sub, err := helper.GetSubClaims(claims)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, err.Error()))
	}

	var CommentRequest model.Comment
	err = ctx.ShouldBindJSON(&CommentRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, helper.CreateResponse(false, nil, err.Error()))
		return
	}

	CommentRequest.UserID = uint(sub.(float64))

	commentCreated, err := pc.commentRepository.Create(CommentRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, helper.CreateResponse(true, gin.H{
		"id":         commentCreated.ID,
		"message":    commentCreated.Message,
		"photo_id":   commentCreated.PhotoID,
		"user_id":    commentCreated.UserID,
		"created_at": commentCreated.CreatedAt,
	}, ""))
}

func (pc *commentController) GetAll(ctx *gin.Context) {
	claims, exist := ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "unauthorized"))
		return
	}

	sub, err := helper.GetSubClaims(claims)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, err.Error()))
	}

	var CommentRequest model.Comment

	CommentRequest.UserID = uint(sub.(float64))

	comments, err := pc.commentRepository.GetAll()

	commentsResponse := make([]model.CommentResponse, 0)
	for _, comment := range comments {
		commentsResponse = append(commentsResponse, model.CommentResponse{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoID:   comment.PhotoID,
			UserID:    comment.UserID,
			UpdatedAt: comment.UpdatedAt,
			CreatedAt: comment.CreatedAt,
			User: model.UpdateCommentResponse{
				ID:       comment.User.ID,
				Email:    comment.User.Email,
				Username: comment.User.UserName,
			},
			Photo: model.PostPhoto{
				ID:       comment.User.ID,
				Title:    comment.Photo.Title,
				Caption:  comment.Photo.Caption,
				PhotoUrl: comment.Photo.PhotoUrl,
				UserID:   comment.Photo.UserID,
			},
		})
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
	}

	ctx.JSON(http.StatusOK, helper.CreateResponse(true, commentsResponse, ""))
}

func (pc *commentController) GetOne(ctx *gin.Context) {
	claims, exist := ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "unauthorized"))
		return
	}

	sub, err := helper.GetSubClaims(claims)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, err.Error()))
	}

	var commentRequest model.Comment

	commentID, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
	}

	commentRequest.UserID = uint(sub.(float64))
	commentRequest.ID = uint(commentID)

	comment, err := pc.commentRepository.GetOne(commentRequest.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
	}

	ctx.JSON(http.StatusOK, helper.CreateResponse(true, comment, ""))

}

func (pc *commentController) Delete(ctx *gin.Context) {
	claims, exist := ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "unauthorized"))
		return
	}

	sub, err := helper.GetSubClaims(claims)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, err.Error()))
	}

	commentID, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	var commentRequest model.Comment

	commentRequest.UserID = uint(sub.(float64))
	commentRequest.ID = uint(commentID)

	err = pc.commentRepository.Delete(commentRequest.ID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your Comment Has Been Successfully Deleted",
	})

}

func (pc *commentController) Update(ctx *gin.Context) {
	claims, exist := ctx.Get("claims")
	if !exist {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, "unauthorized"))
		return
	}

	sub, err := helper.GetSubClaims(claims)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, helper.CreateResponse(false, nil, err.Error()))
	}

	commentID, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	var commentRequest model.Comment

	err = ctx.ShouldBindJSON(&commentRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
		return
	}

	commentRequest.UserID = uint(sub.(float64))
	commentRequest.ID = uint(commentID)

	updatedComment, err := pc.commentRepository.Update(commentRequest.ID, commentRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, helper.CreateResponse(false, nil, err.Error()))
	}

	ctx.JSON(http.StatusOK, helper.CreateResponse(true, gin.H{
		"id":         updatedComment.ID,
		"message":    updatedComment.Message,
		"photo_id":   updatedComment.PhotoID,
		"user_id":    updatedComment.UserID,
		"created_at": updatedComment.CreatedAt,
	}, ""))

}
