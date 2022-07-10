package main

import (
	"fmt"
	//	"log"
	"net/http"
	//	"os"
	//	"strconv"

	"github.com/gin-gonic/gin"
	//"github.com/joho/godotenv"
)

type Numeric struct {
	Num1   float32 `json:"num1"`
	Num2   float32 `json: "num2"`
	Result float32 `json:"result"`
}

func main() {
	r := gin.Default()

	r.POST("/add", add)
	r.POST("/subtract", subtract)
	r.POST("/multiply", multiply)
	r.POST("/divide", divide)

	r.Run(":8090") // run 8080
	fmt.Println("Server is running")
}

func add(c *gin.Context) {
	var num Numeric

	if err := c.BindJSON(&num); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}
	num.Result = num.Num1 + num.Num2
	c.IndentedJSON(http.StatusOK, num)
}

func subtract(c *gin.Context) {
	var num Numeric

	if err := c.BindJSON(&num); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}
	num.Result = num.Num1 - num.Num2
	c.IndentedJSON(http.StatusOK, num)
}

func multiply(c *gin.Context) {
	var num Numeric

	if err := c.BindJSON(&num); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}
	num.Result = num.Num1 * num.Num2
	c.IndentedJSON(http.StatusOK, num)
}

func divide(c *gin.Context) {
	var num Numeric

	if err := c.BindJSON(&num); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}
	num.Result = num.Num1 / num.Num2
	c.IndentedJSON(http.StatusOK, num)
}
