package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Course struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var courses = []Course{ // course slices to create courses
	{ID: "1", Title: "Distributed Systems"},
	{ID: "2", Title: "Introduction to Database Systems"},
	{ID: "3", Title: "Analysis, Design and Software Architecture"},
}

func main() {
	router := gin.Default()
	router.GET("/courses/:id", getCourseByID)
	router.POST("/courses", postCourses)
	router.PUT("/courses/:id", putCourse)
	router.DELETE("/courses/:id", deleteCourse)
	router.Run("localhost:8080")
}

func getCourseByID(c *gin.Context) { // method returns the course which id matches the parameter
	id := c.Param("id")
	for _, a := range courses {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course not found"})
}

func postCourses(c *gin.Context) { // postCourses adds courses from JSON received in body
	var newCourse Course
	if error := c.BindJSON(&newCourse); error != nil {
		fmt.Println(error)
		c.Status(http.StatusMethodNotAllowed)
		return
	}
	courses = append(courses, newCourse) // Add new Course to slide courses
	c.IndentedJSON(http.StatusCreated, newCourse)
}

func putCourse(c *gin.Context) {
	idOfCourse := c.Param("id")
	var newCourse Course

	if error := c.BindJSON(&newCourse); error != nil {
		fmt.Println(error)
		c.Status(http.StatusMethodNotAllowed)
		return
	}

	for index, course := range courses {
		if course.ID == idOfCourse {
			courses[index] = newCourse
			c.JSON(http.StatusOK, courses[index])
			return
		}
	}
	c.Status(http.StatusNotFound)
}

func deleteCourse(c *gin.Context) { //Handles deletion of selected course based on their id
	idOfCourse := c.Param("id")

	for index, course := range courses {
		if course.ID == idOfCourse {
			courses = append(courses[:index], courses[index+1:]...)
			c.Status(http.StatusOK)
			return
		}
	}
	c.Status(http.StatusNotFound)
}
