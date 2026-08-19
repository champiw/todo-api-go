package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/todos", getTodos)                    // get all todos
	router.GET("/todos/:id", getTodoByID)             // get a single todo
	router.POST("/todos", postTodos)                  // add a todo
	router.PATCH("/todos/:id/complete", completeTodo) // complete todo

	router.Run("localhost:8000")
}

type todo struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: 1, Text: "test1", Completed: false},
	{ID: 2, Text: "test2", Completed: true},
	{ID: 3, Text: "test3", Completed: false},
}

// getTodos replies with the list of all todos as json
func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

// postTodos adds a todo from JSON received in the request body
func postTodos(c *gin.Context) {
	var newTodo todo

	// Call BindJSON to bind the received JSON to newTodo
	// & to pass the pointer, so the variable content gets modified
	err := c.BindJSON(&newTodo)

	if err != nil {
		return
	}

	// Add the new todo to the slice
	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

// getTodoByID locates the todo whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getTodoByID(c *gin.Context) {
	// Convert id to int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid ID"})
		return
	}

	// Loop over the list of todos, looking for
	// a todo whose ID value matches the parameter.
	for _, a := range todos {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func completeTodo(c *gin.Context) {
	// Convert id to int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid ID"})
		return
	}

	// Loop over the list of todos, looking for
	// a todo whose ID value matches the parameter.
	for i := range todos {
		if todos[i].ID == id {
			todos[i].Completed = !todos[i].Completed

			c.IndentedJSON(http.StatusOK, todos[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}
