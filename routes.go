package main

import (
	"github.com/gin-gonic/gin"
)

func Routes(){
    router :=gin.Default()
	router.GET("/todos",getTodos)
	router.POST("/todo",addTodo)
	router.GET("/todo/:id",getTodo)
	router.PATCH("/todo/:id",updateTodo)
	router.DELETE("/todo/:id",deleteTodo)
	router.DELETE("/todos",deleteTodos)
	router.Run("localhost:9090")
}