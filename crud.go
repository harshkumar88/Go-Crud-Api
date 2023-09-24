package main

import (
	// "errors"
	"net/http"
    "database/sql"
    "fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type todo struct {
    Id        string    `json:"id" `
    Item      string `json:"item"`
    Completed  bool   `json:"completed"  binding:"required"`
}


func getTodos(context *gin.Context) {
    // Query all TODO items from the database
    rows, err := dB.Query("SELECT * FROM todosTable")
    if err != nil {
        context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
        return
    }
    defer rows.Close()

    // Create a slice to hold the TODO items
    var todos= []todo{}

    // Iterate through the result set and append each row to the 'todos' slice
    for rows.Next() {
        var t todo
        if err := rows.Scan(&t.Id,&t.Item, &t.Completed); err != nil {
            context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Failed to read data from database"})
            return
        }
        todos = append(todos, t)
    }

    // Check for errors during iteration
    if err := rows.Err(); err != nil {
        context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Database error"})
        return
    }

    // Return the 'todos' slice as JSON
    context.IndentedJSON(http.StatusOK, todos)
}



func addTodo(context *gin.Context) {
    var newTodo todo
    // Bind the JSON request body to the 'newTodo' struct

    
    if err := context.ShouldBindJSON(&newTodo); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }


    if(newTodo.Id=="" || newTodo.Item==""){
         context.JSON(http.StatusInternalServerError, gin.H{"message": "All Fields are required"})
         return;
    }


    // Check if the TODO item already exists in the database
    var existingTodo todo
    err := dB.QueryRow("SELECT id, item, completed FROM todosTable WHERE item = ?", newTodo.Item).Scan(&existingTodo.Id,&existingTodo.Item, &existingTodo.Completed)

    if err == nil {
        // If the row was found, return a conflict response
        context.JSON(http.StatusConflict, gin.H{"message": "TODO item already exists"})
        return
    } else if err != sql.ErrNoRows {
        // Handle other database errors
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Database error"})
        return
    }


    // Item doesn't exist, so insert the new TODO item into the database
    _, err = dB.Exec("INSERT INTO todosTable (id,item, completed) VALUES (?, ?, ?)",newTodo.Id,newTodo.Item, newTodo.Completed)

    if err != nil {
        // Handle the database insertion error
        context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add item"})
        return
    }

    context.JSON(http.StatusOK, gin.H{"message": "TODO item added successfully"})

}


func getTodo(context *gin.Context){
    id :=context.Param("id")

    var row todo
    err :=dB.QueryRow("Select id,item,completed from todosTable where id=?",id).Scan(&row.Id,&row.Item,&row.Completed)

    if(err!=nil){
        context.IndentedJSON(http.StatusOK,gin.H{"message":err.Error()})
        return;
    }
	
	context.IndentedJSON(http.StatusOK,row)
}


func deleteTodo(context *gin.Context){

	id:=context.Param("id")
	var row todo;
    err :=dB.QueryRow("Select id from todosTable where id=?",id).Scan(&row.Id)

    if(err!=nil){
        context.IndentedJSON(http.StatusConflict,gin.H{"message":err.Error()})
        return;
    }


     dB.Exec("Delete from todosTable where id=?",id)
     context.IndentedJSON(http.StatusOK,gin.H{"message":"Successfully deleted Todo"})

   
}


func deleteTodos(context *gin.Context){

    _,err := dB.Exec("Delete from todosTable")

     if(err!=nil){
        context.IndentedJSON(http.StatusConflict,gin.H{"message":err.Error()})
        return;
    }
    context.IndentedJSON(http.StatusOK,gin.H{"message":"Successfully deleted Todos"})
}


func updateTodo(context *gin.Context){
	id := context.Param("id")

    var row todo
    err:=dB.QueryRow("Select id from todosTable where id = ?",id).Scan(&row.Id)

    if(err!=nil){
        context.IndentedJSON(http.StatusConflict,gin.H{"message":err.Error()})
        return;
    }

    var updatedTodo todo;

    err = context.ShouldBindJSON(&updatedTodo)

    if(err!=nil){
        fmt.Println("hai")
        context.IndentedJSON(http.StatusConflict,gin.H{"message":err.Error()})
        return;
    }

    dB.Exec("update todosTable  set item=? ,condition=? ",updatedTodo.Item,updatedTodo.Completed)
    context.IndentedJSON(http.StatusOK,gin.H{"message":"Successfully Updated todo"})
	
}