package utils

import (
	"net/http"
	"strconv"
	s "strings"

	"github.com/gin-gonic/gin"
)

func GetClassStudents(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	class := c.Query("class")
	var students []Student
	db.Where("class_id = ?", class).Find(&students)
	c.JSON(http.StatusOK, students)
}

func GetStudentGrades(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Query("id")
	var grades []Grade
	db.Where("student_id = ?", id).Find(&grades)
	c.JSON(http.StatusOK, grades)
}

func GetStudentInfo(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var student Student
	id := c.Query("id")
	db.Where("username = ?", id).First(&student)
	c.JSON(http.StatusOK, student)
}

func PostStudentInfo(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var student Student
	c.Bind(&student)
	if student.Username == "" {
		var lastStudent Student
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
