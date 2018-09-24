// Package model provides the definitions of all of the structs included in the project, the names of the fields piped into JSON, and initializes the MySQL database accordingly.
package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
)

// User is the struct containing the login data and type of users.
//
// Type = 0 for Administrators, 1 for Teachers, 2 for Parents.
type User struct {
	Username string `gorm:"PRIMARY_KEY" form:"Username" json:"Username"`
	Password string `gorm:"not null" form:"Password" json:"Password"`
	Type     int    `gorm:"not null" form:"Type" json:"Type"`
}

// Teacher is the basic teacher struct containing all related info. Fields are self-explanatory.
type Teacher struct {
	Username         string    `gorm:"PRIMARY_KEY" form:"Username" json:"Username"`
	FirstName        string    `gorm:"not null" form:"FirstName" json:"FirstName"`
	LastName         string    `gorm:"not null" form:"LastName" json:"LastName"`
	ProfilePic       string    `form:"ProfilePic" json:"ProfilePic"`
	Email            string    `gorm:"not null" form:"Email" json:"Email"`
	PhoneNumber      string    `gorm:"not null" form:"PhoneNumber" json:"PhoneNumber"`
	DateOfBirth      time.Time `form:"DateOfBirth" json:"DateOfBirth"`
	PlaceOfBirth     string    `form:"PlaceOfBirth" json:"PlaceOfBirth"`
	Nationality      string    `form:"Nationality" json:"Nationality"`
	Address          string    `form:"Address" json:"Address"`
	FiscalCode       string    `form:"FiscalCode" json:"FiscalCode"`
	GradDegree       string    `form:"GradDegree" json:"GradDegree"`
	GradFieldOfStudy string    `form:"GradFieldOfStudy" json:"GradFieldOfStudy"`
	GradGrade        string    `form:"GradGrade" json:"GradGrade"`
	GradSchool       string    `form:"GradSchool" json:"GradSchool"`
	SeniorityLevel   string    `form:"SeniorityLevel" json:"SeniorityLevel"`
	StartDate        time.Time `form:"StartDate" json:"StartDate"`
	EndDate          time.Time `form:"EndDate" json:"EndDate"`
	Status           string    `form:"Status" json:"Status"`
}

// Parent is the basic parent struct containing all related info. Fields are self-explanatory.
type Parent struct {
	Username    string `gorm:"PRIMARY_KEY" form:"Username" json:"Username"`
	FirstName   string `gorm:"not null" form:"FirstName" json:"FirstName"`
	LastName    string `gorm:"not null" form:"LastName" json:"LastName"`
	Email       string `gorm:"not null" form:"Email" json:"Email"`
	PhoneNumber string `gorm:"not null" form:"PhoneNumber" json:"PhoneNumber"`
	Nationality string `form:"Nationality" json:"Nationality"`
	Address     string `form:"Address" json:"Address"`
	FiscalCode  string `form:"FiscalCode" json:"FiscalCode"`
	Status      string `form:"Status" json:"Status"`
}

// Student is the basic student struct containing all related info. Fields are self-explanatory.
type Student struct {
	Username     string    `gorm:"PRIMARY_KEY" form:"Username" json:"Username"`
	FirstName    string    `gorm:"not null" form:"FirstName" json:"FirstName"`
	LastName     string    `gorm:"not null" form:"LastName" json:"LastName"`
	ProfilePic   string    `form:"ProfilePic" json:"ProfilePic"`
	Email        string    `gorm:"not null" form:"Email" json:"Email"`
	PhoneNumber  string    `gorm:"not null" form:"PhoneNumber" json:"PhoneNumber"`
	ClassID      string    `gorm:"not null" form:"ClassID" json:"ClassID"`
	GPA          float64   `gorm:"not null" form:"GPA" json:"GPA"`
	Nationality  string    `form:"Nationality" json:"Nationality"`
	DateOfBirth  time.Time `form:"DateOfBirth" json:"DateOfBirth"`
	PlaceOfBirth string    `form:"PlaceOfBirth" json:"PlaceOfBirth"`
	Address      string    `form:"Address" json:"Address"`
	FiscalCode   string    `form:"FiscalCode" json:"FiscalCode"`
	EnrolledDate time.Time `form:"EnrolledDate" json:"EnrolledDate"`
	EndDate      time.Time `form:"EndDate" json:"EndDate"`
	Status       string    `form:"Status" json:"Status"`
}

