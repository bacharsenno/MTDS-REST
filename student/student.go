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
// Example URL: http://localhost:8080/api/v1/class/C1/students
func GetClassStudents(c *gin.Context) {
	db := initDb()
	defer db.Close()
	class := c.Params.ByName("cid")
	var students []m.Student
	db.Where("class_id = ?", class).Order("LENGTH(username), username").Find(&students)
	if len(students) > 0 {
		var objectsWithLink []m.StudentWithLink
		var tempObj m.StudentWithLink
		for i := 0; i < len(students); i++ {
			tempObj.Student = students[i]
			tempObj.Link = "http://localhost:8080/api/v1/student/" + students[i].Username + "/info"
			objectsWithLink = append(objectsWithLink, tempObj)
		}
		c.JSON(http.StatusOK, objectsWithLink)
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
// Example URL: http://localhost:8080/api/v1/student/S1/grades
func GetStudentGrades(c *gin.Context) {
	db := initDb()
	defer db.Close()
	id := c.Params.ByName("sid")
	var grades []m.Grade
	db.Where("student_id = ?", id).Find(&grades)
	if len(grades) > 0 {
		c.JSON(http.StatusOK, grades)
	} else {
		c.JSON(http.StatusOK, make([]string, 0))
	}
}

// GetStudentInfo returns the information pertaining to a specific student.
//
// Input: Student ID.
//
// Output: Student Object.
//
// Example URL: http://localhost:8080/api/v1/student/S1/info
func GetStudentInfo(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var student m.Student
	id := c.Params.ByName("sid")
	db.Where("username = ?", id).First(&student)
	if student.Username == "" || student.FirstName == "" {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

// GetStudentSubjects returns the subjects studied by a specific student.
//
// Input: Student ID.
//
// Output: TeachClass Array.
//
// Example URL: http://localhost:8080/api/v1/parent/P1/students/S1/subjects
func GetStudentSubjects(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var teachClasses []m.TeachClass
	var tcWithLink []m.TeachClassWithLink
	var temptc m.TeachClassWithLink
	id := c.Params.ByName("sid")
	db.Table("teach_classes tc, students s").Where("s.username = ? and tc.class_id = s.class_id", id).Find(&teachClasses)
	if len(teachClasses) > 0 {
		for i := 0; i < len(teachClasses); i++ {
			temptc.TeachClass = teachClasses[i]
			temptc.Link = "http://localhost:8080/api/v1/teacher/" + teachClasses[i].TeacherID + "/classes?class=" + teachClasses[i].ClassID
			tcWithLink = append(tcWithLink, temptc)
		}
		c.JSON(http.StatusOK, tcWithLink)
	} else {
		c.JSON(http.StatusOK, make([]string, 0))
	}
}

// GetStudentParents returns the Parents of a specific student.
//
// Input: Student ID.
//
// Output: Parent Array.
//
// Example URL: http://localhost:8080/api/v1/student/S1/parents
func GetStudentParents(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var parents []m.ParentOf
	id := c.Params.ByName("sid")
	db.Where("student_id = ?", id).Find(&parents)
	var parentsWithLink []m.ParentWithLink
	var temp m.ParentWithLink
	if len(parents) > 0 {
		for i := 0; i < len(parents); i++ {
			temp.ParentOf = parents[i]
			temp.Link = "http://localhost:8080/api/v1/parent/" + parents[i].ParentID + "/info"
			parentsWithLink = append(parentsWithLink, temp)
		}
		c.JSON(http.StatusOK, parentsWithLink)
	} else {
		c.JSON(http.StatusOK, nil)
	}
}

// PostStudentInfo edits the information of a specific student if his ID exists in the database; otherwise, it creates a new student with the provided data.
//
// Input: Student Data (ID Optional).
//
// Output: Post Response
//
// Example URL: http://localhost:8080/api/v1/student/info
func PostStudentInfo(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var student m.Student
	var post m.PostResponse
	c.Bind(&student)
	if student.FirstName == "" || student.LastName == "" || student.ClassID == "" {
		post.Code = 400
		post.Message = "Missing Parameters"
		c.JSON(http.StatusBadRequest, post)
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
		post.Code = 200
		post.Message = "Student added/updated successfully."
		c.JSON(http.StatusOK, post)
	}
}
