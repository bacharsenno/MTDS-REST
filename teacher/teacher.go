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
)

var initDb = m.InitDb

// GetTeacherInfo return the information of a specific teacher.
//
// Input: Teacher ID
//
// Output: Teacher Object
//
// Example URL: http://localhost:8080/api/v1/teacher/info?id=T1
func GetTeacherInfo(c *gin.Context) {
	db := initDb()
	defer db.Close()
	username := c.Query("id")
	var teacher m.Teacher
	db.Where("username = ?", username).First(&teacher)
	if teacher.Username == "" || teacher.FirstName == "" {
		c.JSON(http.StatusBadRequest, nil)
	} else {
		c.JSON(http.StatusOK, teacher)
	}
}

// GetTeacherNotifications returns the notifications that have this specific teacher, "TEACHERS", or "ALL" as destination.
//
// Input: Teacher ID
//
// Output: []Notification
//
// Example URL: http://localhost:8080/api/v1/teacher/notifications?id=T1
func GetTeacherNotifications(c *gin.Context) {
	db := initDb()
	defer db.Close()
	username := c.Query("id")
	var notifications []m.Notification
	var teachClasses []m.TeachClass
	var classes []string
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
}

// GetTeacherAppointments returns the scheduled appointments for a specific teacher. The scope of the request can be specified (day/week).
//
// Input:  Teacher ID, [Scope=day/week/all, default all]
//
// Output: []Appointment
//
// Example URL: http://localhost:8080/api/v1/teacher/appointments?id=T1
func GetTeacherAppointments(c *gin.Context) {
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
		db.Where("teacher_id = ? AND date(start_time) = ?", username, date).Find(&appointments)
	case "week":
		today := m.GetDateString(0)
		week := m.GetDateString(7)
		db.Where("teacher_id = ? AND date(start_time) >= ? and date(start_time) <= ?", username, today, week).Find(&appointments)
	case "all":
		db.Where("teacher_id = ?", username).Find(&appointments)
	}
	if len(appointments) > 0 {
		for i := 0; i < len(appointments); i++ {
			row := db.Table("parents p").Select("Concat(p.first_name, ' ', p.last_name) as Name").Where("p.username = ?", appointments[i].ParentID).Row()
			row.Scan(&appointments[i].ParentID)
		}
		c.JSON(http.StatusOK, appointments)
	} else {
		c.JSON(http.StatusOK, make([]string, 0))
	}
}

// GetTeacherAgenda returns the schedule of a specific teacher (i.e. time, location other info of lessons). The scope of the request can be specified (day/week).
//
// Input: Teacher ID, [Class ID], [Scope=day/week, default week], [Semester]
//
// Output: []ClassSchedule
//
// Example URL: http://localhost:8080/api/v1/teacher/agenda?id=T1&scope=day&semester=2&class=C3
func GetTeacherAgenda(c *gin.Context) {
	db := initDb()
	defer db.Close()
	username := c.Query("id")
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
			temp := m.ClassSchedule{
				TeachClass: teachClasses[i],
				Time:       schedules,
			}
			classSchedules = append(classSchedules, temp)
		}
		c.JSON(http.StatusOK, classSchedules)
	} else {
		c.JSON(http.StatusOK, make([]string, 0))
	}
}

// GetTeacherClasses returns the classes taught by a specific teacher.
//
// Input: Teacher ID, [Class ID]
//
// Output: []Classes
//
// Example URL: http://localhost:8080/api/v1/teacher/classes?id=T1
func GetTeacherClasses(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var classes []m.TeachClass
	username := c.Query("id")
	class := c.Query("class")
	if class != "" {
		db.Where("teacher_id = ? AND class_id = ?", username, class).Find(&classes)
	} else {
		db.Where("teacher_id = ?", username).Order("LENGTH(class_id), class_id").Find(&classes)
	}
	if len(classes) > 0 {
		c.JSON(http.StatusOK, classes)
	} else {
		c.JSON(http.StatusOK, make([]string, 0))
	}
}