// Grade is the basic grade struct containing all related info. Fields are self-explanatory.
type Grade struct {
	TeacherID string    `form:"TeacherID" json:"TeacherID"`
	StudentID string    `form:"StudentID" json:"StudentID"`
	Subject   string    `form:"Subject" json:"Subject"`
	Year      int       `form:"Year" json:"Year"`
	Semester  int       `form:"Semester" json:"Semester"`
	Type      string    `form:"Type" json:"Type"`
	Date      time.Time `form:"Date" json:"Date"`
	Grade     float64   `form:"Grade" json:"Grade"`
	Remarks   string    `form:"Remarks" json:"Remarks"`
	Link      string    `form:"Link" json:"Link"`
}

// GradeSummary is the struct containing all related info, representing the overall performance of a student in a specific course for a specific semester. Fields are self-explanatory.
type GradeSummary struct {
	TeacherID string    `form:"TeacherID" json:"TeacherID"`
	StudentID string    `form:"StudentID" json:"StudentID"`
	Subject   string    `form:"Subject" json:"Subject"`
	Year      int       `form:"Year" json:"Year"`
	Semester  int       `form:"Semester" json:"Semester"`
	Date      time.Time `form:"Date" json:"Date"`
	Grade     float64   `form:"Grade" json:"Grade"`
	Remarks   string    `form:"Remarks" json:"Remarks"`
}

// Payment is the basic payment struct containing all related info.
//
// Status = 1 for pending payments, 2 for completed.
type Payment struct {
	PaymentID   string    `gorm:"PRIMARY_KEY" form:"PaymentID" json:"PaymentID"`
	ParentID    string    `form:"ParentID" json:"ParentID"`
	StudentID   string    `form:"StudentID" json:"StudentID"`
	Amount      float64   `form:"Amount" json:"Amount"`
	Deadline    time.Time `form:"Deadline" json:"Deadline"`
	CreatedAt   time.Time `gorm:"type:timestamp" form:"CreatedOn" json:"CreatedOn"`
	Status      string    `form:"Status" json:"Status"`
	Description string    `form:"Description" json:"Description"`
}

// Notification is the basic notification struct containing all related info.
//
// Status = 1 for active, 2 for discarded, 0 for expired.
//
// Topic = Trip, Parent-Teacher Conference, Holidays etc...
type Notification struct {
	SenderID      string    `form:"SenderID" json:"SenderID"`
	DestinationID string    `form:"DestinationID" json:"DestinationID"`
	Topic         string    `form:"Topic" json:"Topic"`
	Title         string    `form:"Title" json:"Title"`
	Description   string    `form:"Description" json:"Description"`
	Priority      string    `form:"Priority" json:"Priority"`
	StartDate     time.Time `form:"StartDate" json:"StartDate"`
	EndDate       time.Time `form:"EndDate" json:"EndDate"`
	Status        string    `form:"Status" json:"Status"`
}

// Appointment is the basic appointment struct containing all related info.
//
// Status = 1 for active, 2 for discarded, 0 for expired.
//
// StatusTeacher/StatusParent = 1 for approved, 0 for rejected.
type Appointment struct {
	AppointmentID int       `gorm:"PRIMARY_KEY;AUTO_INCREMENT" form:"AppointmentID" json:"AppointmentID"`
	TeacherID     string    `form:"TeacherID" json:"TeacherID"`
	ParentID      string    `form:"ParentID" json:"ParentID"`
	FullDay       bool      `form:"FullDay" json:"FullDay"`
	StartTime     time.Time `form:"StartTime" json:"StartTime"`
	EndTime       time.Time `form:"EndTime" json:"EndTime"`
	Remarks       string    `form:"Remarks" json:"Remarks"`
	Status        int       `form:"Status" json:"Status"`
	StatusTeacher int       `form:"StatusTeacher" json:"StatusTeacher"`
	StatusParent  int       `form:"StatusParent" json:"StatusParent"`
}

// Schedule represents the timetable of a specific subject for a specific class. Fields are self-explanatory.
type Schedule struct {
	ScheduleID int       `form:"ScheduleID" json:"ScheduleID"`
	Day        string    `form:"Day" json:"Day"`
	StartTime  time.Time `form:"StartTime" json:"StartTime"`
	EndTime    time.Time `form:"EndTime" json:"EndTime"`
	Semester   int       `form:"Semester" json:"Semester"`
}

