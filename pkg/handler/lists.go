package handler

import (
	"Zefirnutiy/simpleTodoApp/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	
	var input structs.TodoList
	
	if err := c.BindJSON(&input); err != nil{
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	listId, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, "user id not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"listId": listId,
	})
}

type getAllLists struct {
	Data	[]structs.TodoList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	todoList, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, getAllLists{
		Data: todoList,
	})
}

func (h *Handler) getListById(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	todoList, err := h.services.TodoList.GetListById(listId, userId)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, todoList)
}

func (h *Handler) updateList(c *gin.Context){

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, "invalid id param")
		return
	}
	
	var input structs.UpdateTodoList
	
	if err := c.BindJSON(&input); err != nil{
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.TodoList.Update(userId, listId, input); err != nil{
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusCode{
		Status: "ok",
	})
}

func (h *Handler) deleteList(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	err = h.services.TodoList.DeleteList(userId, listId)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, StatusCode{
		Status: "ok",
	})
}