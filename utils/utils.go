package utils

import (
	"awesomeProject/model"
	"fmt"
	"strconv"
	s "strings"

	"github.com/gin-gonic/gin"
)

var R = gin.Default()

type User = model.User
type Teacher = model.Teacher
type Parent = model.Parent
type Student = model.Student
type Class = model.Class
type Subject = model.Subject
type TeachClass = model.TeachClass
type ParentOf = model.ParentOf

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
		teacher.POST("/", PostTeacher)
		teacher.GET("/", GetTeachers)
		teacher.GET("/:id", GetTeacher)
		teacher.PUT("/:id", UpdateTeacher)
		teacher.DELETE("/:id", DeleteTeacher)
	}

	parent := R.Group("api/v1/parent")
	{
		parent.POST("/", PostParent)
		parent.GET("/", GetParents)
		parent.GET("/:id", GetParent)
		parent.PUT("/:id", UpdateParent)
		parent.DELETE("/:id", DeleteParent)
	}

	subject := R.Group("api/v1/subject")
	{
		subject.POST("/", PostSubject)
		subject.GET("/", GetSubject)
		subject.GET("/:id", GetSubject)
		subject.PUT("/:id", UpdateSubject)
		subject.DELETE("/:id", DeleteSubject)
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
	for i := 1; i <= 10; i++ {
		class := Class{
			ClassID: "C" + strconv.Itoa(i),
		}
		db.Create(&class)
	}
	k := 1
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
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
		subject := Subject{
			ID:   i,
			Name: "SName" + strconv.Itoa(i),
		}
		db.Create(&subject)
	}
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			teachClass := TeachClass{
				TeacherID: "T" + strconv.Itoa(i),
				ClassID:   "C" + strconv.Itoa(j),
				SubjectID: i,
			}
			db.Create(teachClass)
		}
	}
	for i := 1; i <= 100; i++ {
		parentOf := ParentOf{
			StudentID: "S" + strconv.Itoa(i),
			ParentID:  "P" + strconv.Itoa(i),
		}
		db.Create(&parentOf)
	}

}

func PostLogin(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var user User
	c.Bind(&user)
	username := user.Username
	password := user.Password
	fmt.Println("Username: " + username)
	fmt.Println("Password: " + password)
	if s.HasPrefix(username, "T") {
		var teacher Teacher
		db.Where("username = ? AND password = ?", username, password).First(&teacher)
		if teacher.Username != "" {
			c.JSON(200, teacher)
		} else {
			c.JSON(404, gin.H{"error": "Teacher not found"})
		}
	} else if s.HasPrefix(username, "P") {
		var parent Parent
		db.Where("username = ? AND password = ?", username, password).First(&parent)
		if parent.Username != "" {
			c.JSON(200, parent)
		} else {
			c.JSON(404, gin.H{"error": "Parent not found"})
		}
	} else {
		c.JSON(408, gin.H{"error": "Invalid Username"})
	}
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
	db := InitDb()
	defer db.Close()
	username := c.Params.ByName("Username")
	var teacher Teacher
	db.Where("username = ?", username).First(&teacher)
	if teacher.Username != "" {
		c.JSON(200, teacher)
	} else {
		c.JSON(404, gin.H{"error": "Teacher not found"})
	}
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