// TeachClass contains the details of classes taught by teachers. Fields are self-explanatory.
type TeachClass struct {
	TeacherID  string `gorm:"PRIMARY_KEY" form:"TeacherID" json:"TeacherID"`
	ClassID    string `gorm:"PRIMARY_KEY" form:"ClassID" json:"ClassID"`
	Subject    string `gorm:"PRIMARY_KEY" form:"Subject" json:"Subject"`
	ScheduleID int    `form:"Schedule" json:"Schedule"`
	Location   string `form:"Location" json:"Location"`
	Year       string `form:"Year" json:"Year"`
	Program    string `form:"Program" json:"Program"`
	Book       string `form:"Book" json:"Book"`
}

// Class contains the details of classes. Fields are self-explanatory.
type Class struct {
	ClassID  string `gorm:"PRIMARY_KEY" form:"ClassID" json:"ClassID"`
	Location string `form:"Location" json:"Location"`
	Year     string `form:"Year" json:"Year"`
}

// ParentOf defines the parent-child relationships between parents and students.
//
// Relationship = 1 for Mother, 2 for Father.
//
// Status = 1 for active, 0 for removed.
type ParentOf struct {
	StudentID    string `form:"StudentID" json:"StudentID"`
	ParentID     string `form:"ParentID" json:"ParentID"`
	Relationship string `form:"Relationship" json:"Relationship"`
	Status       string `form:"Status" json:"Status"`
}

// CreditCard is an auxiliary class used for payment processing.
type CreditCard struct {
	NameOnCard string `form:"NameOnCard" json:"NameOnCard"`
	Circuit    string `form:"Circuit" json:"Circuit"`
	CCN        string `form:"CCN" json:"CCN"`
	CVV        string `form:"CVV" json:"CVV"`
	Expiry     string `form:"Expiry" json:"Expiry"`
}

// ClassSchedule is a custom struct created for JSON construction purposes.
type ClassSchedule struct {
	TeachClassWithLink `form:"TeachClassWithLink" json:"TeachClassWithLink"`
	Time               []Schedule `form:"Schedule" json:"Schedule"`
}

// BasicStudent is a custom struct created for JSON construction purposes.
type BasicStudent struct {
	StudentID  string `form:"StudentID" json:"StudentID"`
	FirstName  string `form:"FirstName" json:"FirstName"`
	LastName   string `form:"LastName" json:"LastName"`
	ProfilePic string `form:"ProfilePic" json:"ProfilePic"`
	Link       string `form:"Link" json:"Link"`
}

// StudentWithGrade is a custom struct created for JSON construction purposes.
type StudentWithGrade struct {
	BasicStudent   `form:"BasicStudent" json:"BasicStudent"`
	Grades         []Grade        `form:"Grade" json:"Grade"`
	GradeSummaries []GradeSummary `form:"GradeSummary" json:"GradeSummary"`
}

// StudentGradesBySubject is a custom struct created for JSON construction purposes.
type StudentGradesBySubject struct {
	Subject        string         `form:"Subject" json:"Subject"`
	Grades         []Grade        `form:"Grades" json:"Grades"`
	GradeSummaries []GradeSummary `form:"GradeSummaries" json:"GradeSummaries"`
}

// StudentParentGrades is a custom struct created for JSON construction purposes.
type StudentParentGrades struct {
	BasicStudent  `form:"BasicStudent" json:"BasicStudent"`
	SubjectGrades []StudentGradesBySubject `form:"SubjectGrades" json:"SubjectGrades"`
}

// GradesList is a custom struct created for JSON construction purposes.
type GradesList struct {
	Grades []Grade `form:"Grades" json:"Grades"`
}

// AppointmentRequest is a custom struct created for JSON construction purposes.
type AppointmentRequest struct {
	StudentID string    `form:"StudentID" json:"StudentID"`
	TeacherID string    `form:"TeacherID" json:"TeacherID"`
	ParentID  string    `form:"ParentID" json:"ParentID"`
	FullDay   bool      `form:"FullDay" json:"FullDay"`
	StartTime time.Time `form:"StartTime" json:"StartTime"`
	EndTime   time.Time `form:"EndTime" json:"EndTime"`
	Remarks   string    `form:"Remarks" json:"Remarks"`
}

