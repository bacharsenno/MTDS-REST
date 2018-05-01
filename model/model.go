// Package model provides the definitions of all of the structs included in the project, the names of the fields piped into JSON, and initializes the MySQL database accordingly.
package model

import (
	s "strings"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	Username string `gorm:"PRIMARY_KEY" form:"Username" json:"Username"`
	Password string `gorm:"not null" form:"Password" json:"Password"`
	Type     int    `gorm:"not null" form:"Type" json:"Type"`
}

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
}

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

type Appointment struct {
	AppointmentID int       `gorm:"PRIMARY_KEY" form:"AppointmentID" json:"AppointmentID"`
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

type Schedule struct {
	ScheduleID int       `form:"ScheduleID" json:"ScheduleID"`
	Day        string    `form:"Day" json:"Day"`
	StartTime  time.Time `form:"StartTime" json:"StartTime"`
	EndTime    time.Time `form:"EndTime" json:"EndTime"`
	Semester   int       `form:"Semester" json:"Semester"`
}

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

type ParentOf struct {
	StudentID    string `form:"StudentID" json:"StudentID"`
	ParentID     string `form:"ParentID" json:"ParentID"`
	Relationship string `form:"Relationship" json:"Relationship"`
	Status       string `form:"Status" json:"Status"`
}

type ClassSchedule struct {
	TeachClass `form:"TeachClass" json:"TeachClass"`
	Time       []Schedule `form:"Schedule" json:"Schedule"`
}

type BasicStudent struct {
	FirstName  string `form:"FirstName" json:"FirstName"`
	LastName   string `form:"LastName" json:"LastName"`
	ProfilePic string `form:"ProfilePic" json:"ProfilePic"`
}

type GradeWithName struct {
	Grade     `form:"Grade" json:"Grade"`
	FirstName string `form:"FirstName" json:"FirstName"`
	LastName  string `form:"LastName" json:"LastName"`
}

type StudentWithGrade struct {
	BasicStudent   `form:"BasicStudent" json:"BasicStudent"`
	Grades         []Grade        `form:"Grade" json:"Grade"`
	GradeSummaries []GradeSummary `form:"GradeSummary" json:"GradeSummary"`
}

type StudentGradesBySubject struct {
	Subject        string         `form:"Subject" json:"Subject"`
	Grades         []Grade        `form:"Grades" json:"Grades"`
	GradeSummaries []GradeSummary `form:"GradeSummaries" json:"GradeSummaries"`
}

type StudentParentGrades struct {
	BasicStudent  `form:"BasicStudent" json:"BasicStudent"`
	SubjectGrades []StudentGradesBySubject `form:"SubjectGrades" json:"SubjectGrades"`
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
	//Tables: User Teacher Parent Student Subject Class Grade Payment Notification NotificationTopics Appointment TeachClass ParentOf
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
func GetDateString(scope string, offset int) string {
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
