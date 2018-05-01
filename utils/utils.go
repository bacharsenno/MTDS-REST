// Package utils provides the implementations for setting up the routes, generating test data, plus a basic implementation of a login function.
package utils

import (
	m "MTDS-REST/model"
	p "MTDS-REST/parent"
	d "MTDS-REST/student"
	t "MTDS-REST/teacher"

	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// R is the default Gin router
var R = gin.Default()

var InitDb = m.InitDb

// Cors is the function that does the handling of the headers (allowed origin, headers, methods etc...)
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

// SetupRoutes is a function that sets up the different API routes and specify the corresponding implementations.
//
// Paths are divided based on category (login, teacher, parent, student, class etc...)
func SetupRoutes() {

	R.Use(Cors())

	login := R.Group("api/v1/login")
	{
		login.POST("/", PostLogin)
	}

	teacher := R.Group("api/v1/teacher")
	{
		teacher.GET("/info", t.GetTeacherInfo)
		teacher.GET("/notifications", t.GetTeacherNotifications)
		teacher.GET("/appointments", t.GetTeacherAppointments)
		teacher.GET("/agenda", t.GetTeacherAgenda)
		teacher.GET("/classes", t.GetTeacherClasses)
		teacher.GET("/grades", t.GetTeacherClassGrades)
		teacher.POST("/grades", t.PostTeacherClassGrades)
		teacher.POST("/info", t.PostTeacherInfo)
		teacher.POST("/appointments", t.PostTeacherAppointment)
	}

	parent := R.Group("api/v1/parent")
	{
		parent.GET("/info", p.GetParentInfo)
		parent.GET("/notifications", p.GetParentNotifications)
		parent.GET("/appointments", p.GetParentAppointments)
		parent.GET("/students", p.GetParentStudents)
		parent.GET("/students/grades", p.GetParentStudentsGrades)
		parent.GET("/payments", p.GetParentPayments)
		parent.POST("/info", p.PostParentInfo)
		parent.POST("/appointments", p.PostParentAppointment)
		parent.POST("/payments", p.PostParentPayment)
	}

	class := R.Group("api/v1/class")
	{
		class.GET("/students", d.GetClassStudents)
	}

	student := R.Group("api/v1/student")
	{
		student.GET("/info", d.GetStudentInfo)
		student.GET("/grades", d.GetStudentGrades)
		student.POST("/info", d.PostStudentInfo)
	}

	test := R.Group("api/v1/test")
	{
		test.GET("/", GenerateTestData)
	}
}

// GenerateTestData is an automated data-generation function that generates 10 teachers, 100 parents, 200 students, classes, schedules,
// appointments, payments, and various other components for testing and visualization purposes. The data is stored in a MySQL database.
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
		teacher := m.Teacher{
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
		user := m.User{
			Username: "T" + strconv.Itoa(i),
			Password: "TP" + strconv.Itoa(i),
			Type:     1,
		}
		db.Create(&teacher)
		db.Create(&user)
	}
	for i := 1; i <= 100; i++ {
		parent := m.Parent{
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
		user := m.User{
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
			student := m.Student{
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
			teachClass := m.TeachClass{
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
			parentOf := m.ParentOf{
				StudentID:    "S" + strconv.Itoa(k),
				ParentID:     "P" + strconv.Itoa(i),
				Relationship: strconv.Itoa(j),
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
		notification := m.Notification{
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
	for k := 1; k <= 2; k++ {
		for j := 1; j <= 10; j++ {
			for i := 1; i <= 5; i++ {
				start, _ := time.Parse(time.RFC3339, "2018-01-01T"+strconv.Itoa(i+j+6)+":00:00Z")
				end, _ := time.Parse(time.RFC3339, "2018-01-01T"+strconv.Itoa(i+j+7)+":00:00Z")
				schedule := m.Schedule{
					ScheduleID: 721 + j,
					Day:        days[i-1],
					StartTime:  start,
					EndTime:    end,
					Semester:   k,
				}
				db.Create(&schedule)
			}
		}
	}
	for i := 1; i <= 4; i++ {
		appointment := m.Appointment{
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
	appointment := m.Appointment{
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
	a := 1
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 2; j++ {
			payment := m.Payment{
				PaymentID:   "PID" + strconv.Itoa(a),
				ParentID:    "P" + strconv.Itoa(i),
				StudentID:   "S" + strconv.Itoa(a),
				Amount:      truncate(r1.Float64()*1000+1000, 2),
				Deadline:    time.Now().AddDate(0, 1, 0),
				Status:      strconv.Itoa(j),
				Description: "PayDesc" + strconv.Itoa(a),
			}
			db.Create(&payment)
			a++
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
					grade := m.Grade{
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
					grade := m.Grade{
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
				gradeSummary := m.GradeSummary{
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

// PostLogin is the function that handles basic login. It returns a User object if login is successful, 404 otherwise.
func PostLogin(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var user m.User
	c.Bind(&user)
	username := user.Username
	password := user.Password
	var dbUser m.User
	db.Where("username = ? and password = ?", username, password).Find(&dbUser)
	if dbUser.Username != "" {
		dbUser.Password = ""
		c.JSON(http.StatusOK, dbUser)
	} else {
		c.String(http.StatusNotFound, "User Not Found")
	}
}

// truncate simply truncates the decimal part of a number down to n decimal places.
func truncate(x float64, n int) float64 {
	return math.Floor(x*math.Pow(10, float64(n))) * math.Pow(10, -float64(n))
}