// GetTeacherClassGrades returns the grades of the students in a specific class for a specific teacher. Semester-based filtering available.
//
// Input: TeacherID, Class ID, [Semester]
//
// Output: []StudentWithGrades
//
// Example URL: http://localhost:8080/api/v1/teacher/grades?id=T1&class=C3&semester=2
func GetTeacherClassGrades(c *gin.Context) {
	db := initDb()
	defer db.Close()
	id := c.Query("id")
	class := c.Query("class")
	var studentsWithGrades []m.StudentWithGrade
	var swgtemp m.StudentWithGrade
	var classStudents []m.Student
	db.Where("class_id = ?", class).Order("LENGTH(username), username").Find(&classStudents)
	if len(classStudents) > 0 {
		for i := 0; i < len(classStudents); i++ {
			swgtemp.BasicStudent.StudentID = classStudents[i].Username
			swgtemp.BasicStudent.FirstName = classStudents[i].FirstName
			swgtemp.BasicStudent.LastName = classStudents[i].LastName
			swgtemp.BasicStudent.ProfilePic = classStudents[i].ProfilePic
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
}

// PostTeacherClassGrades saves the grades provided by a teacher in the database.
//
// Input: []Grades
//
// Output: Post Response
//
// Example URL: http://localhost:8080/api/v1/teacher/grades
func PostTeacherClassGrades(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var gradesList m.GradesList
	var post m.PostResponse
	c.Bind(&gradesList)
	grades := gradesList.Grades
	for i := 0; i < len(grades); i++ {
		db.Save(&grades[i])
	}
	post.Code = 200
	post.Message = "Grades created/updated successfully."
	c.JSON(http.StatusOK, post)
}

// PostTeacherInfo updates the information of a specified teacher, or creates a new teacher with the given information otherwise.
//
// Input: Teacher Data (ID Optional)
//
// Output: Post Response
//
// Example URL: http://localhost:8080/api/v1/teacher/info
func PostTeacherInfo(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var teacher m.Teacher
	var post m.PostResponse
	c.Bind(&teacher)
	if teacher.FirstName != "" && teacher.LastName != "" && teacher.ProfilePic != "" {
		if teacher.Username == "" {
			var lastTeacher m.Teacher
			db.Limit(1).Order("LENGTH(username) desc, username desc").Find(&lastTeacher)
			id := lastTeacher.Username
			id = s.Trim(id, "T")
			num, _ := strconv.Atoi(id)
			num++
			teacher.Username = "T" + strconv.Itoa(num)
			user := m.User{
				Username: teacher.Username,
				Password: "TP" + strconv.Itoa(num),
				Type:     1,
			}
			db.Save(&user)
		}
		db.Save(&teacher)
		post.Code = 200
		post.Message = "Teacher created/updated successfully."
		c.JSON(http.StatusOK, post)
	} else {
		post.Code = 400
		post.Message = "Missing Parameters"
		c.JSON(http.StatusBadRequest, post)
	}
}

// PostAppointmentInfo updates an appointment between a teacher and a parent in the database.
//
// Input: Appointment
//
// Output: Post Response
//
// Example URL: http://localhost:8080/api/v1/teacher/appointments
func PostAppointmentInfo(c *gin.Context) {
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

// PostTeacherAppointment creates a new appointment between a teacher and a parent in the database.
//
// Input: Appointment
//
// Output: Post Response
//
// Example URL: http://localhost:8080/api/v1/teacher/appointment
func PostTeacherAppointment(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var appointmentReq m.AppointmentRequest
	c.Bind(&appointmentReq)
	if appointmentReq.ParentID == "" {
		var parents []m.ParentOf
		db.Where("student_id = ? AND Status = 1", appointmentReq.StudentID).Find(&parents)
		for i := 0; i < len(parents); i++ {
			app := m.Appointment{
				TeacherID:     appointmentReq.TeacherID,
				ParentID:      parents[i].ParentID,
				FullDay:       appointmentReq.FullDay,
				StartTime:     appointmentReq.StartTime,
				EndTime:       appointmentReq.EndTime,
				Remarks:       appointmentReq.Remarks,
				Status:        0,
				StatusTeacher: 1,
				StatusParent:  0,
			}
			db.Save(&app)
		}
	} else {
		app := m.Appointment{
			TeacherID:     appointmentReq.TeacherID,
			ParentID:      appointmentReq.ParentID,
			FullDay:       appointmentReq.FullDay,
			StartTime:     appointmentReq.StartTime,
			EndTime:       appointmentReq.EndTime,
			Remarks:       appointmentReq.Remarks,
			Status:        0,
			StatusTeacher: 1,
			StatusParent:  0,
		}
		db.Save(&app)
	}
	c.String(http.StatusOK, "Appointment created/updated successfully.")
}
