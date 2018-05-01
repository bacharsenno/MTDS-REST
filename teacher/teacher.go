// Package teacher provides implementation for various teacher-related methods.
package teacher

import (
	m "MTDS-REST/model"

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
func GetTeacherInfo(c *gin.Context) {
	db := initDb()
	defer db.Close()
	username := c.Query("id")
	var teacher m.Teacher
	db.Where("username = ?", username).First(&teacher)
	c.JSON(http.StatusOK, teacher)
}

// GetTeacherNotifications returns the notifications that have this specific teacher, "TEACHERS", or "ALL" as destination.
//
// Input: Teacher ID
//
// Output: []Notification
func GetTeacherNotifications(c *gin.Context) {
	db := initDb()
	defer db.Close()
	username := c.Query("id")
	var notifications []m.Notification
	db.Where("destination_id in ('ALL', 'TEACHERS', ?) AND start_date < ? AND end_date > ?", username, time.Now(), time.Now()).Find(&notifications)
	c.JSON(http.StatusOK, notifications)
}

// GetTeacherAppointments returns the scheduled appointments for a specific teacher. The scope of the request can be specified (day/week).
//
// Input:  Teacher ID, [Scope=day/week, week default]
//
// Output: []Appointment

func GetTeacherAppointments(c *gin.Context) {
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
		db.Where("teacher_id = ? AND date(start_time) = ?", username, date).Find(&appointments)
	case "week":
		today := m.GetDateString("day", 0)
		week := m.GetDateString("day", 7)
		db.Where("teacher_id = ? AND date(start_time) >= ? and date(start_time) <= ?", username, today, week).Find(&appointments)
	}
	for i := 0; i < len(appointments); i++ {
		row := db.Table("parents p").Select("Concat(p.first_name, ' ', p.last_name) as Name").Where("p.username = ?", appointments[i].ParentID).Row()
		row.Scan(&appointments[i].ParentID)
	}
	c.JSON(http.StatusOK, appointments)
}

// GetTeacherAgenda returns the schedule of a specific teacher (i.e. time, location other info of lessons). The scope of the request can be specified (day/week).
//
// Input: Teacher ID, [Class ID], [Scope=day/week, default week], [Semester]
//
// Output: []ClassSchedule
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
}

// GetTeacherClasses returns the classes taught by a specific teacher.
//
// Input: Teacher ID, [Class ID]
//
// Output: []Classes
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
	c.JSON(http.StatusOK, classes)
}

// GetTeacherClassGrades returns the grades of the students in a specific class for a specific teacher. Semester-based filtering available.
//
// Input: TeacherID, Class ID
//
// Output: []StudentWithGrades
func GetTeacherClassGrades(c *gin.Context) {
	db := initDb()
	defer db.Close()
	id := c.Query("id")
	class := c.Query("class")
	var studentsWithGrades []m.StudentWithGrade
	var swgtemp m.StudentWithGrade
	var classStudents []m.Student
	db.Where("class_id = ?", class).Order("LENGTH(username), username").Find(&classStudents)
	for i := 0; i < len(classStudents); i++ {
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
}

// PostTeacherClassGrades saves the grades provided by a teacher in the database.
//
// Input: []Grades
//
// Output: []Grades
func PostTeacherClassGrades(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var grades []m.Grade
	c.Bind(&grades)
	for i := 0; i < len(grades); i++ {
		db.Create(&grades[i])
	}
	c.JSON(http.StatusOK, grades)
}

// PostTeacherInfo updates the information of a specified teacher, or creates a new teacher with the given information otherwise.
//
// Input: Teacher Data (ID Optional)
//
// Output: Newly created/edited student.
func PostTeacherInfo(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var teacher m.Teacher
	c.Bind(&teacher)
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
	c.JSON(http.StatusOK, teacher)
}

// PostTeacherAppointment creates a new appointment between a teacher and a parent in the database.
//
// Input: Appointment
//
// Output: Appointment
func PostTeacherAppointment(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var appointment m.Appointment
	c.Bind(&appointment)
	db.Save(&appointment)
	c.JSON(http.StatusOK, appointment)
}
