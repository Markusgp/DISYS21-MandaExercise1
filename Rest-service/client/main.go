package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {
	getCourseByID(2)
	getCourseByID(1)
	getCourseByID(3)
	postCourse("Grundlæggende programmering", 10)
	getCourseByID(10)
	putCourse("Sysdab", 10)
	getCourseByID(10)

}

func getCourseByID(id int) {
	fmt.Printf("Retrieving course by id: %d", id)
	response, error := http.Get("http://localhost:8080/courses/" + strconv.Itoa(id) + "/")

	if error != nil {
		log.Fatal(error)
	} //Error check

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}

func postCourse(title string, id int) {
	fmt.Println("Posting: " + title + "with ID: " + strconv.Itoa(id))

	course := struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}{
		ID:    strconv.Itoa(id),
		Title: title,
	}

	json_data, error := json.Marshal(course)

	if error != nil {
		log.Fatal(error)
	} //Error check

	response, error := http.Post("http://localhost:8080/courses/", "application/json", bytes.NewBuffer(json_data))

	if error != nil {
		log.Fatal(error)
	} //Error check

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}

func putCourse(title string, id int) {

	course := struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}{
		ID:    strconv.Itoa(id),
		Title: title,
	}

	json_data, err := json.Marshal(course)

	if err != nil {
		log.Fatal(err)
	}

	res, err := http.NewRequest(http.MethodPut, "http://localhost:8080/courses/"+strconv.Itoa(id)+"/", bytes.NewBuffer(json_data)) //Der ligger en fejl i de her NewRequest calls

	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

func deleteCourse(id int) {
	fmt.Println("Deleting course with id '" + strconv.Itoa(id) + "'...")
	_, err := http.NewRequest(http.MethodDelete, "http://localhost:8080/courses/"+strconv.Itoa(id)+"/", bytes.NewBuffer(nil))

	if err != nil {
		log.Fatal(err)
	} //Error check
}
