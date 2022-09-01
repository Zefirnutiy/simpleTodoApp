package structs

type User struct {
	Id 		 int 	`json:"-"`
	Name 	 string `json:"name" binding:"required"`
	UserName string `json:"userName" binding:"required"`
	Password string	`json:"password" binding:"required"`
}

type Todo struct {
	Id 		 	int 	`json:"-"`
	Title 	 	string  `json:"title"`
	Description string  `json:"description"`
	Done 		bool	`json:"done"`
}

type TodoList struct {
	Id 		 	int 	`json:"-"`
	Title 	 	string  `json:"title"`
	Description string  `json:"description"`
}

type UsersLists struct {
	Id 		int	`json:"-"`
	UserID	int	`json:"userId"`
	ListID	int	`json:"listId"`
}

type ListsTodos struct {
	Id 		int	`json:"-"`
	ListID	int	`json:"listId"`
	TodoID	int	`json:"todoId"`
}