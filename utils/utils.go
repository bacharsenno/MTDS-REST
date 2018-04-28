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
type Payment = model.Payment
type Grade = model.Grade
type GradeSummary = model.GradeSummary

type ClassSchedule = model.ClassSchedule
type GradeWithName = model.GradeWithName
type StudentWithGrade = model.StudentWithGrade
type BasicStudent = model.BasicStudent
type StudentParentGrades = model.StudentParentGrades
type StudentGradesBySubject = model.StudentGradesBySubject

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
		parent.GET("/students/grades", GetParentStudentsGrades)
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
	pic := []string{"https://cdn4.iconfinder.com/data/icons/cool-avatars-2/190/00-17-512.png",
		"https://www.teachngo.com/images/student_avatar.jpg"}
	pic2 := []string{"https://off2class-sol5y8kuafeozy9kld6.netdna-ssl.com/wp-content/themes/stylish-child/assets/styles/images/student/teacher_first.png",
		"https://pbs.twimg.com/profile_images/490643057822273537/pMkrGQPT_400x400.jpeg"}
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
			ProfilePic:       pic2[(i+1)%2],
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
				ProfilePic:   pic[(k+1)%2],
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
	typeA := []string{"Homework", "Oral", "Quiz", "Exam"}
	typeB := []string{"Homework", "Exam"}
	t := 0
	for i := 1; i <= 200; i++ {
		for j := 1; j <= 10; j++ {
			if i%2 != 0 {
				for k := 1; k <= 20; k++ {
					if k <= 3 || (k >= 11 && k < 13) {
						t = 0
					} else if (k > 3 && k <= 5) || (k > 12 && k < 16) {
						t = 1
					} else if (k > 5 && k <= 8) || (k > 15 && k < 18) {
						t = 2
					} else if (k > 8 && k <= 10) || k > 17 {
						t = 3
					}
					grade := Grade{
						TeacherID: "T" + strconv.Itoa(j),
						StudentID: "S" + strconv.Itoa(i),
						Subject:   "SubjectName" + strconv.Itoa(j),
						Year:      time.Now().Year(),
						Date:      time.Now().AddDate(0, (k-1)/5, 0),
						Semester:  ((k - 1) / 10) + 1,
						Type:      typeA[t],
						Grade:     truncate(r1.Float64()*99+1, 1),
						Remarks:   "REMARK STUDENT S" + strconv.Itoa(i) + " BY TEACHER T" + strconv.Itoa(j),
					}
					db.Create(&grade)
				}
			} else {
				for k := 1; k <= 6; k++ {
					t := 0
					if k%3 == 0 {
						t = 1
					}
					grade := Grade{
						TeacherID: "T" + strconv.Itoa(j),
						StudentID: "S" + strconv.Itoa(i),
						Subject:   "SubjectName" + strconv.Itoa(j),
						Year:      time.Now().Year(),
						Date:      time.Now().AddDate(0, k-1, 0),
						Semester:  (k-1)/3 + 1,
						Type:      typeB[t],
						Grade:     truncate(r1.Float64()*99+1, 1),
						Remarks:   "REMARK STUDENT S" + strconv.Itoa(i) + " BY TEACHER T" + strconv.Itoa(j),
					}
					db.Create(&grade)
				}
			}
		}
	}
	for i := 1; i <= 200; i++ {
		for j := 1; j <= 10; j++ {
			for k := 1; k <= 2; k++ {
				gradeSummary := GradeSummary{
					TeacherID: "T" + strconv.Itoa(j),
					StudentID: "S" + strconv.Itoa(i),
					Subject:   "SubjectName" + strconv.Itoa(j),
					Year:      time.Now().Year(),
					Date:      time.Now().AddDate(0, (k-1)*8, 0),
					Semester:  k,
					Grade:     truncate(r1.Float64()*99+1, 1),
					Remarks:   "REMARK STUDENT S" + strconv.Itoa(i) + " BY TEACHER T" + strconv.Itoa(j),
				}
				db.Create(&gradeSummary)
			}
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
