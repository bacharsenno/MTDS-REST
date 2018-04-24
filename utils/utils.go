package utils

import (
	"MTDS-REST/model"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	s "strings"
	"time"

	"github.com/gin-gonic/gin"
)

var R = gin.Default()

//Mon Jan 2 15:04:05 MST 2006

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
type Grade = model.Grade
type GradeWithName = model.GradeWithName
type StudentWithGrade = model.StudentWithGrade

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
		teacher.GET("/info", GetTeacherInfo)
		teacher.GET("/notifications", GetTeacherNotifications)
		teacher.GET("/appointments", GetTeacherAppointments)
		teacher.GET("/agenda", GetTeacherAgenda)
		teacher.GET("/classes", GetTeacherClasses)
		teacher.GET("/grades", GetTeacherClassGrades)
		teacher.POST("/grades", PostTeacherClassGrades)
		teacher.POST("/info", PostTeacherInfo)
		teacher.POST("/appointments", PostTeacherAppointment)
	}

	parent := R.Group("api/v1/parent")
	{
		parent.GET("/info", GetParentInfo)
		parent.GET("/notifications", GetParentNotifications)
		parent.GET("/appointments", GetParentAppointments)
		parent.GET("/students", GetParentStudents)
		parent.GET("/payments", GetParentPayments)
		parent.POST("/info", PostParentInfo)
		parent.POST("/appointments", PostParentAppointment)
		parent.POST("/payments", PostParentPayment)
	}

	class := R.Group("api/v1/class")
	{
		class.GET("/students", GetClassStudents)
	}

	student := R.Group("api/v1/student")
	{
		student.GET("/info", GetStudentInfo)
		student.GET("/grades", GetStudentGrades)
		student.POST("/info", PostStudentInfo)
	}

	test := R.Group("api/v1/test")
	{
		test.GET("/", GenerateTestData)
	}
}

