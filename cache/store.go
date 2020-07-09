package cache

import (
	"github.com/vpalahia/Cache/db"
	"github.com/vpalahia/Cache/types"
)

type cache struct {
	TodoItems types.TodoMap
}

var Cache = cache{
	TodoItems: types.TodoMap{},
}

func updateCache(todoArr []types.Todo) {
	ma := make(types.TodoMap)
	for i, v := range todoArr {
		ma[i+1] = v
	}
	Cache.TodoItems = ma
}

func InsertIntoCache(todo types.Todo) {
	todo.Id = len(Cache.TodoItems) + 1

	Cache.TodoItems[todo.Id] = todo

	db.InsertIntoDB(todo)
}
