package utils

import (
	"net/http"
	"strconv"
	s "strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTeacherInfo(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	username := c.Query("id")
	var teacher Teacher
	db.Where("username = ?", username).First(&teacher)
	c.JSON(http.StatusOK, teacher)
}

func GetTeacherNotifications(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	username := c.Query("id")
	var notifications []Notification
	db.Where("destination_id in ('ALL', 'TEACHERS', ?) AND start_date < ? AND end_date > ?", username, time.Now(), time.Now()).Find(&notifications)
	c.JSON(http.StatusOK, notifications)
}

func GetTeacherAppointments(c *gin.Context) {
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
		db.Where("teacher_id = ? AND date(start_time) = ?", username, date).Find(&appointments)
	case "week":
		today := getDateString("day", 0)
		week := getDateString("day", 7)
		db.Where("teacher_id = ? AND date(start_time) >= ? and date(start_time) <= ?", username, today, week).Find(&appointments)
	}
	for i := 0; i < len(appointments); i++ {
		row := db.Table("parents p").Select("Concat(p.first_name, ' ', p.last_name) as Name").Where("p.username = ?", appointments[i].ParentID).Row()
		row.Scan(&appointments[i].ParentID)
	}
	c.JSON(http.StatusOK, appointments)
}

func GetTeacherAgenda(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	username := c.Query("id")
	class := c.Query("class")
	scope := c.Query("scope")
	if scope == "" {
		scope = "week"
	}
	var teachClasses []TeachClass
	var schedules []Schedule
	var classSchedules []ClassSchedule
	currentDay := time.Now().Weekday().String()
	condition := ""
	if class != "" {
		condition = " and tc.class_id = '" + class + "'"
	}
	if scope == "day" {
		db.Table("teach_classes tc, schedules s").Where("tc.teacher_id = ? and tc.schedule_id = s.schedule_id and s.day = ?"+condition, username, currentDay).Find(&teachClasses)
	}
	if scope == "week" {
		db.Table("teach_classes tc").Where("tc.teacher_id = ?"+condition, username).Find(&teachClasses)
	}
	for i := 0; i < len(teachClasses); i++ {
		if scope == "day" {
			db.Where("schedule_id = ? and Day = ?", teachClasses[i].ScheduleID, currentDay).Find(&schedules)
		} else if scope == "week" {
			db.Where("schedule_id = ? ", teachClasses[i].ScheduleID).Find(&schedules)
		}
		temp := ClassSchedule{
			TeachClass: teachClasses[i],
			Time:       schedules,
		}
		classSchedules = append(classSchedules, temp)
	}
	c.JSON(http.StatusOK, classSchedules)
}

func GetTeacherClasses(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var classes []TeachClass
	username := c.Query("id")
	class := c.Query("class")
	if class != "" {
		db.Where("teacher_id = ? AND class_id = ?", username, class).Find(&classes)
	} else {
		db.Where("teacher_id = ?", username).Find(&classes)
	}
	c.JSON(http.StatusOK, classes)
}

func GetTeacherClassGrades(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Query("id")
	class := c.Query("class")
	subject := c.Query("subject")
	condition := ""
	if subject != "" {
		condition = " and g.subject = '" + subject + "'"
	}
	object := c.Query("object")
	if object == "" {
		object = "student"
	}
	if object == "grade" {
		var grades []Grade
		db.Table("grades g, students s").Where("g.student_id = s.username and s.class_id = ? and g.teacher_id = ? "+condition, class, id).Find(&grades)
		var gradesWithNames []GradeWithName
		var temp GradeWithName
		var firstname string
		var lastname string
		for i := 0; i < len(grades); i++ {
			row := db.Table("students s").Select("s.first_name as firstname, s.last_name as lastname").Where("s.username = ?", grades[i].StudentID).Row()
			row.Scan(&firstname, &lastname)
			temp = GradeWithName{
				Grade:     grades[i],
				FirstName: firstname,
				LastName:  lastname,
			}
			gradesWithNames = append(gradesWithNames, temp)
		}
		c.JSON(http.StatusOK, gradesWithNames)
	} else if object == "student" {
		var studentsWithGrades []StudentWithGrade
		var swgtemp StudentWithGrade
		var classStudents []Student
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
}

func PostTeacherClassGrades(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var grades []Grade
	c.Bind(&grades)
	for i := 0; i < len(grades); i++ {
		db.Create(&grades[i])
	}
	c.JSON(http.StatusOK, grades)
}

func PostTeacherInfo(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var teacher Teacher
	c.Bind(&teacher)
	if teacher.Username == "" {
		var lastTeacher Teacher
		db.Limit(1).Order("LENGTH(username) desc, username desc").Find(&lastTeacher)
		id := lastTeacher.Username
		id = s.Trim(id, "T")
		num, _ := strconv.Atoi(id)
		num++
		teacher.Username = "T" + strconv.Itoa(num)
	}
	db.Save(&teacher)
	c.JSON(http.StatusOK, teacher)
}

func PostTeacherAppointment(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var appointment Appointment
	c.Bind(&appointment)
	db.Save(&appointment)
	c.JSON(http.StatusOK, appointment)
}
