package utils

import (
	"net/http"
	"strconv"
	s "strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetParentInfo(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	username := c.Query("id")
	var parent Parent
	db.Where("username = ?", username).Find(&parent)
	c.JSON(http.StatusOK, parent)
}

func GetParentNotifications(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	username := c.Query("id")
	var notifications []Notification
	db.Where("destination_id in ('ALL', 'PARENTS', ?) AND start_date < ? AND end_date > ?", username, time.Now(), time.Now()).Find(&notifications)
	c.JSON(http.StatusOK, notifications)
}

func GetParentAppointments(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	username := c.Query("id")
	scope := c.Query("scope")
	if scope == "" {
		scope = "week"
	}
	var appointments []Appointment
	switch scope {
	case "day":
		date := getDateString(scope, 0)
		db.Where("parent_id = ? AND date(start_time) = ?", username, date).Find(&appointments)
	case "week":
		today := getDateString("day", 0)
		week := getDateString("day", 7)
		db.Where("parent_id = ? AND date(start_time) >= ? and date(start_time) <= ?", username, today, week).Find(&appointments)
	}
	for i := 0; i < len(appointments); i++ {
		row := db.Table("teachers t").Select("Concat(t.first_name, ' ', t.last_name) as Name").Where("t.username = ?", appointments[i].TeacherID).Row()
		row.Scan(&appointments[i].TeacherID)
	}
	c.JSON(http.StatusOK, appointments)
}

func GetParentStudents(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	username := c.Query("id")
	var students []Student
	db.Table("students s, parent_ofs po").Where("po.parent_id = ? and po.student_id = s.username", username).Find(&students)
	c.JSON(http.StatusOK, students)
}

func GetParentStudentsGrades(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Query("id")
	var students []Student
	var studentParentGrades []StudentParentGrades
	var grades []Grade
	var temp StudentParentGrades
	var temp2 StudentGradesBySubject
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

func GetParentPayments(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	username := c.Query("id")
	var payments []Payment
	db.Where("parent_id = ?", username).Find(&payments)
	c.JSON(http.StatusOK, payments)
}

func PostParentInfo(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var parent Parent
	c.Bind(&parent)
	if parent.Username == "" {
		var lastParent Parent
		db.Limit(1).Order("LENGTH(username) desc, username desc").Find(&lastParent)
		id := lastParent.Username
		id = s.Trim(id, "P")
		num, _ := strconv.Atoi(id)
		num++
		parent.Username = "P" + strconv.Itoa(num)
	}
	db.Save(&parent)
	c.JSON(http.StatusOK, parent)
}

func PostParentAppointment(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var appointment Appointment
	c.Bind(&appointment)
	db.Save(&appointment)
	c.JSON(http.StatusOK, appointment)
}

func PostParentPayment(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var payment Payment
	c.Bind(&payment)
	db.Save(&payment)
	c.JSON(http.StatusOK, payment)
}
