// Package parent provides implementation for various parent-related methods.
package parent

import (
	m "MTDS-REST/model"

	"net/http"
	"strconv"
	s "strings"
	"time"

	"github.com/gin-gonic/gin"
)

var initDb = m.InitDb

// GetParentInfo return the information of a specific parent.
//
// Input: Parent ID
//
// Output: Parent Object
func GetParentInfo(c *gin.Context) {
	db := initDb()
	defer db.Close()
	username := c.Query("id")
	var parent m.Parent
	db.Where("username = ?", username).Find(&parent)
	c.JSON(http.StatusOK, parent)
}

// GetParentNotifications returns the notification that have this specific parent, "PARENTS", or "ALL" as destination.
//
// Input: Parent ID
//
// Output: []Notification
func GetParentNotifications(c *gin.Context) {
	db := initDb()
	defer db.Close()
	username := c.Query("id")
	var notifications []m.Notification
	db.Where("destination_id in ('ALL', 'PARENTS', ?) AND start_date < ? AND end_date > ?", username, time.Now(), time.Now()).Find(&notifications)
	c.JSON(http.StatusOK, notifications)
}

// GetParentAppointments returns the scheduled appointments for a specific parent. The scope of the request can be specified (day/week).
//
// Input: Teacher ID, [scope=day/week, default week]
//
// Output: []Appointment
func GetParentAppointments(c *gin.Context) {
	db := initDb()
	defer db.Close()
	username := c.Query("id")
	scope := c.Query("scope")
	if scope == "" {
		scope = "week"
	}
	var appointments []m.Appointment
	switch scope {
	case "day":
		date := m.GetDateString(scope, 0)
		db.Where("parent_id = ? AND date(start_time) = ?", username, date).Find(&appointments)
	case "week":
		today := m.GetDateString("day", 0)
		week := m.GetDateString("day", 7)
		db.Where("parent_id = ? AND date(start_time) >= ? and date(start_time) <= ?", username, today, week).Find(&appointments)
	}
	for i := 0; i < len(appointments); i++ {
		row := db.Table("teachers t").Select("Concat(t.first_name, ' ', t.last_name) as Name").Where("t.username = ?", appointments[i].TeacherID).Row()
		row.Scan(&appointments[i].TeacherID)
	}
	c.JSON(http.StatusOK, appointments)
}

// GetParentStudents returns the students associated with a specific parent.
//
// Input: Parent ID
//
// Output: []Student
func GetParentStudents(c *gin.Context) {
	db := initDb()
	defer db.Close()
	username := c.Query("id")
	var students []m.Student
	db.Table("students s, parent_ofs po").Where("po.parent_id = ? and po.student_id = s.username", username).Find(&students)
	c.JSON(http.StatusOK, students)
}

// GetParentStudentsGrades returns the grades of the students associated with a specific parents. The grades are grouped by suject, and can be filtered by semester.
//
// Input: Parent ID, [Semester]
//
// Output: []StudentParentGrades
func GetParentStudentsGrades(c *gin.Context) {
	db := initDb()
	defer db.Close()
	id := c.Query("id")
	var students []m.Student
	var studentParentGrades []m.StudentParentGrades
	var grades []m.Grade
	var temp m.StudentParentGrades
	var temp2 m.StudentGradesBySubject
	semester := c.Query("semester")
	db.Table("students s, parent_ofs po").Where("po.parent_id = ? and po.student_id = s.username", id).Find(&students)
	for i := 0; i < len(students); i++ {
		temp.BasicStudent.FirstName = students[i].FirstName
		temp.BasicStudent.LastName = students[i].LastName
		temp.BasicStudent.ProfilePic = students[i].ProfilePic
		if semester != "" {
			sem, _ := strconv.Atoi(semester)
			db.Where("student_id = ? and semester = ?", students[i].Username, sem).Order("length(subject), subject").Find(&grades)
		} else {
			db.Where("student_id = ?", students[i].Username).Order("length(subject), subject").Find(&grades)
		}
		sub := grades[0].Subject
		temp2.Subject = sub
		for j := 0; j < len(grades); j++ {
			if grades[j].Subject == sub {
				temp2.Grades = append(temp2.Grades, grades[j])
				if j == len(grades)-1 {
					db.Where("student_id = ? and subject = ? and year = ?", students[i].Username, sub, time.Now().Year()).Find(&temp2.GradeSummaries)
					temp.SubjectGrades = append(temp.SubjectGrades, temp2)
					continue
				}
			} else {
				db.Where("student_id = ? and subject = ? and year = ?", students[i].Username, sub, time.Now().Year()).Find(&temp2.GradeSummaries)
				temp.SubjectGrades = append(temp.SubjectGrades, temp2)
				sub = grades[j].Subject
				temp2.Subject = sub
				temp2.Grades = nil
				temp2.GradeSummaries = nil
				j--
			}
		}
		studentParentGrades = append(studentParentGrades, temp)
		temp.SubjectGrades = nil
	}
	c.JSON(http.StatusOK, studentParentGrades)
}

// GetParentPayments returns the pending payments associated with a specific parent.
//
// Input: Parent ID
//
// Output: []Payment
func GetParentPayments(c *gin.Context) {
	db := initDb()
	defer db.Close()
	username := c.Query("id")
	var payments []m.Payment
	db.Where("parent_id = ?", username).Find(&payments)
	c.JSON(http.StatusOK, payments)
}

// PostParentInfo updates the information of a specified parent, or creates a new parent with the given information otherwise.
//
// Input: Teacher Data (ID Optional)
//
// Output: Newly created/edited student.
func PostParentInfo(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var parent m.Parent
	c.Bind(&parent)
	if parent.Username == "" {
		var lastParent m.Parent
		db.Limit(1).Order("LENGTH(username) desc, username desc").Find(&lastParent)
		id := lastParent.Username
		id = s.Trim(id, "P")
		num, _ := strconv.Atoi(id)
		num++
		parent.Username = "P" + strconv.Itoa(num)
		user := m.User{
			Username: parent.Username,
			Password: "PP" + strconv.Itoa(num),
			Type:     2,
		}
		db.Save(&user)
	}
	db.Save(&parent)
	c.JSON(http.StatusOK, parent)
}

// PostParentAppointment creates a new appointment between a parent and a teacher in the database.
//
// Input: Appointment
//
// Output: Appointment
func PostParentAppointment(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var appointment m.Appointment
	c.Bind(&appointment)
	db.Save(&appointment)
	c.JSON(http.StatusOK, appointment)
}

// PostParentPayment updates the payment details of specified payment associated with the specified parent in the database.
//
// Input: Payment
//
// Output: Payment
func PostParentPayment(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var payment m.Payment
	c.Bind(&payment)
	db.Save(&payment)
	c.JSON(http.StatusOK, payment)
}