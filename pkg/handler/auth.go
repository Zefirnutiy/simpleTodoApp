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

func (h *Handler) singIn(c *gin.Context){

}