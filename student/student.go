// Package student implements the various methods related to getting and setting student data.
package student

import (
	m "MTDS-REST/model"

	"net/http"
	"strconv"
	s "strings"

	"github.com/gin-gonic/gin"
)

var initDb = m.InitDb

// GetClassStudents returns the students in a specific class.
//
// Input: Class ID.
//
// Output: []Student.
//
// Example URL: http://localhost:8080/api/v1/class/students?class=C1
func GetClassStudents(c *gin.Context) {
	db := initDb()
	defer db.Close()
	class := c.Query("class")
	var students []m.Student
	db.Where("class_id = ?", class).Order("LENGTH(username), username").Find(&students)
	if len(students) > 0 {
		c.JSON(http.StatusOK, students)
	} else {
		c.JSON(http.StatusBadRequest, make([]string, 0))
	}
}

// GetStudentGrades returns the grades of a specific student.
//
// Input: Student ID.
//
// Output: []Grade.
//
// Example URL: http://localhost:8080/api/v1/student/grades?id=S1
func GetStudentGrades(c *gin.Context) {
	db := initDb()
	defer db.Close()
	id := c.Query("id")
	var grades []m.Grade
	db.Where("student_id = ?", id).Find(&grades)
	if len(grades) > 0 {
		c.JSON(http.StatusOK, grades)
	} else {
		c.JSON(http.StatusBadRequest, make([]string, 0))
	}
}

// GetStudentInfo returns the information pertaining to a specific student.
//
// Input: Student ID.
//
// Output: Student Object.
//
// Example URL: http://localhost:8080/api/v1/student/info?id=S1
func GetStudentInfo(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var student m.Student
	id := c.Query("id")
	db.Where("username = ?", id).First(&student)
	if student.Username == "" || student.FirstName == "" {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

// PostStudentInfo edits the information of a specific student if his ID exists in the database; otherwise, it creates a new student with the provided data.
//
// Input: Student Data (ID Optional).
//
// Output: Newly created/edited student.
//
// Example URL: http://localhost:8080/api/v1/student/info
func PostStudentInfo(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var student m.Student
	c.Bind(&student)
	if student.FirstName == "" || student.LastName == "" || student.ClassID == "" {
		c.String(http.StatusBadRequest, "Invalid Input")
	} else {
		if student.Username == "" {
			var lastStudent m.Student
			db.Limit(1).Order("LENGTH(username) desc, username desc").Find(&lastStudent)
			id := lastStudent.Username
			id = s.Trim(id, "S")
			num, _ := strconv.Atoi(id)
			num++
			student.Username = "S" + strconv.Itoa(num)
		}
		db.Save(&student)
		c.JSON(http.StatusOK, student)
	}
}
