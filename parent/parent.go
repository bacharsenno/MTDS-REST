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
//
// Example URL: http://localhost:8080/api/v1/parent/info?id=P1
func GetParentInfo(c *gin.Context) {
	db := initDb()
	defer db.Close()
	username := c.Query("id")
	var parent m.Parent
	db.Where("username = ?", username).Find(&parent)
	if parent.Username == "" || parent.FirstName == "" {
		c.JSON(http.StatusBadRequest, nil)
	} else {
		c.JSON(http.StatusOK, parent)
	}
}

// GetParentNotifications returns the notification that have this specific parent, "PARENTS", or "ALL" as destination.
//
// Input: Parent ID
//
// Output: []Notification
//
// Example URL: http://localhost:8080/api/v1/parent/notifications?id=P1
func GetParentNotifications(c *gin.Context) {
	db := initDb()
	defer db.Close()
	username := c.Query("id")
	var notifications []m.Notification
	var students []m.Student
	db.Table("parent_ofs po, students s").Select("s.*").Where("po.parent_id = ? and po.student_id = s.username", username).Find(&students)
	var classes []string
	for i := 0; i < len(students); i++ {
		classes = append(classes, students[i].ClassID)
	}
	classesString := s.Join(classes, "','")
	db.Where("destination_id in ('ALL', 'PARENTS', ?, ?) AND start_date < ? AND end_date > ?", username, classesString, time.Now(), time.Now()).Find(&notifications)
	if len(notifications) > 0 {
		c.JSON(http.StatusOK, notifications)
	} else {
		c.JSON(http.StatusOK, make([]string, 0))
	}
}

// GetParentAppointments returns the scheduled appointments for a specific parent. The scope of the request can be specified (day/week).
//
// Input: Parent ID, [scope=day/week/all, default all]
//
// Output: []Appointment
//
// Example URL: http://localhost:8080/api/v1/parent/appointments?id=P1
func GetParentAppointments(c *gin.Context) {
	db := initDb()
	defer db.Close()
	username := c.Query("id")
	scope := c.Query("scope")
	if scope == "" {
		scope = "all"
	}
	var appointments []m.Appointment
	switch scope {
	case "day":
		date := m.GetDateString(0)
		db.Where("parent_id = ? AND date(start_time) = ?", username, date).Find(&appointments)
	case "week":
		today := m.GetDateString(0)
		week := m.GetDateString(7)
		db.Where("parent_id = ? AND date(start_time) >= ? and date(start_time) <= ?", username, today, week).Find(&appointments)
	case "all":
		db.Where("parent_id = ?", username).Find(&appointments)
	}
	if len(appointments) > 0 {
		for i := 0; i < len(appointments); i++ {
			row := db.Table("teachers t").Select("Concat(t.first_name, ' ', t.last_name) as Name").Where("t.username = ?", appointments[i].TeacherID).Row()
			row.Scan(&appointments[i].TeacherID)
		}
		c.JSON(http.StatusOK, appointments)
	} else {
		c.JSON(http.StatusOK, make([]string, 0))
	}
}

// GetParentStudents returns the students associated with a specific parent.
//
// Input: Parent ID
//
// Output: []Student
//
// Example URL: http://localhost:8080/api/v1/parent/students?id=P1
func GetParentStudents(c *gin.Context) {
	db := initDb()
	defer db.Close()
	username := c.Query("id")
	var students []m.Student
	db.Table("students s, parent_ofs po").Where("po.parent_id = ? and po.student_id = s.username", username).Find(&students)
	if len(students) > 0 {
		c.JSON(http.StatusOK, students)
	} else {
		c.JSON(http.StatusOK, make([]string, 0))
	}
}

// GetParentStudentsGrades returns the grades of the students associated with a specific parents. The grades are grouped by suject, and can be filtered by semester.
//
// Input: Parent ID, [Semester]
//
// Output: []StudentParentGrades
//
// Example URL: http://localhost:8080/api/v1/parent/students/grades?id=P1&semester=2&studentid=S2
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
	sid := c.Query("studentid")
	if sid == "" {
		db.Table("students s, parent_ofs po").Where("po.parent_id = ? and po.student_id = s.username", id).Find(&students)
	} else {
		db.Where("username = ?", sid).Find(&students)
	}
	if len(students) > 0 {
		for i := 0; i < len(students); i++ {
			temp.BasicStudent.StudentID = students[i].Username
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
	} else {
		c.JSON(http.StatusOK, make([]string, 0))
	}
}

// GetParentPayments returns the pending payments associated with a specific parent.
//
// Input: Parent ID, [status=1|2] (Status = 1 for pending payments, 2 for completed.)
//
// Output: []Payment
//
// Example URL: http://localhost:8080/api/v1/parent/payments&id=P1&status=pending
func GetParentPayments(c *gin.Context) {
	db := initDb()
	defer db.Close()
	username := c.Query("id")
	var payments []m.Payment
	status := c.Query("status")
	if status != "" {
		db.Where("parent_id = ? and status = ?", username, status).Find(&payments)
	} else {
		db.Where("parent_id = ?", username).Find(&payments)
	}
	if len(payments) > 0 {
		c.JSON(http.StatusOK, payments)
	} else {
		c.JSON(http.StatusOK, make([]string, 0))
	}
}

// PostParentInfo updates the information of a specified parent, or creates a new parent with the given information otherwise.
//
// Input: Parent Data (ID Optional)
//
// Output: Post Response.
//
// Example URL: http://localhost:8080/api/v1/parent/info
func PostParentInfo(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var parent m.Parent
	var post m.PostResponse
	c.Bind(&parent)
	if parent.FirstName != "" && parent.LastName != "" && parent.Email != "" {
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
		post.Code = 200
		post.Message = "Parent created/updated successfully."
		c.JSON(http.StatusOK, post)
	} else {
		post.Code = 400
		post.Message = "Missing Parameters"
		c.JSON(http.StatusBadRequest, post)
	}
}

// PostParentAppointment creates a new appointment between a parent and a teacher in the database.
//
// Input: Appointment
//
// Output: Post Response
//
// Example URL: http://localhost:8080/api/v1/parent/appointments
func PostParentAppointment(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var appointment m.Appointment
	var post m.PostResponse
	c.Bind(&appointment)
	if appointment.AppointmentID == 0 {
		var lastAppointment m.Appointment
		db.Limit(1).Order("LENGTH(appointment_id) desc, appointment_id desc").Find(&lastAppointment)
		appointment.AppointmentID = lastAppointment.AppointmentID + 1
	}
	db.Save(&appointment)
	post.Code = 200
	post.Message = "Appointment created/updated successfully."
	c.JSON(http.StatusOK, post)
}

// PostParentPayment updates the payment details of specified payment associated with the specified parent in the database.
//
// Input: Payment
//
// Output: Post Response
//
// Example URL: http://localhost:8080/api/v1/parent/payments
func PostParentPayment(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var payment m.Payment
	var post m.PostResponse
	c.Bind(&payment)
	if payment.PaymentID == "" || payment.ParentID == "" {
		var lastPayment m.Payment
		db.Limit(1).Order("LENGTH(payment_id) desc, payment_id desc").Find(&lastPayment)
		pid := lastPayment.PaymentID
		num, _ := strconv.Atoi(s.Trim(pid, "PID"))
		num++
		payment.PaymentID = "PID" + strconv.Itoa(num)
	}
	db.Save(&payment)
	post.Code = 200
	post.Message = "Payment updated successfully."
	c.JSON(http.StatusOK, post)
}
