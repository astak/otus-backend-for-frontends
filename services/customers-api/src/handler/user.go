package handler

import (
	"net/http"

	"github.com/Astak/otus-docker-basics-homework/web-service-gin/auth"
	common_auth "github.com/Astak/otus-docker-basics-homework/web-service-gin/common/auth"
	"github.com/Astak/otus-docker-basics-homework/web-service-gin/dto"
	"github.com/Astak/otus-docker-basics-homework/web-service-gin/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProfile(c *gin.Context) {
	claims, ok := c.Get(common_auth.JwtKey)
	if !ok {
		errModel := dto.NewNoAccountInfoError()
		c.IndentedJSON(http.StatusBadRequest, errModel)
		return
	}
	jwtClaims := claims.(auth.JWTClaim)
	user, err := h.UserRepo.GetProfile(jwtClaims.Subject)
	if err != nil {
		errModel := dto.NewBadRequestError(err.Error())
		c.IndentedJSON(http.StatusBadRequest, errModel)
		return
	}
	response := dto.ProfileResponse{
		AccountId: jwtClaims.Subject,
		UserName:  jwtClaims.Username,
		Email:     jwtClaims.Email,
	}
	if user != nil {
		response.FirstName = user.FirstName
		response.LastName = user.LastName
		response.Phone = user.Phone
	}
	c.IndentedJSON(http.StatusOK, response)
}

func (h *Handler) UpdateProfile(c *gin.Context) {
	claims, ok := c.Get(common_auth.JwtKey)
	if !ok {
		errModel := dto.NewNoAccountInfoError()
		c.IndentedJSON(http.StatusBadRequest, errModel)
		return
	}
	jwtClaims := claims.(auth.JWTClaim)
	request := new(dto.ProfileRequest)
	if err := c.ShouldBindJSON(request); err != nil {
		errModel := dto.NewBadRequestError(err.Error())
		c.IndentedJSON(http.StatusBadRequest, errModel)
		return
	}
	u := &models.Profile{
		AccountId: jwtClaims.Subject,
		UserName:  jwtClaims.Username,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     jwtClaims.Email,
		Phone:     request.Phone,
	}
	newUser, err := h.UserRepo.UpdateProfile(u)
	if err != nil {
		errModel := dto.NewBadRequestError(err.Error())
		c.IndentedJSON(http.StatusBadRequest, errModel)
		return
	}
	response := dto.ProfileResponse{
		AccountId: newUser.AccountId,
		UserName:  newUser.UserName,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		Phone:     newUser.Phone,
	}
	c.IndentedJSON(http.StatusOK, response)
}
