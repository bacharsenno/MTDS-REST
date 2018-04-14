package utils

import (
	"MTDS-REST/model"
	"net/http"
	"strconv"
	s "strings"
	"time"

	"github.com/gin-gonic/gin"
)

var R = gin.Default()

type User = model.User
type Teacher = model.Teacher
type Parent = model.Parent
type Student = model.Student
type TeachClass = model.TeachClass
type ParentOf = model.ParentOf
type Notification = model.Notification
type Schedule = model.Schedule
type Appointment = model.Appointment
type ClassSchedule = model.ClassSchedule

// type Class = model.Class

var InitDb = model.InitDb

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func SetupRoutes() {

	login := R.Group("api/v1/login")
	{
		login.POST("/", PostLogin)
	}

	teacher := R.Group("api/v1/teacher")
	{
		teacher.GET("/notifications", GetTeacherNotifications)
		teacher.GET("/appointments", GetTeacherAppointments)
		teacher.GET("/agenda", GetTeacherAgenda)
		teacher.GET("/classes", GetTeacherClasses)
	}

	parent := R.Group("api/v1/parent")
	{
		parent.POST("/", PostParent)
		parent.GET("/", GetParents)
		parent.GET("/:id", GetParent)
		parent.PUT("/:id", UpdateParent)
		parent.DELETE("/:id", DeleteParent)
	}

	class := R.Group("api/v1/class")
	{
		class.POST("/", PostClass)
		class.GET("/", GetClass)
		class.GET("/:id", GetClass)
		class.PUT("/:id", UpdateClass)
		class.DELETE("/:id", DeleteClass)
	}

	test := R.Group("api/v1/test")
	{
		test.GET("/", GenerateTestData)
	}
}

func GenerateTestData(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	for i := 1; i <= 10; i++ {
		teacher := Teacher{
			Username:    "T" + strconv.Itoa(i),
			FirstName:   "TFirstName" + strconv.Itoa(i),
			LastName:    "TLastName" + strconv.Itoa(i),
			Email:       "TEmail" + strconv.Itoa(i),
			PhoneNumber: "TPhoneNumber" + strconv.Itoa(i),
		}
		user := User{
			Username: "T" + strconv.Itoa(i),
			Password: "TP" + strconv.Itoa(i),
			Type:     1,
		}
		db.Create(&teacher)
		db.Create(&user)
	}
	for i := 1; i <= 100; i++ {
		parent := Parent{
			Username:    "P" + strconv.Itoa(i),
			FirstName:   "PFirstName" + strconv.Itoa(i),
			LastName:    "PLastName" + strconv.Itoa(i),
			Email:       "PEmail" + strconv.Itoa(i),
			PhoneNumber: "PPhoneNumber" + strconv.Itoa(i),
		}
		user := User{
			Username: "P" + strconv.Itoa(i),
			Password: "PP" + strconv.Itoa(i),
			Type:     2,
		}
		db.Create(&parent)
		db.Create(&user)
	}
	// for i := 1; i <= 10; i++ {
	// 	class := Class{
	// 		ClassID: "C" + strconv.Itoa(i),
	// 	}
	// 	db.Create(&class)
	// }
	k := 1
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 20; j++ {
			student := Student{
				Username:    "S" + strconv.Itoa(k),
				FirstName:   "SFirstName" + strconv.Itoa(k),
				LastName:    "SLastName" + strconv.Itoa(k),
				Email:       "SEmail" + strconv.Itoa(k),
				PhoneNumber: "SPhoneNumber" + strconv.Itoa(k),
				ClassID:     "C" + strconv.Itoa(i),
			}
			db.Create(&student)
			k++
		}
	}
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			teachClass := TeachClass{
				TeacherID: "T" + strconv.Itoa(i),
				ClassID:   "C" + strconv.Itoa(j),
				Location:  "R" + strconv.Itoa(j),
				Subject:   "SubjectName" + strconv.Itoa(i),
			}
			if i == 1 {
				teachClass.ScheduleID = 721 + j
			}
			db.Create(teachClass)
		}
	}
	k = 1
	for i := 1; i <= 100; i++ {
		for j := 1; j <= 2; j++ {
			parentOf := ParentOf{
				StudentID: "S" + strconv.Itoa(k),
				ParentID:  "P" + strconv.Itoa(i),
			}
			k++
			db.Create(&parentOf)
		}
	}
	for i := 1; i <= 5; i++ {
		notification := Notification{
			SenderID:      "School",
			DestinationID: "ALL",
			Topic:         "NTOPIC" + strconv.Itoa(i),
			Title:         "NTITLE" + strconv.Itoa(i),
			Description:   "NDESCRIPTION" + strconv.Itoa(i),
			Priority:      strconv.Itoa(i),
		}
		db.Create(&notification)
	}
	days := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday"}
	for j := 1; j <= 10; j++ {
		for i := 1; i <= 5; i++ {
			schedule := Schedule{
				ScheduleID: 721 + j,
				Day:        days[i-1],
				StartTime:  strconv.Itoa(i + j + 6),
				EndTime:    strconv.Itoa(i + j + 7),
			}
			db.Create(&schedule)
		}
	}
	for i := 1; i <= 5; i++ {
		appointment := Appointment{
			AppointmentID: i,
			TeacherID:     "T1",
			ParentID:      "P" + strconv.Itoa(i),
			Date:          "1" + strconv.Itoa(i+2) + "-04-2018", //hhhhhhere, need to change this to generate appointments with the correct date
			FullDay:       false,
			StartTime:     strconv.Itoa(i + 12),
			EndTime:       strconv.Itoa(i + 13),
		}
		db.Create(&appointment)
	}
}

