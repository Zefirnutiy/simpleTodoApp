package handler

import (
	"Zefirnutiy/simpleTodoApp/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) singUp(c *gin.Context){
	var input structs.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)

	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":id,
	})
}

type singInInput struct {
	UserName string `json:"userName" binding:"required"`
	Password string	`json:"password" binding:"required"`
}

func (h *Handler) singIn(c *gin.Context){
	var input singInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.UserName, input.Password)

	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":token,
	})
}