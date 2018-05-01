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
func GetClassStudents(c *gin.Context) {
	db := initDb()
	defer db.Close()
	class := c.Query("class")
	var students []m.Student
	db.Where("class_id = ?", class).Order("LENGTH(username), username").Find(&students)
	c.JSON(http.StatusOK, students)
}

// GetStudentGrades returns the grades of a specific student.
func GetStudentGrades(c *gin.Context) {
	db := initDb()
	defer db.Close()
	id := c.Query("id")
	var grades []m.Grade
	db.Where("student_id = ?", id).Find(&grades)
	c.JSON(http.StatusOK, grades)
}

// GetStudentInfo returns the information pertaining to a specific student.
func GetStudentInfo(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var student m.Student
	id := c.Query("id")
	db.Where("username = ?", id).First(&student)
	c.JSON(http.StatusOK, student)
}

// PostStudentInfo edits the information of a specific student if his ID exists in the database; otherwise, it creates a new student with the provided data.
func PostStudentInfo(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var student m.Student
	c.Bind(&student)
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
