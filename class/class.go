// Package class contains methods related to viewing Class data
package class

import (
	m "MTDS-REST/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var initDb = m.InitDb

// GetClasses returns the list of all the classes.
//
// Input: [Class ID]
//
// Output: []Classes
//
// Example URL: http://localhost:8080/api/v1/class/C1
func GetClasses(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var classes []m.Class
	class := c.Params.ByName("cid")

	if m.IsAuthorizedUserType(c, db, 1) {
		if class != "" {
			db.Table("teach_classes").Select("class_id, location, year").Group("class_id").Where("class_id = ?", class).Order("LENGTH(class_id), class_id").Scan(&classes)
		} else {
			db.Table("teach_classes").Select("class_id, location, year").Group("class_id").Order("LENGTH(class_id), class_id").Scan(&classes)
		}
		if len(classes) > 0 {
			c.JSON(http.StatusOK, classes)
		} else {
			c.JSON(http.StatusOK, make([]string, 0))
		}
	} else {
		c.JSON(http.StatusUnauthorized, m.UNAUTHORIZED_RESPONSE)
	}
}

// GetClassStudents returns the students in a specific class.
//
// Input: Class ID.
//
// Output: []Student.
//
// Example URL: http://localhost:8080/api/v1/class/C1/students
func GetClassStudents(c *gin.Context) {
	db := initDb()
	defer db.Close()
	class := c.Params.ByName("cid")
	var students []m.Student

	if m.IsAuthorizedUserType(c, db, 1) {
		db.Where("class_id = ?", class).Find(&students)
		if len(students) > 0 {
			var objectsWithLink []m.StudentWithLink
			var tempObj m.StudentWithLink
			for i := 0; i < len(students); i++ {
				tempObj.Student = students[i]
				tempObj.Link = "http://localhost:8080/api/v1/student/" + students[i].Username + "/info"
				objectsWithLink = append(objectsWithLink, tempObj)
			}
			c.JSON(http.StatusOK, objectsWithLink)
		} else {
			c.JSON(http.StatusBadRequest, make([]string, 0))
		}
	} else {
		c.JSON(http.StatusUnauthorized, m.UNAUTHORIZED_RESPONSE)
	}
}
