package handler

import (
	"Zefirnutiy/simpleTodoApp/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createTodo(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input structs.Todo

	if err := c.BindJSON(&input); err != nil {
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Todo.Create(userId, listId, input)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *Handler) getAllTodos(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusBadRequest, "invalid id param")
		return
	}

	todos, err := h.services.Todo.GetAll(userId, listId)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, todos)
}

func (h *Handler) getTodoById(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	todo, err := h.services.Todo.GetTodoById(todoId, userId)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, todo)
}

func (h *Handler) updateTodo(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, "invalid id param")
		return
	}
	
	var input structs.UpdateTodo
	
	if err := c.BindJSON(&input); err != nil{
		newErrorResponce(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Todo.Update(userId, todoId, input); err != nil{
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusCode{
		Status: "ok",
	})
}

func (h *Handler) deleteTodo(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	todoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	err = h.services.Todo.Delete(todoId, userId)
	if err != nil {
		newErrorResponce(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	c.JSON(http.StatusOK, StatusCode{
		Status: "ok",
	})
}