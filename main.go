package main

import (
	"errors"
	"net/http"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type todo struct {
    ID          string  `json:"id"`
    Item        string  `json:"item"`
    Completed   bool    `json:"completed"`
}

var todos = []todo{
    {ID:"1", Item:"clean room", Completed: false},
    {ID:"2", Item:"clean poopee", Completed: false},
    {ID:"3", Item:"clean tichou", Completed: true},
}

func getTodos(context *gin.Context) {
    context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
    var newTodo todo

    if err:= context.BindJSON(&newTodo); err !=nil {
        return
    }

    todos = append(todos, newTodo)
    context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodo(context *gin.Context){
   id :=  context.Param("id")
   todo, err:=getTodoById(id)

   if err !=nil {
    context.IndentedJSON(http.StatusNotFound, gin.H{"message":"todo not found"})
    return
   }

   context.IndentedJSON(http.StatusOK, todo)

}

func getTodoById(id string) (*todo, error) {
    for i, todo:= range todos {
        if todo.ID ==id {
            return &todos[i], nil
        }
    }
    return nil, errors.New("todo not found")
}

func toggleTodoStatus(context *gin.Context){
    id :=  context.Param("id")
    todo, err:=getTodoById(id)

    if err !=nil {
        context.IndentedJSON(http.StatusNotFound, gin.H{"message":"todo not found"})
        return
    }

    todo.Completed = !todo.Completed
    context.IndentedJSON(http.StatusOK, todo)
}

func getAuthConfig(c *gin.Context) {
    path := filepath.Join(".", "auth_config.json")
    c.File(path)
}


func main() {
    router := gin.Default()
    router.Use(cors.Default())
    router.GET("/todos", getTodos)
    router.GET("/todos/:id", getTodo)
    router.GET("/auth_config.json", getAuthConfig)
    router.PATCH("/todos/:id", toggleTodoStatus)
    router.POST("/todos", addTodo)
    router.Run("localhost:9090")
}