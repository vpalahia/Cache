package types

type Todo struct {
	Id          int    `json:"id" db:"id"`
	Task        string `json:"task" db:"task"`
	Description string `json:"description" db:"description"`
}

type TodoMap map[int]Todo
