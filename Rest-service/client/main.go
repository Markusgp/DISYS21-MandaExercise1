package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	getCourseByID(1)
	getCourseByID(2)
	getCourseByID(3)
	postCourse("Grundlæggende programmering", 10)
	getCourseByID(10)
	putCourse("Sysdab", 10)
	getCourseByID(10)
	deleteCourse(10)
	getCourseByID(10)
}

//The GET call that takes in an ID in the form of an integer.
func getCourseByID(id int) {
	fmt.Println("Retrieving course by id: " + strconv.Itoa(id))

	response, error := http.Get("http://localhost:8080/courses/" + strconv.Itoa(id) + "/") //Call to server.

	if error != nil {
		fmt.Println(error)
	} //Error check

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body)) //print out the response.
}


//The POST call that takes in a title: string and an id: integer
func postCourse(title string, id int) {
	fmt.Println("Posting: " + title + " with ID: " + strconv.Itoa(id))

	course := struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}{
		ID:    strconv.Itoa(id),
		Title: title,
	} //Creating the data as a course struct

	json_data, error := json.Marshal(course) //Creating the struct as json data.

	if error != nil {
		fmt.Println(error)
	} //Error check

	response, error := http.Post("http://localhost:8080/courses/", "application/json", bytes.NewBuffer(json_data)) //Call to the POST execution

	if error != nil {
		fmt.Println(error)
	} //Error check

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body)) //print out the response.
}

func putCourse(title string, id int) {
	fmt.Println("Changing the course with id: " + strconv.Itoa(id) + " to title: " + "'" + title + "'")

	course := struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	}{
		ID:    strconv.Itoa(id),
		Title: title,
	} //Creating the data as a course struct

	json_data, error := json.Marshal(course) //Creating json data from the struct

	if error != nil {
		fmt.Println(error)
	} //Error check

	client := &http.Client{} //Establishing the client
	request, _ := http.NewRequest(http.MethodPut, "http://localhost:8080/courses/" + strconv.Itoa(id) + "/", bytes.NewBuffer(json_data)) //Call to PUT request.
	response, error := client.Do(request) //Making the client do the request.
	
	if error != nil {
		fmt.Println(error)
		return
	} //Error check

	defer response.Body.Close() //Waits until end of execution to execute.
	
	//body, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(body)) //Print out response
}

func deleteCourse(id int) {
	fmt.Println("Deleting course with id: " + strconv.Itoa(id))


	client := &http.Client{} //Establishing the client
	request, error := http.NewRequest(http.MethodDelete, "http://localhost:8080/courses/"+strconv.Itoa(id)+"/", bytes.NewBuffer(nil)) //Call to the DELETE request.
	response, error := client.Do(request) //Making the client do the request.

	if error != nil {
		fmt.Println(error)
	} //Error check

	defer response.Body.Close() //Waits until end of execution to execute.

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body)) //Print out response
}
