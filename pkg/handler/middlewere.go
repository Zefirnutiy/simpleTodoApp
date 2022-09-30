package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx = "userId"
)

func (h *Handler) userIdentity(c *gin.Context){
	header := c.Request.Header[authorizationHeader]

	if header[0] == "" {
		newErrorResponce(c, http.StatusUnauthorized, "header is empty")
		return
	}

	headerParts := strings.Split(header[0], " ")

	if len(headerParts) != 2 {
		newErrorResponce(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.ParseToken(headerParts[1])

	if err != nil {
		newErrorResponce(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
}


func getUserId(c *gin.Context)(int, error){
	userId, ok := c.Get(userCtx)

	if !ok {
		newErrorResponce(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id is not found")
	}
	
	indInt, ok := userId.(int)
	
	if !ok{
		newErrorResponce(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}

	return indInt, nil
}