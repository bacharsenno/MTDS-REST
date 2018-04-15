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
type Payment = model.Payment

// type Class = model.Class

var InitDb = model.InitDb

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func SetupRoutes() {

	R.Use(Cors())

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
		parent.GET("/notifications", GetParentNotifications)
		parent.GET("/appointments", GetParentAppointments)
		parent.GET("/students", GetParentStudents)
		parent.GET("/payments", GetParentPayments)
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
			Username:         "T" + strconv.Itoa(i),
			FirstName:        "TFirstName" + strconv.Itoa(i),
			LastName:         "TLastName" + strconv.Itoa(i),
			Email:            "TEmail" + strconv.Itoa(i),
			PhoneNumber:      "TPhoneNumber" + strconv.Itoa(i),
			DateOfBirth:      "TDoB" + strconv.Itoa(i),
			PlaceOfBirth:     "TPoB" + strconv.Itoa(i),
			Nationality:      "TNationality" + strconv.Itoa(i),
			Address:          "TAddr" + strconv.Itoa(i),
			FiscalCode:       "TFiscCode" + strconv.Itoa(i),
			GradDegree:       "TGradDeg" + strconv.Itoa(i),
			GradFieldOfStudy: "TGradField" + strconv.Itoa(i),
			GradGrade:        "TGradGrade" + strconv.Itoa(i),
			GradSchool:       "TGradSchool" + strconv.Itoa(i),
			SeniorityLevel:   "TSenLevel" + strconv.Itoa(i),
			StartDate:        "TStartDate" + strconv.Itoa(i),
			EndDate:          "TEndDate" + strconv.Itoa(i),
			Status:           "1",
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
			Nationality: "PNationality" + strconv.Itoa(i),
			Address:     "PAddress" + strconv.Itoa(i),
			FiscalCode:  "PFiscCode" + strconv.Itoa(i),
			Status:      "1",
		}
		user := User{
			Username: "P" + strconv.Itoa(i),
			Password: "PP" + strconv.Itoa(i),
			Type:     2,
		}
		db.Create(&parent)
		db.Create(&user)
	}
	k := 1
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 20; j++ {
			student := Student{
				Username:     "S" + strconv.Itoa(k),
				FirstName:    "SFirstName" + strconv.Itoa(k),
				LastName:     "SLastName" + strconv.Itoa(k),
				Email:        "SEmail" + strconv.Itoa(k),
				PhoneNumber:  "SPhoneNumber" + strconv.Itoa(k),
				ClassID:      "C" + strconv.Itoa(i),
				GPA:          "80.0",
				Nationality:  "SNationality" + strconv.Itoa(k),
				DateOfBirth:  "SDoB" + strconv.Itoa(k),
				PlaceOfBirth: "SPoB" + strconv.Itoa(k),
				Address:      "SAddress" + strconv.Itoa(k),
				FiscalCode:   "SFiscCode" + strconv.Itoa(k),
				EnrolledDate: "SEnrollDate" + strconv.Itoa(k),
				EndDate:      "SEndDate" + strconv.Itoa(k),
				Status:       "1",
			}
			db.Create(&student)
			k++
		}
	}
	k = 1
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			teachClass := TeachClass{
				TeacherID:  "T" + strconv.Itoa(i),
				ClassID:    "C" + strconv.Itoa(j),
				Subject:    "SubjectName" + strconv.Itoa(i),
				ScheduleID: 721 + k,
				Location:   "R" + strconv.Itoa(j),
				Year:       "2018",
				Program:    "TCProgram" + strconv.Itoa(k),
				Book:       "TCBook" + strconv.Itoa(k),
			}
			k++
			db.Create(teachClass)
		}
	}
	k = 1
	for i := 1; i <= 100; i++ {
		for j := 1; j <= 2; j++ {
			r := "Father"
			if i%2 == 0 {
				r = "Mother"
			}
			parentOf := ParentOf{
				StudentID:    "S" + strconv.Itoa(k),
				ParentID:     "P" + strconv.Itoa(i),
				Relationship: r,
				Status:       "1",
			}
			k++
			db.Create(&parentOf)
		}
	}
	for i := 1; i <= 20; i++ {
		DestinationID := "ALL"
		if i > 10 && i <= 15 {
			DestinationID = "TEACHERS"
		}
		if i > 15 {
			DestinationID = "PARENTS"
		}
		notification := Notification{
			SenderID:      "School",
			DestinationID: DestinationID,
			Topic:         "NTOPIC" + strconv.Itoa(i),
			Title:         "NTITLE" + strconv.Itoa(i),
			Description:   "NDESCRIPTION" + strconv.Itoa(i),
			Priority:      strconv.Itoa(i),
			StartDate:     "NStartDate" + strconv.Itoa(i),
			EndDate:       "NEndDate" + strconv.Itoa(i),
			Status:        "1",
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
			Date:          getDateString("day", i-1),
			FullDay:       false,
			StartTime:     strconv.Itoa(i + 12),
			EndTime:       strconv.Itoa(i + 13),
		}
		db.Create(&appointment)
	}
	m := 1
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 2; j++ {
			payment := Payment{
				PaymentID:   "PID" + strconv.Itoa(m),
				ParentID:    "P" + strconv.Itoa(i),
				StudentID:   "S" + strconv.Itoa(m),
				Amount:      "100,000",
				Deadline:    "PayDeadline" + strconv.Itoa(m),
				CreatedOn:   "PayCreated" + strconv.Itoa(m),
				Status:      strconv.Itoa(j),
				Description: "PayDesc" + strconv.Itoa(m),
			}
			db.Create(&payment)
			m++
		}
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
		date := getDateString(scope, 0)
		db.Where("teacher_id = ? AND date = ?", username, date).Find(&appointments)
	case "week":
		date := getDateString(scope, 0)
		db.Raw("SELECT * FROM appointments WHERE (teacher_id = ? AND date in ('"+date+"'))", username).Find(&appointments)
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

func GetParentNotifications(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	username := c.Query("id")
	var notifications []Notification
	db.Where("destination_id in ('ALL', 'PARENTS', ?)", username).Find(&notifications)
	c.JSON(http.StatusOK, notifications)
}

func GetParentAppointments(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	username := c.Query("id")
	scope := c.Query("scope")
	var appointments []Appointment
	switch scope {
	case "day":
		date := getDateString(scope, 0)
		db.Where("parent_id = ? AND date = ?", username, date).Find(&appointments)
	case "week":
		date := getDateString(scope, 0)
		db.Raw("SELECT * FROM appointments WHERE (parent_id = ? AND date in ('"+date+"'))", username).Find(&appointments)
		//db.Where("teacher_id = ? AND date in (?)", username, date).Find(&appointments)
	}
	for i := 0; i < len(appointments); i++ {
		row := db.Table("teachers t").Select("Concat(t.first_name, ' ', t.last_name) as Name").Where("t.username = ?", appointments[i].TeacherID).Row()
		row.Scan(&appointments[i].ParentID)
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

func GetParentPayments(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	username := c.Query("id")
	var payments []Payment
	db.Where("parent_id = ?", username).Find(&payments)
	c.JSON(http.StatusOK, payments)
}

func getDateString(scope string, offset int) string {
	if scope == "day" {
		dateString := time.Now().AddDate(0, 0, offset).Format("02-01-2006")
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

func OptionsUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
