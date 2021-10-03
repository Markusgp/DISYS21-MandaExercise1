package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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
	router.DELETE("/courses/:id", deleteCourse) //For some reason this will never be called, have yet to debug
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
	if err := c.BindJSON(&newCourse); err != nil {
		return
	}
	courses = append(courses, newCourse) // Add new Course to slides
	c.IndentedJSON(http.StatusCreated, newCourse)
}

func putCourse(c *gin.Context) {
	id := c.Param("courseID")
	var newCourse Course

	if err := c.BindJSON(&newCourse); err != nil {
		log.Fatal(err)
		c.Status(http.StatusMethodNotAllowed)
		return
	}

	for i, course := range courses {
		if course.ID == id {
			courses[i] = newCourse
			c.JSON(http.StatusOK, courses[i])
			return
		}
	}
	c.Status(http.StatusNotFound)

}

func deleteCourse(c *gin.Context) { //Handles deletion of selected course based on their id
	courseId := c.Param("courseId")

	for i, course := range courses {
		if course.ID == courseId {
			courses = append(courses[:i], courses[i+1:]...)
			c.Status(http.StatusOK)
			return
		}
	}
	c.Status(http.StatusNotFound)
}
