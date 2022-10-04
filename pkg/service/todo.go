package service

import (
	"Zefirnutiy/simpleTodoApp/pkg/repository"
	"Zefirnutiy/simpleTodoApp/structs"
)

type TodoService struct{
	repo repository.Todo
	listRepo repository.TodoList
}

func NewTodoService(repo repository.Todo, listRepo repository.TodoList) *TodoService{
	return &TodoService{
		repo: repo,
		listRepo: listRepo,
	}
}


func (s *TodoService) Create(userId, listId int, todo structs.Todo) (int, error){
	_, err := s.listRepo.GetListById(listId, userId)
	if err != nil {
		// list does not exist or does not belongs to user
		return 0, err
	}

	return s.repo.Create(listId, todo)
}

func (s *TodoService) GetAll(userId, listId int)([]structs.Todo, error){
	return s.repo.GetAll(userId, listId )
}

func (s *TodoService) GetTodoById(todoId, userId int)(structs.Todo, error){
	return s.repo.GetTodoById(todoId, userId)
}

func (s *TodoService) Delete(todoId, userId int)(error){
	return s.repo.Delete(todoId, userId)
}

func (s *TodoService) Update(userId, todoId int, input structs.UpdateTodo) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, todoId, input)
}