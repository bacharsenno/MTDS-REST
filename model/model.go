package model

import "github.com/jinzhu/gorm"

type User struct {
	Username string `gorm:"not null, PRIMARY_KEY" form:"Username" json:"Username"`
	Password string `gorm:"not null" form:"Password" json:"Password"`
	Type     int    `gorm:"not null" form:"Type" json:"Type"`
}

type Teacher struct {
	Username         string `gorm:"not null, PRIMARY_KEY" form:"Username" json:"Username"`
	FirstName        string `gorm:"not null" form:"FirstName" json:"FirstName"`
	LastName         string `gorm:"not null" form:"LastName" json:"LastName"`
	Email            string `gorm:"not null" form:"Email" json:"Email"`
	PhoneNumber      string `gorm:"not null" form:"PhoneNumber" json:"PhoneNumber"`
	DateOfBirth      string `form:"DateOfBirth" json:"DateOfBirth"`
	PlaceOfBirth     string `form:"PlaceOfBirth" json:"PlaceOfBirth"`
	Nationality      string `form:"Nationality" json:"Nationality"`
	Address          string `form:"Address" json:"Address"`
	FiscalCode       string `form:"FiscalCode" json:"FiscalCode"`
	GradDegree       string `form:"GradDegree" json:"GradDegree"`
	GradFieldOfStudy string `form:"GradFieldOfStudy" json:"GradFieldOfStudy"`
	GradGrade        string `form:"GradGrade" json:"GradGrade"`
	GradSchool       string `form:"GradSchool" json:"GradSchool"`
	SeniorityLevel   string `form:"SeniorityLevel" json:"SeniorityLevel"`
	StartDate        string `form:"StartDate" json:"StartDate"`
	EndDate          string `form:"EndDate" json:"EndDate"`
	Status           string `form:"Status" json:"Status"`
}

type Parent struct {
	Username    string `gorm:"not null, PRIMARY_KEY" form:"Username" json:"Username"`
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
	Username     string `gorm:"not null, PRIMARY_KEY" form:"Username" json:"Username"`
	FirstName    string `gorm:"not null" form:"FirstName " json:"FirstName "`
	LastName     string `gorm:"not null" form:"LastName" json:"LastName"`
	Email        string `gorm:"not null" form:"Email" json:"Email"`
	PhoneNumber  string `gorm:"not null" form:"PhoneNumber" json:"PhoneNumber"`
	ClassID      string `gorm:"not null" form:"ClassID" json:"ClassID"`
	Nationality  string `form:"Nationality" json:"Nationality"`
	DateOfBirth  string `form:"DateOfBirth" json:"DateOfBirth"`
	PlaceOfBirth string `form:"PlaceOfBirth" json:"PlaceOfBirth"`
	Address      string `form:"Address" json:"Address"`
	FiscalCode   string `form:"FiscalCode" json:"FiscalCode"`
	EnrolledDate string `form:"EnrolledDate" json:"EnrolledDate"`
	EndDate      string `form:"EndDate" json:"EndDate"`
	Status       string `form:"Status" json:"Status"`
}

// type Class struct {
// 	ClassID  string `gorm:"PRIMARY_KEY" form:"ID" json:"ID"`
// 	Location string `gorm:"not null" form:"Location" json:"Location"`
// 	Year     string `gorm:"not null" form:"Year" json:"Year"`
// }

type Grade struct {
	TeacherID string `form:"TeacherID" json:"TeacherID"`
	StudentID string `form:"StudentID" json:"StudentID"`
	SubjectID int    `form:"SubjectID" json:"SubjectID"`
	Year      string `form:"Year" json:"Year"`
	Date      string `form:"Date" json:"Date"`
	Grade     int    `form:"Grade" json:"Grade"`
	Remarks   string `form:"Remarks" json:"Remarks"`
}

type Payment struct {
	PaymentID   string `form:"PaymentID" json:"PaymentID"`
	ParentID    string `form:"ParentID" json:"ParentID"`
	StudentID   string `form:"StudentID" json:"StudentID"`
	Amount      int    `form:"Amount" json:"Amount"`
	Deadline    string `form:"Deadline" json:"Deadline"`
	CreatedOn   string `form:"CreatedOn" json:"CreatedOn"`
	Status      string `form:"Status" json:"Status"`
	Description string `form:"Description" json:"Description"`
}

type Notification struct {
	SenderID      string `form:"SenderID" json:"SenderID"`
	DestinationID string `form:"DestinationID" json:"DestinationID"`
	//Topic i.e. Payment due, School Trip, Parent-Teacher meeting, Student Medical Checkups etc...
	Topic       string `form:"Topic" json:"Topic"`
	Title       string `form:"Title" json:"Title"`
	Description string `form:"Description" json:"Description"`
	Priority    string `form:"Priority" json:"Priority"`
	StartDate   string `form:"StartDate" json:"StartDate"`
	EndDate     string `form:"EndDate" json:"EndDate"`
	Status      string `form:"Status" json:"Status"`
}

type Appointment struct {
	AppointmentID int    `form:"AppointmentID" json:"AppointmentID"`
	TeacherID     string `form:"TeacherID" json:"TeacherID"`
	ParentID      string `form:"ParentID" json:"ParentID"`
	Date          string `form:"Date" json:"Date"`
	FullDay       bool   `form:"FullDay" json:"FullDay"`
	StartTime     string `form:"StartTime" json:"StartTime"`
	EndTime       string `form:"EndTime" json:"EndTime"`
	Remarks       string `form:"Remarks" json:"Remarks"`
	Status        string `form:"Status" json:"Status"`
}

type Schedule struct {
	ScheduleID int    `form:"ScheduleID" json:"ScheduleID"`
	Day        string `form:"Day" json:"Day"`
	StartTime  string `form:"StartTime" json:"StartTime"`
	EndTime    string `form:"EndTime" json:"EndTime"`
	StartDate  string `form:"StartDate" json:"StartDate"`
	EndDate    string `form:"EndDate" json:"EndDate"`
}

type TeachClass struct {
	TeacherID  string `form:"TeacherID" json:"TeacherID"`
	ClassID    string `form:"ClassID" json:"ClassID"`
	Subject    string `form:"Subject" json:"Subject"`
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

func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("mysql", "root:@/testdb")
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
	// if !db.HasTable(&Class{}) {
	// 	db.CreateTable(&Class{})
	// 	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Class{})
	// }
	if !db.HasTable(&Schedule{}) {
		db.CreateTable(&Schedule{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Schedule{})
	}
	if !db.HasTable(&Grade{}) {
		db.CreateTable(&Grade{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Grade{})
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
