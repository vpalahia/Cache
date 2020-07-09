package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vpalahia/Cache-Nokia/cache"
	"github.com/vpalahia/Cache-Nokia/types"
)

func Insert(c *gin.Context) {
	var todo types.Todo
	c.BindJSON(&todo)

	cache.InsertIntoCache(todo)

	c.JSON(201, "Inserted Successfully")
}
