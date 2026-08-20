package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/health", health)                     // get health
	router.GET("/todos", getTodos)                    // get all todos
	router.GET("/todos/:id", getTodoByID)             // get single todo
	router.POST("/todos", postTodos)                  // add todo
	router.PATCH("/todos/:id", updateTodo)            // update todo
	router.PATCH("/todos/:id/complete", completeTodo) // complete todo
	router.DELETE("/todos/:id", deleteTodo)           // delete todo

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

// returns server health
// TODO: update with sqlite when implemented
func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
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

// getTodo: Utility function that takes an ID and returns the todo and if it was found
func getTodo(id int) (*todo, bool) {
	for i := range todos {
		if todos[i].ID == id {
			return &todos[i], true
		}
	}

	return nil, false
}

// getTodoByID locates the todo whose ID value matches the id
// parameter sent by the client, then returns that todo as a response.
func getTodoByID(c *gin.Context) {
	// Convert id to int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid ID"})
		return
	}

	t, found := getTodo(id)

	// todo not found
	if !found {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}

	// todo found
	c.IndentedJSON(http.StatusOK, t)
}

// updateTodo updates the properties of a todo
func updateTodo(c *gin.Context) {

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

// deleteTodo deletes a todo
func deleteTodo(c *gin.Context) {
	// Convert id to int
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid ID"})
		return
	}

	for i := range todos {
		if todos[i].ID == id {

			// Using append function to combine two slices
			// first slice is the slice of all the elements before the given index
			// second slice is the slice of all the elements after the given index
			// append function appends the second slice to the end of the first slice
			todos = append(todos[:i], todos[i+1:]...)

			c.Status(http.StatusNoContent)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}
