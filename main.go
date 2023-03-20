package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID       float64 `json:"id"`
	Username string  `json:"username"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
}

var users = []User{
	{ID: 1, Username: "edcnogueira", Name: "Edjancarlos", Email: "edjan.nogueira14@gmail.com"},
	{ID: 2, Username: "lucevald", Name: "Lucas", Email: "lucas.evaldo14@gmail.com"},
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.Run("localhost:8080")
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
