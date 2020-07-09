package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vpalahia/Cache/cache"
	"github.com/vpalahia/Cache/types"
)

func Insert(c *gin.Context) {
	var todo types.Todo
	c.BindJSON(&todo)

	cache.InsertIntoCache(todo)

	c.JSON(201, "Inserted Successfully")
}