func PostLogin(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var user User
	c.Bind(&user)
	username := user.Username
	password := user.Password
	if s.HasPrefix(username, "T") {
		var loggedUser User
		db.Where("username = ? AND password = ?", username, password).First(&loggedUser)
		if loggedUser.Username != "" {
			var teacher Teacher
			db.Where("username = ?", username).First(&teacher)
			c.JSON(http.StatusOK, teacher)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Username/Password combination not found"})
		}
	} else if s.HasPrefix(username, "P") {
		var loggedUser User
		db.Where("username = ? AND password = ?", username, password).First(&loggedUser)
		if loggedUser.Username != "" {
			var parent Parent
			db.Where("username = ?", username).First(&parent)
			//db.Joins("JOIN parent_ofs po on po.student_id = students.username").Where("po.parent_id = ?", username).Find(&students)
			c.JSON(http.StatusOK, parent)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "Username/Password combination not found"})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invalid Username"})
	}
}

func GetTeacherNotifications(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	username := c.Query("id")
	var notifications []Notification
	db.Where("destination_id in ('ALL', 'TEACHERS', ?)", username).Find(&notifications)
	c.JSON(http.StatusOK, notifications)
}

func GetTeacherAppointments(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	username := c.Query("id")
	scope := c.Query("scope")
	var appointments []Appointment
	switch scope {
	case "day":
		date := getDateString(scope)
		db.Where("teacher_id = ? AND date = ?", username, date).Find(&appointments)
	case "week":
		date := getDateString(scope)
		db.Raw("SELECT * FROM appointments WHERE (teacher_id = 'T1' AND date in ('" + date + "'))").Find(&appointments)
		//db.Where("teacher_id = ? AND date in (?)", username, date).Find(&appointments)
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
	scope := c.Query("scope")
	var teachClasses []TeachClass
	var schedules []Schedule
	var classSchedules []ClassSchedule
	currentDay := time.Now().Weekday().String()
	if scope == "day" {
		db.Table("teach_classes tc, schedules s").Where("tc.teacher_id = ? and tc.schedule_id = s.schedule_id and s.day = ?", username, currentDay).Find(&teachClasses)
	}
	if scope == "week" {
		db.Where("teacher_id = ?", username).Find(&teachClasses)
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
	db.Where("teacher_id = ?", username).Find(&classes)
	c.JSON(http.StatusOK, classes)
}

func PostTeacher(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var teacher Teacher
	c.Bind(&teacher)
	db.Create(&teacher)
	c.JSON(201, gin.H{"success": teacher})
}

func GetTeachers(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var teachers []Teacher
	db.Find(&teachers)
	c.JSON(200, teachers)
}

func GetTeacher(c *gin.Context) {
	// db := InitDb()
	// defer db.Close()
	// username := c.Params.ByName("Username")
	// var teacher Teacher
	// db.Where("username = ?", username).First(&teacher)
	// if teacher.Username != "" {
	// 	c.JSON(200, teacher)
	// } else {
	// 	c.JSON(404, gin.H{"error": "Teacher not found"})
	// }
}

func UpdateTeacher(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	username := c.Params.ByName("Username")
	var teacher Teacher
	db.Where("username = ?", username).First(&teacher)
	var newTeacher Teacher
	c.Bind(&newTeacher)
	result := Teacher{
		Username:    teacher.Username,
		FirstName:   newTeacher.FirstName,
		LastName:    newTeacher.LastName,
		Email:       newTeacher.Email,
		PhoneNumber: newTeacher.PhoneNumber,
	}
	db.Save(&result)
	c.JSON(200, gin.H{"success": result})
}

func DeleteTeacher(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	username := c.Params.ByName("Username")
	var teacher Teacher
	db.Where("username = ?", username).First(&teacher)
	if teacher.Username != "" {
		db.Delete(&teacher)
		c.JSON(200, gin.H{"success": "Teacher with Username " + teacher.Username + " deleted"})
	} else {
		c.JSON(404, gin.H{"error": "Teacher not found"})
	}
}

func getDateString(scope string) string {
	if scope == "day" {
		dateString := time.Now().Format("02-01-2006")
		return dateString
	}
	if scope == "week" {
		date := []string{time.Now().Format("02-01-2006")}
		for i := 1; i <= 6; i++ {
			date = append(date, time.Now().AddDate(0, 0, i).Format("02-01-2006"))
		}
		dateString := ""
		for i := 0; i < len(date); i++ {
			dateString += date[i] + "', '"
		}
		dateString = s.TrimSuffix(dateString, "', '")
		return dateString
	}
	return ""
}

func PostParent(c *gin.Context) {
	// db := InitDb()
	// defer db.Close()
	// var parent Parent
	// c.Bind(&parent)
	// db.Create(&parent)
	// c.JSON(201, gin.H{"success": parent})
}

func GetParents(c *gin.Context) {
	// db := InitDb()
	// defer db.Close()
	// var parents []Parent
	// db.Find(&parents)
	// c.JSON(200, parents)
}

func GetParent(c *gin.Context) {
	// db := InitDb()
	// defer db.Close()
	// id := c.Params.ByName("id")
	// var parent Parent
	// db.First(&parent, id)
	// if parent.ID != "" {
	// 	c.JSON(200, parent)
	// } else {
	// 	c.JSON(404, gin.H{"error": "Parent not found"})
	// }
}

func UpdateParent(c *gin.Context) {
	// db := InitDb()
	// defer db.Close()
	// id := c.Params.ByName("id")
	// var parent Parent
	// db.First(&parent, id)
	// var newParent Parent
	// c.Bind(&newParent)
	// result := Parent{
	// 	ID:          parent.ID,
	// 	Firstname:   newParent.Firstname,
	// 	Lastname:    newParent.Lastname,
	// 	Email:       newParent.Email,
	// 	Username:    newParent.Username,
	// 	Password:    newParent.Password,
	// 	PhoneNumber: newParent.PhoneNumber,
	// }
	// db.Save(&result)
	// c.JSON(200, gin.H{"success": result})
}

func DeleteParent(c *gin.Context) {
	// db := InitDb()
	// defer db.Close()
	// id := c.Params.ByName("id")
	// var parent Parent
	// db.First(&parent, id)
	// if parent.ID != "" {
	// 	db.Delete(&parent)
	// 	c.JSON(200, gin.H{"success": "Parent #" + id + " deleted"})
	// } else {
	// 	c.JSON(404, gin.H{"error": "Parent not found"})
	// }
}

func PostSubject(c *gin.Context) {
	// db := InitDb()
	// defer db.Close()
	// var subject Subject
	// c.Bind(&subject)
	// db.Create(&subject)
	// c.JSON(201, gin.H{"success": subject})
}

func GetSubjects(c *gin.Context) {
	// db := InitDb()
	// defer db.Close()
	// var subjects []Subject
	// db.Find(&subjects)
	// c.JSON(200, subjects)
}

func GetSubject(c *gin.Context) {
	// db := InitDb()
	// defer db.Close()
	// id := c.Params.ByName("id")
	// var subject Subject
	// db.First(&subject, id)
	// if subject.ID != 0 {
	// 	c.JSON(200, subject)
	// } else {
	// 	c.JSON(404, gin.H{"error": "Subject not found"})
	// }
}

func UpdateSubject(c *gin.Context) {
	// db := InitDb()
	// defer db.Close()
	// id := c.Params.ByName("id")
	// var subject Subject
	// db.First(&subject, id)
	// var newSubject Subject
	// c.Bind(&newSubject)
	// result := Subject{
	// 	ID:    subject.ID,
	// 	Name:  newSubject.Name,
	// 	Class: newSubject.Class,
	// }
	// db.Save(&result)
	// c.JSON(200, gin.H{"success": result})
}

func DeleteSubject(c *gin.Context) {
	// db := InitDb()
	// defer db.Close()
	// id := c.Params.ByName("id")
	// var subject Subject
	// db.First(&subject, id)
	// if subject.ID != 0 {
	// 	db.Delete(&subject)
	// 	c.JSON(200, gin.H{"success": "Subject #" + id + " deleted"})
	// } else {
	// 	c.JSON(404, gin.H{"error": "Subject not found"})
	// }
}

func PostClass(c *gin.Context) {
	// db := InitDb()
	// defer db.Close()
	// var class Class
	// c.Bind(&class)
	// db.Create(&class)
	// c.JSON(201, gin.H{"success": class})
}

func GetClasses(c *gin.Context) {
	// db := InitDb()
	// defer db.Close()
	// var classes []Class
	// db.Find(&classes)
	// c.JSON(200, classes)
}

func GetClass(c *gin.Context) {
	// db := InitDb()
	// defer db.Close()
	// id := c.Params.ByName("id")
	// var class Class
	// db.First(&class, id)
	// if class.ID != 0 {
	// 	c.JSON(200, class)
	// } else {
	// 	c.JSON(404, gin.H{"error": "Class not found"})
	// }
}

func UpdateClass(c *gin.Context) {
	// db := InitDb()
	// defer db.Close()
	// id := c.Params.ByName("id")
	// var class Class
	// db.First(&class, id)
	// var newClass Class
	// c.Bind(&newClass)
	// result := Class{
	// 	ID:       class.ID,
	// 	Name:     newClass.Name,
	// 	Location: newClass.Location,
	// }
	// db.Save(&result)
	// c.JSON(200, gin.H{"success": result})
}

func DeleteClass(c *gin.Context) {
	// db := InitDb()
	// defer db.Close()
	// id := c.Params.ByName("id")
	// var class Class
	// db.First(&class, id)
	// if class.ID != 0 {
	// 	db.Delete(&class)
	// 	c.JSON(200, gin.H{"success": "Class #" + id + " deleted"})
	// } else {
	// 	c.JSON(404, gin.H{"error": "Class not found"})
	// }
}

func OptionsUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
