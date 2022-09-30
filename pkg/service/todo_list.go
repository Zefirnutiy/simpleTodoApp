package service

import (
	"Zefirnutiy/simpleTodoApp/pkg/repository"
	"Zefirnutiy/simpleTodoApp/structs"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo:repo}
}

func (s *TodoListService) Create(userId int, list structs.TodoList)(int, error){
	return s.repo.Create(userId, list) 
}

func (s *TodoListService) GetAll(userId int)([]structs.TodoList, error){
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetListById(listId, userId int)(structs.TodoList, error){
	return s.repo.GetListById(listId, userId)
}

func (s *TodoListService) DeleteList(userId, listId int) error{
	return s.repo.DeleteList(userId, listId)
}

func (s *TodoListService) Update(userId, listId int, input structs.UpdateTodoList) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)
}