package api

import (
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vpalahia/Cache-Nokia/cache"
	"github.com/vpalahia/Cache-Nokia/types"
)

func Fetch(c *gin.Context) {
	limitstring := c.Query("limit")
	offsetstring := c.Query("offset")

	limit, err := strconv.Atoi(limitstring)

	if err != nil {
		c.JSON(400, "please provide interger value for 'limit' query parameter")
	}

	offset, err := strconv.Atoi(offsetstring)
	if err != nil {
		c.JSON(400, "please provide interger value for 'offset' query parameter")
	}

	todoList := []types.Todo{}
	CacheList := cache.Cache

	var keys []int
	for k := range CacheList.TodoItems {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for i, k := range keys {
		if i+1 <= offset {
			continue
		}
		todo := types.Todo{
			Id:          CacheList.TodoItems[k].Id,
			Task:        CacheList.TodoItems[k].Task,
			Description: CacheList.TodoItems[k].Description,
		}
		todoList = append(todoList, todo)
		if len(todoList) == limit {
			break
		}
	}

	c.JSON(200, todoList)
}
