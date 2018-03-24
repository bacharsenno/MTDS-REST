package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Teacher struct {
	ID          int    `gorm:"PRIMARY_KEY, AUTOINCREMENT" form:"ID" json:"ID"`
	Firstname   string `gorm:"not null" form:"firstname" json:"firstname"`
	Lastname    string `gorm:"not null" form:"lastname" json:"lastname"`
	Email       string `gorm:"not null" form:"email" json:"email"`
	Username    string `gorm:"not null" form:"username" json:"username"`
	Password    string `gorm:"not null" form:"password" json:"password"`
	PhoneNumber string `gorm:"not null" form:"phonenumber" json:"phonenumber"`
}

type Parent struct {
	ID          int    `gorm:"PRIMARY_KEY, AUTO_INCREMENT" form:"ID" json:"ID"`
	Firstname   string `gorm:"not null" form:"firstname" json:"firstname"`
	Lastname    string `gorm:"not null" form:"lastname" json:"lastname"`
	Email       string `gorm:"not null" form:"email" json:"email"`
	Username    string `gorm:"not null" form:"username" json:"username"`
	Password    string `gorm:"not null" form:"password" json:"password"`
	PhoneNumber string `gorm:"not null" form:"phonenumber" json:"phonenumber"`
}

type Subject struct {
	ID    int    `gorm:"PRIMARY_KEY, AUTO_INCREMENT" form:"ID" json:"ID"`
	Name  string `gorm:"not null" form:"name" json:"name"`
	Class int    `gorm:"not null" form:"class" json:"class"`
}