func GenerateTestData(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := 1; i <= 10; i++ {
		Dob := time.Now()
		if i < 10 {
			date := "12-0" + strconv.Itoa(i) + "-1985"
			Dob, _ = time.Parse("02-01-2006", date)
		} else {
			date := "12-" + strconv.Itoa(i) + "-1985"
			Dob, _ = time.Parse("02-01-2006", date)
		}
		teacher := Teacher{
			Username:         "T" + strconv.Itoa(i),
			FirstName:        "TFirstName" + strconv.Itoa(i),
			LastName:         "TLastName" + strconv.Itoa(i),
			Email:            "TEmail" + strconv.Itoa(i),
			PhoneNumber:      "TPhoneNumber" + strconv.Itoa(i),
			DateOfBirth:      Dob,
			PlaceOfBirth:     "TPoB" + strconv.Itoa(i),
			Nationality:      "TNationality" + strconv.Itoa(i),
			Address:          "TAddr" + strconv.Itoa(i),
			FiscalCode:       "TFiscCode" + strconv.Itoa(i),
			GradDegree:       "TGradDeg" + strconv.Itoa(i),
			GradFieldOfStudy: "TGradField" + strconv.Itoa(i),
			GradGrade:        "TGradGrade" + strconv.Itoa(i),
			GradSchool:       "TGradSchool" + strconv.Itoa(i),
			SeniorityLevel:   "TSenLevel" + strconv.Itoa(i),
			StartDate:        time.Now().AddDate(0, -1, 0),
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
	date, _ := time.Parse("02-01-2006", "01-01-2000")
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 20; j++ {
			student := Student{
				Username:     "S" + strconv.Itoa(k),
				FirstName:    "SFirstName" + strconv.Itoa(k),
				LastName:     "SLastName" + strconv.Itoa(k),
				Email:        "SEmail" + strconv.Itoa(k),
				PhoneNumber:  "SPhoneNumber" + strconv.Itoa(k),
				ClassID:      "C" + strconv.Itoa(i),
				GPA:          truncate(r1.Float64()*99+1, 1),
				Nationality:  "SNationality" + strconv.Itoa(k),
				DateOfBirth:  date,
				PlaceOfBirth: "SPoB" + strconv.Itoa(k),
				Address:      "SAddress" + strconv.Itoa(k),
				FiscalCode:   "SFiscCode" + strconv.Itoa(k),
				EnrolledDate: date.AddDate(3, 0, 0),
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
			StartDate:     time.Now(),
			EndDate:       time.Now().AddDate(0, 1, 0),
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
	for i := 1; i <= 4; i++ {
		appointment := Appointment{
			AppointmentID: i,
			TeacherID:     "T1",
			ParentID:      "P" + strconv.Itoa(i),
			FullDay:       false,
			StartTime:     time.Now().AddDate(0, 0, i-1),
			EndTime:       time.Now().AddDate(0, 0, i-1).Add(time.Hour),
			Status:        1,
			StatusTeacher: 1,
			StatusParent:  1,
		}
		db.Create(&appointment)
	}
	appointment := Appointment{
		AppointmentID: 5,
		TeacherID:     "T1",
		ParentID:      "P" + strconv.Itoa(5),
		StartTime:     time.Now().AddDate(0, 0, 4),
		EndTime:       time.Now().AddDate(0, 0, 4),
		FullDay:       true,
		Status:        1,
		StatusTeacher: 1,
		StatusParent:  1,
	}
	db.Create(&appointment)
	m := 1
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 2; j++ {
			payment := Payment{
				PaymentID:   "PID" + strconv.Itoa(m),
				ParentID:    "P" + strconv.Itoa(i),
				StudentID:   "S" + strconv.Itoa(m),
				Amount:      truncate(r1.Float64()*1000+1000, 2),
				Deadline:    time.Now().AddDate(0, 1, 0),
				Status:      strconv.Itoa(j),
				Description: "PayDesc" + strconv.Itoa(m),
			}
			db.Create(&payment)
			m++
		}
	}
	for i := 1; i <= 200; i++ {
		for j := 1; j <= 10; j++ {
			grade := Grade{
				TeacherID: "T" + strconv.Itoa(j),
				StudentID: "S" + strconv.Itoa(i),
				Subject:   "SubjectName" + strconv.Itoa(j),
				Year:      time.Now().Year(),
				Date:      time.Now(),
				Grade:     truncate(r1.Float64()*99+1, 1),
				Remarks:   "REMARK STUDENT S" + strconv.Itoa(i) + " BY TEACHER T" + strconv.Itoa(j),
			}
			db.Create(&grade)
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
	var dbUser User
	db.Where("username = ? and password = ?", username, password).Find(&dbUser)
	if dbUser.Username != "" {
		dbUser.Password = ""
		c.JSON(http.StatusOK, dbUser)
	} else {
		c.String(http.StatusNotFound, "User Not Found")
	}
}

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
	var grades []Grade
	db.Table("grades g, students s").Where("g.student_id = s.username and s.class_id = ? and g.teacher_id = ? "+condition, class, id).Find(&grades)
	object := c.Query("object")
	if object == "" {
		object = "grade"
	}
	if object == "grade" {
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
		var temp Student
		for i := 0; i < len(grades); i++ {
			db.Where("username = ?", grades[i].StudentID).First(&temp)
			temp2 := StudentWithGrade{
				Student: temp,
				Grade:   grades[i],
			}
			studentsWithGrades = append(studentsWithGrades, temp2)
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

func GetClassStudents(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	class := c.Query("class")
	var students []Student
	db.Where("class_id = ?", class).Find(&students)
	c.JSON(http.StatusOK, students)
}

func GetStudentGrades(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Query("id")
	var grades []Grade
	db.Where("student_id = ?", id).Find(&grades)
	c.JSON(http.StatusOK, grades)
}

func GetStudentInfo(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var student Student
	id := c.Query("id")
	db.Where("username = ?", id).First(&student)
	c.JSON(http.StatusOK, student)
}

func PostStudentInfo(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var student Student
	c.Bind(&student)
	if student.Username == "" {
		var lastStudent Student
		db.Limit(1).Order("LENGTH(username) desc, username desc").Find(&lastStudent)
		id := lastStudent.Username
		id = s.Trim(id, "S")
		num, _ := strconv.Atoi(id)
		num++
		student.Username = "S" + strconv.Itoa(num)
	}
	db.Save(&student)
	c.JSON(http.StatusOK, student)
}

func getDateString(scope string, offset int) string {
	if scope == "day" {
		dateString := time.Now().AddDate(0, 0, offset).Format("2006-01-02")
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

func truncate(x float64, n int) float64 {
	return math.Floor(x*math.Pow(10, float64(n))) * math.Pow(10, -float64(n))
}

func OptionsUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
