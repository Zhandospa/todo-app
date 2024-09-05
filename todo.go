package todo

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"descripton"`
}
type UserList struct {
	Id      int
	UserID  int
	TitleID int
}
type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"descripton"`
	Done        bool   `json:"done`
}
type ListsItem struct {
	Id      int
	UserID  int
	TitleID int
}