// PaymentInfo is a custom struct created for JSON construction purposes.
type PaymentInfo struct {
	CreditCard `form:"CreditCard" json:"CreditCard"`
	Payment    `form:"Payment" json:"Payment"`
}

// AppointmentWithLink is Appointment object with deeplink
type AppointmentWithLink struct {
	Appointment `form:"Appointment" json:"Appointment"`
	Link        string `form:"Link" json:"Link"`
}

// StudentWithLink is Student object with deeplink
type StudentWithLink struct {
	Student `form:"Student" json:"Student"`
	Link    string `form:"Link" json:"Link"`
}

// ParentWithLink is Parent object with deeplink
type ParentWithLink struct {
	ParentOf `form:"Parent" json:"Parent"`
	Link     string `form:"Link" json:"Link"`
}

// TeachClassWithLink is TeachClass object with deeplink
type TeachClassWithLink struct {
	TeachClass `form:"TeachClass" json:"TeachClass"`
	Link       string `form:"Link" json:"Link"`
}

// PaymentWithLink is Payment object with deeplink
type PaymentWithLink struct {
	Payment `form:"Payment" json:"Payment"`
	Link    string `form:"Link" json:"Link"`
}

// PostResponse is the default struct returned as response for any POST request.
type PostResponse struct {
	Code    int    `form:"code" json:"code"`
	Message string `form:"message" json:"message"`
}

// InitDb creates the connection with the MySQL database and creates the needed/missing tables based on the struct definitions previously specified.
func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("mysql", "root:@/testdb?parseTime=True")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{})
	}
	if !db.HasTable(&Teacher{}) {
		db.CreateTable(&Teacher{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Teacher{})
	}
	if !db.HasTable(&Parent{}) {
		db.CreateTable(&Parent{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Parent{})
	}
	if !db.HasTable(&Student{}) {
		db.CreateTable(&Student{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Student{})
	}
	if !db.HasTable(&Schedule{}) {
		db.CreateTable(&Schedule{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Schedule{})
	}
	if !db.HasTable(&Grade{}) {
		db.CreateTable(&Grade{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Grade{})
	}
	if !db.HasTable(&GradeSummary{}) {
		db.CreateTable(&GradeSummary{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&GradeSummary{})
	}
	if !db.HasTable(&Payment{}) {
		db.CreateTable(&Payment{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Payment{})
	}
	if !db.HasTable(&Notification{}) {
		db.CreateTable(&Notification{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Notification{})
	}
	if !db.HasTable(&Appointment{}) {
		db.CreateTable(&Appointment{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Appointment{})
	}
	if !db.HasTable(&TeachClass{}) {
		db.CreateTable(&TeachClass{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&TeachClass{})
	}
	if !db.HasTable(&ParentOf{}) {
		db.CreateTable(&ParentOf{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&ParentOf{})
	}
	return db
}

// GetDateString is a utility function that generates a date string under a certain format with a possible offset for day/month/year etc...
func GetDateString(offset int) string {
	dateString := time.Now().AddDate(0, 0, offset).Format("2006-01-02")
	return dateString
}

var UNAUTHORIZED_RESPONSE = PostResponse{Code: 401, Message: "You are not authorized to access this resource."}

// isAuthorized check if the logged user is authorized to access the required resource
func IsAuthorized(c *gin.Context, db *gorm.DB, paramId string) bool {
	requestKey := c.GetHeader("X-Auth-Key");
	
	var user User
	db.Where("username = ?", requestKey).First(&user)

	//if the logged user is an Admin or its id correspond to the one on the request returns true
	if user.Type == 0 || requestKey == paramId {
		return true
	} else {
		return false
	}
}

// isAuthorized check if the logged user is authorized to access the required resource according to its type
func IsAuthorizedUserType(c *gin.Context, db *gorm.DB, userType int) bool {
	requestKey := c.GetHeader("X-Auth-Key");
	
	var user User
	db.Where("username = ?", requestKey).First(&user)

	//if the logged user is an Admin or its id correspond to the one on the request returns true
	if user.Type == 0 || user.Type == userType {
		return true
	} else {
		return false
	}
}
