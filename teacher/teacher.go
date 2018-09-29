// Package teacher provides implementation for various teacher-related methods.
package teacher

import (
	m "MTDS-REST/model"
	"fmt"

	"net/http"
	"strconv"
	s "strings"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var initDb = m.InitDb

// GetTeacherInfo return the information of a specific teacher.
//
// Input: Teacher ID
//
// Output: Teacher Object
//
// Example URL: http://localhost:8080/api/v1/teacher/T1/info
func GetTeacherInfo(c *gin.Context) {
	db := initDb()
	defer db.Close()
	id := c.Params.ByName("tid")

	var teacher m.Teacher

	if m.IsAuthorized(c, db, id) {
		db.Where("username = ?", id).First(&teacher)
		c.JSON(http.StatusOK, teacher)
	} else {
		c.JSON(http.StatusUnauthorized, m.UNAUTHORIZED_RESPONSE)
	}
}

// GetTeacherNotifications returns the notifications that have this specific teacher, "TEACHERS", or "ALL" as destination.
//
// Input: Teacher ID
//
// Output: []Notification
//
// Example URL: http://localhost:8080/api/v1/teacher/T1/notifications
func GetTeacherNotifications(c *gin.Context) {
	db := initDb()
	defer db.Close()
	username := c.Params.ByName("tid")
	var notifications []m.Notification
	var teachClasses []m.TeachClass
	var classes []string

	if m.IsAuthorized(c, db, username) {
		db.Where("teacher_id = ?", username).Find(&teachClasses)
		for i := 0; i < len(teachClasses); i++ {
			classes = append(classes, teachClasses[i].ClassID)
		}
		classesString := s.Join(classes, "','")
		fmt.Println(classesString)
		db.Where("destination_id in ('ALL', 'TEACHERS', ?, ?) AND start_date < ? AND end_date > ?", username, classesString, time.Now(), time.Now()).Find(&notifications)
		if len(notifications) > 0 {
			c.JSON(http.StatusOK, notifications)
		} else {
			c.JSON(http.StatusOK, make([]string, 0))
		}
	} else {
		c.JSON(http.StatusUnauthorized, m.UNAUTHORIZED_RESPONSE)
	}
}

// GetTeacherAppointments returns the scheduled appointments for a specific teacher. The scope of the request can be specified (day/week).
//
// Input:  Teacher ID, [Scope=day/week/all, default all]
//
// Output: []Appointment
//
// Example URL: http://localhost:8080/api/v1/teacher/T1/appointments
func GetTeacherAppointments(c *gin.Context) {
	db := initDb()
	defer db.Close()
	username := c.Params.ByName("tid")
	scope := c.Query("scope")
	if scope == "" {
		scope = "all"
	}
	var appointments []m.Appointment

	if m.IsAuthorized(c, db, username) {
		switch scope {
		case "day":
			date := m.GetDateString(0)
			db.Where("teacher_id = ? AND date(start_time) = ?", username, date).Find(&appointments)
		case "week":
			today := m.GetDateString(0)
			week := m.GetDateString(7)
			db.Where("teacher_id = ? AND date(start_time) >= ? and date(start_time) <= ?", username, today, week).Find(&appointments)
		case "all":
			db.Where("teacher_id = ?", username).Find(&appointments)
		}
		if len(appointments) > 0 {
			var objectsWithLink []m.AppointmentWithLink
			var tempObj m.AppointmentWithLink
			for i := 0; i < len(appointments); i++ {
				tempObj.Appointment = appointments[i]
				tempObj.Link = "http://localhost:8080/api/v1/parent/" + appointments[i].ParentID + "/info"
				objectsWithLink = append(objectsWithLink, tempObj)
			}
			c.JSON(http.StatusOK, objectsWithLink)
		} else {
			c.JSON(http.StatusOK, make([]string, 0))
		}
	} else {
		c.JSON(http.StatusUnauthorized, m.UNAUTHORIZED_RESPONSE)
	}
}

// GetTeacherAgenda returns the schedule of a specific teacher (i.e. time, location other info of lessons). The scope of the request can be specified (day/week).
//
// Input: Teacher ID, [Class ID], [Scope=day/week, default week], [Semester]
//
// Output: []ClassSchedule
//
// Example URL: http://localhost:8080/api/v1/teacher/T1/agenda?semester=2&class=C3
func GetTeacherAgenda(c *gin.Context) {
	db := initDb()
	defer db.Close()
	username := c.Params.ByName("tid")
	class := c.Query("class")
	scope := c.Query("scope")
	if scope == "" {
		scope = "week"
	}
	semester := c.Query("semester")
	if semester == "" {
		semester = "1"
	}
	var teachClasses []m.TeachClass
	var schedules []m.Schedule
	var classSchedules []m.ClassSchedule
	var temptc m.TeachClassWithLink

	if m.IsAuthorized(c, db, username) {
		currentDay := time.Now().Weekday().String()
		condition := ""
		if class != "" {
			condition = " and tc.class_id = '" + class + "'"
		}
		if scope == "day" {
			db.Table("teach_classes tc, schedules s").Where("tc.teacher_id = ? and tc.schedule_id = s.schedule_id and s.semester = ? and s.day = ?"+condition, username, semester, currentDay).Order("LENGTH(tc.class_id), tc.class_id").Find(&teachClasses)
		}
		if scope == "week" {
			db.Table("teach_classes tc").Where("tc.teacher_id = ?"+condition, username).Order("LENGTH(tc.class_id), tc.class_id").Find(&teachClasses)
		}
		if len(teachClasses) > 0 {
			for i := 0; i < len(teachClasses); i++ {
				if scope == "day" {
					db.Where("schedule_id = ? and day = ? and semester = ?", teachClasses[i].ScheduleID, currentDay, semester).Find(&schedules)
				} else if scope == "week" {
					db.Where("schedule_id = ? and semester = ?", teachClasses[i].ScheduleID, semester).Find(&schedules)
				}
				temptc.TeachClass = teachClasses[i]
				temptc.Link = "http://localhost:8080/api/v1/teacher/" + teachClasses[i].TeacherID + "/classes?class=" + teachClasses[i].ClassID
				temp := m.ClassSchedule{
					TeachClassWithLink: temptc,
					Time:               schedules,
				}
				classSchedules = append(classSchedules, temp)
			}
			c.JSON(http.StatusOK, classSchedules)
		} else {
			c.JSON(http.StatusOK, make([]string, 0))
		}
	} else {
		c.JSON(http.StatusUnauthorized, m.UNAUTHORIZED_RESPONSE)
	}
}

// GetTeacherClasses returns the classes taught by a specific teacher.
//
// Input: Teacher ID, [Class ID]
//
// Output: []TeachClasses
//
// Example URL: http://localhost:8080/api/v1/teacher/T1/classes
func GetTeacherClasses(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var classes []m.TeachClass
	var tcWithLink []m.TeachClassWithLink
	var tmp m.TeachClassWithLink
	username := c.Params.ByName("tid")
	class := c.Query("class")

	if m.IsAuthorized(c, db, username) {
		if class != "" {
			db.Where("teacher_id = ? AND class_id = ?", username, class).Find(&classes)
		} else {
			db.Where("teacher_id = ?", username).Order("LENGTH(class_id), class_id").Find(&classes)
		}
		if len(classes) > 0 {
			for i := 0; i < len(classes); i++ {
				tmp.TeachClass = classes[i]
				tmp.Link = "http://localhost:8080/api/v1/teacher/" + username + "/agenda?class=" + classes[i].ClassID
				tcWithLink = append(tcWithLink, tmp)
			}
			c.JSON(http.StatusOK, tcWithLink)
		} else {
			c.JSON(http.StatusOK, make([]string, 0))
		}
	} else {
		c.JSON(http.StatusUnauthorized, m.UNAUTHORIZED_RESPONSE)
	}
}

// GetTeacherClassGrades returns the grades of the students in a specific class for a specific teacher. Semester-based filtering available.
//
// Input: TeacherID, Class ID, [Semester]
//
// Output: []StudentWithGrades
//
// Example URL: http://localhost:8080/api/v1/teacher/T1/classes/C1/grades?semester=2
func GetTeacherClassGrades(c *gin.Context) {
	db := initDb()
	defer db.Close()
	id := c.Params.ByName("tid")
	class := c.Params.ByName("cid")
	var studentsWithGrades []m.StudentWithGrade
	var swgtemp m.StudentWithGrade
	var classStudents []m.Student

	if m.IsAuthorized(c, db, id) {
		db.Where("class_id = ?", class).Order("LENGTH(username), username").Find(&classStudents)
		if len(classStudents) > 0 {
			for i := 0; i < len(classStudents); i++ {
				var parent m.ParentOf
				db.Where("student_id = ?", classStudents[i].Username).First(&parent)
				swgtemp.BasicStudent.StudentID = classStudents[i].Username
				swgtemp.BasicStudent.FirstName = classStudents[i].FirstName
				swgtemp.BasicStudent.LastName = classStudents[i].LastName
				swgtemp.BasicStudent.ProfilePic = classStudents[i].ProfilePic
				swgtemp.BasicStudent.Link = "http://localhost:8080/api/v1/student/" + classStudents[i].Username + "/info"
				swgtemp.BasicStudent.Link += " ; http://localhost:8080/api/v1/parent/" + parent.ParentID + "/info"
				semester := c.Query("semester")
				if semester == "" {
					db.Where("teacher_id = ? and student_id = ? and year = ?", id, classStudents[i].Username, time.Now().Year()).Find(&swgtemp.Grades)
				} else {
					sem, _ := strconv.Atoi(semester)
					db.Where("teacher_id = ? and student_id = ? and year = ? and semester = ?", id, classStudents[i].Username, time.Now().Year(), sem).Find(&swgtemp.Grades)
				}
				db.Where("teacher_id = ? and student_id = ? and year = ?", id, classStudents[i].Username, time.Now().Year()).Find(&swgtemp.GradeSummaries)
				studentsWithGrades = append(studentsWithGrades, swgtemp)
			}
			c.JSON(http.StatusOK, studentsWithGrades)
		} else {
			c.JSON(http.StatusOK, make([]string, 0))
		}
	} else {
		c.JSON(http.StatusUnauthorized, m.UNAUTHORIZED_RESPONSE)
	}
}

// PostTeacherClassGrades saves the grades provided by a teacher in the database.
//
// Input: []Grades
//
// Output: Post Response
//
// Example URL: http://localhost:8080/api/v1/teacher/T1/grades
func PostTeacherClassGrades(c *gin.Context) {
	db := initDb()
	defer db.Close()
	id := c.Params.ByName("tid")
	var grades []m.Grade
	var post m.PostResponse
	c.Bind(&grades)

	isAuthorized := true
	for i := 0; i < len(grades); i++ {
		if !m.IsAuthorized(c, db, grades[i].TeacherID) {
			isAuthorized = false
			c.JSON(http.StatusUnauthorized, m.UNAUTHORIZED_RESPONSE)
			return
		}
		if grades[i].StudentID == "" {
			post.Code = 400
			post.Message = "Missing Parameters"
			c.JSON(http.StatusBadRequest, post)
			return
		}
	}
	if m.IsAuthorized(c, db, id) && isAuthorized {
		for i := 0; i < len(grades); i++ {
			db.Save(&grades[i])
		}
		c.JSON(http.StatusOK, grades)
	} else {
		c.JSON(http.StatusUnauthorized, m.UNAUTHORIZED_RESPONSE)
	}

}

// PostTeacherInfo updates the information of a specified teacher, or creates a new teacher with the given information otherwise.
//
// Input: Teacher Data (ID Optional)
//
// Output: Post Response
//
// Example URL: http://localhost:8080/api/v1/teacher/T1/info
func PostTeacherInfo(c *gin.Context) {
	db := initDb()
	defer db.Close()
	id := c.Params.ByName("tid")
	var teacher m.Teacher
	var post m.PostResponse
	c.Bind(&teacher)
	if m.IsAuthorized(c, db, id) && m.IsAuthorized(c, db, teacher.Username) {
		if teacher.FirstName != "" && teacher.LastName != "" && teacher.ProfilePic != "" && teacher.DateOfBirth.String() != "0001-01-01 00:00:00 +0000 UTC" {
			if teacher.Username == "" {
				var lastTeacher m.Teacher
				db.Limit(1).Order("LENGTH(username) desc, username desc").Find(&lastTeacher)
				id := lastTeacher.Username
				id = s.Trim(id, "T")
				num, _ := strconv.Atoi(id)
				num++
				teacher.Username = "T" + strconv.Itoa(num)
				password, _ := bcrypt.GenerateFromPassword([]byte("TP"+strconv.Itoa(num)), bcrypt.DefaultCost)
				user := m.User{
					Username: teacher.Username,
					Password: string(password),
					Type:     1,
				}
				db.Save(&user)
			}
			db.Save(&teacher)
			c.JSON(http.StatusOK, teacher)
		} else {
			post.Code = 400
			post.Message = "Missing Parameters"
			c.JSON(http.StatusBadRequest, post)
		}
	} else {
		c.JSON(http.StatusUnauthorized, m.UNAUTHORIZED_RESPONSE)
	}
}

// PostAppointmentInfo updates an appointment between a teacher and a parent in the database.
//
// Input: Appointment
//
// Output: Post Response
//
// Example URL: http://localhost:8080/api/v1/teacher/T1/appointments
func PostAppointmentInfo(c *gin.Context) {
	db := initDb()
	defer db.Close()
	id := c.Params.ByName("tid")
	var appointment m.Appointment
	c.Bind(&appointment)

	if m.IsAuthorized(c, db, id) && m.IsAuthorized(c, db, appointment.TeacherID) {
		if appointment.AppointmentID == 0 {
			var lastAppointment m.Appointment
			db.Limit(1).Order("LENGTH(appointment_id) desc, appointment_id desc").Find(&lastAppointment)
			appointment.AppointmentID = lastAppointment.AppointmentID + 1
		}
		db.Save(&appointment)
		c.JSON(http.StatusOK, appointment)
	} else {
		c.JSON(http.StatusUnauthorized, m.UNAUTHORIZED_RESPONSE)
	}
}