type Class struct {
	ID       int    `gorm:"PRIMARY_KEY, AUTO_INCREMENT" form:"ID" json:"ID"`
	Name     string `gorm:"not null" form:"name" json:"name"`
	Location string `gorm:"not null" form:"location" json:"location"`
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
	if !db.HasTable(&Teacher{}) {
		db.CreateTable(&Teacher{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Teacher{})
	}
	if !db.HasTable(&Parent{}) {
		db.CreateTable(&Parent{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Parent{})
	}
	if !db.HasTable(&Subject{}) {
		db.CreateTable(&Subject{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Subject{})
	}
	if !db.HasTable(&Class{}) {
		db.CreateTable(&Class{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Class{})
	}
	return db
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(Cors())

	teacher := r.Group("api/v1/teacher")
	{
		teacher.POST("/", PostTeacher)
		teacher.GET("/", GetTeachers)
		teacher.GET("/:id", GetTeacher)
		teacher.PUT("/:id", UpdateTeacher)
		teacher.DELETE("/:id", DeleteTeacher)
	}

	parent := r.Group("api/v1/parent")
	{
		parent.POST("/", PostParent)
		parent.GET("/", GetParents)
		parent.GET("/:id", GetParent)
		parent.PUT("/:id", UpdateParent)
		parent.DELETE("/:id", DeleteParent)
	}

	subject := r.Group("api/v1/subject")
	{
		subject.POST("/", PostSubject)
		subject.GET("/", GetSubject)
		subject.GET("/:id", GetSubject)
		subject.PUT("/:id", UpdateSubject)
		subject.DELETE("/:id", DeleteSubject)
	}

	class := r.Group("api/v1/class")
	{
		class.POST("/", PostClass)
		class.GET("/", GetClass)
		class.GET("/:id", GetClass)
		class.PUT("/:id", UpdateClass)
		class.DELETE("/:id", DeleteClass)
	}

	r.Run(":8080")
}

func PostTeacher(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var teacher Teacher
	c.Bind(&teacher)
	db.Create(&teacher)
	c.JSON(201, gin.H{"success": teacher})
	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/users
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
	id := c.Params.ByName("id")
	var teacher Teacher
	db.First(&teacher, id)
	if teacher.ID != 0 {
		c.JSON(200, teacher)
	} else {
		c.JSON(404, gin.H{"error": "Teacher not found"})
	}
}

func UpdateTeacher(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var teacher Teacher
	db.First(&teacher, id)
	var newTeacher Teacher
	c.Bind(&newTeacher)
	result := Teacher{
		ID:          teacher.ID,
		Firstname:   newTeacher.Firstname,
		Lastname:    newTeacher.Lastname,
		Email:       newTeacher.Email,
		Username:    newTeacher.Username,
		Password:    newTeacher.Password,
		PhoneNumber: newTeacher.PhoneNumber,
	}
	db.Save(&result)
	c.JSON(200, gin.H{"success": result})
	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/users/1
}

func DeleteTeacher(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var teacher Teacher
	db.First(&teacher, id)
	if teacher.ID != 0 {
		db.Delete(&teacher)
		c.JSON(200, gin.H{"success": "Teacher #" + id + " deleted"})
	} else {
		c.JSON(404, gin.H{"error": "Teacher not found"})
	}
	// curl -i -X DELETE http://localhost:8080/api/v1/users/1
}

func PostParent(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var parent Parent
	c.Bind(&parent)
	db.Create(&parent)
	c.JSON(201, gin.H{"success": parent})
	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/users
}

func GetParents(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var parents []Parent
	db.Find(&parents)
	c.JSON(200, parents)
}

func GetParent(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var parent Parent
	db.First(&parent, id)
	if parent.ID != 0 {
		c.JSON(200, parent)
	} else {
		c.JSON(404, gin.H{"error": "Parent not found"})
	}
}

func UpdateParent(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var parent Parent
	db.First(&parent, id)
	var newParent Parent
	c.Bind(&newParent)
	result := Parent{
		ID:          parent.ID,
		Firstname:   newParent.Firstname,
		Lastname:    newParent.Lastname,
		Email:       newParent.Email,
		Username:    newParent.Username,
		Password:    newParent.Password,
		PhoneNumber: newParent.PhoneNumber,
	}
	db.Save(&result)
	c.JSON(200, gin.H{"success": result})
	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/users/1
}

func DeleteParent(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var parent Parent
	db.First(&parent, id)
	if parent.ID != 0 {
		db.Delete(&parent)
		c.JSON(200, gin.H{"success": "Parent #" + id + " deleted"})
	} else {
		c.JSON(404, gin.H{"error": "Parent not found"})
	}
	// curl -i -X DELETE http://localhost:8080/api/v1/users/1
}

func PostSubject(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var subject Subject
	c.Bind(&subject)
	db.Create(&subject)
	c.JSON(201, gin.H{"success": subject})
	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/users
}

func GetSubjects(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var subjects []Subject
	db.Find(&subjects)
	c.JSON(200, subjects)
}

func GetSubject(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var subject Subject
	db.First(&subject, id)
	if subject.ID != 0 {
		c.JSON(200, subject)
	} else {
		c.JSON(404, gin.H{"error": "Subject not found"})
	}
}

func UpdateSubject(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var subject Subject
	db.First(&subject, id)
	var newSubject Subject
	c.Bind(&newSubject)
	result := Subject{
		ID:    subject.ID,
		Name:  newSubject.Name,
		Class: newSubject.Class,
	}
	db.Save(&result)
	c.JSON(200, gin.H{"success": result})
	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/users/1
}

func DeleteSubject(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var subject Subject
	db.First(&subject, id)
	if subject.ID != 0 {
		db.Delete(&subject)
		c.JSON(200, gin.H{"success": "Subject #" + id + " deleted"})
	} else {
		c.JSON(404, gin.H{"error": "Subject not found"})
	}
	// curl -i -X DELETE http://localhost:8080/api/v1/users/1
}

func PostClass(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var class Class
	c.Bind(&class)
	db.Create(&class)
	c.JSON(201, gin.H{"success": class})
	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/users
}

func GetClasses(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	var classes []Class
	db.Find(&classes)
	c.JSON(200, classes)
}

func GetClass(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var class Class
	db.First(&class, id)
	if class.ID != 0 {
		c.JSON(200, class)
	} else {
		c.JSON(404, gin.H{"error": "Class not found"})
	}
}

func UpdateClass(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var class Class
	db.First(&class, id)
	var newClass Class
	c.Bind(&newClass)
	result := Class{
		ID:       class.ID,
		Name:     newClass.Name,
		Location: newClass.Location,
	}
	db.Save(&result)
	c.JSON(200, gin.H{"success": result})
	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/users/1
}

func DeleteClass(c *gin.Context) {
	db := InitDb()
	defer db.Close()
	id := c.Params.ByName("id")
	var class Class
	db.First(&class, id)
	if class.ID != 0 {
		db.Delete(&class)
		c.JSON(200, gin.H{"success": "Class #" + id + " deleted"})
	} else {
		c.JSON(404, gin.H{"error": "Class not found"})
	}
	// curl -i -X DELETE http://localhost:8080/api/v1/users/1
}

func OptionsUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
