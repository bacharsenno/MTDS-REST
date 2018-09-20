// Package admin contains methods related to admin functionalities
package admin

import (
	m "MTDS-REST/model"
	"net/http"
	"strconv"
	s "strings"
	"time"

	"github.com/gin-gonic/gin"
)

var initDb = m.InitDb

// PostAdminInfo creates/updates an administrator account.
//
// Input: User Object (Type = 0)
//
// Output: Post Response
//
// Example URL: http://localhost:8080/api/v1/admin
func PostAdminInfo(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var user m.User
	var post m.PostResponse
	c.Bind(&user)
	if user.Type == 0 {
		db.Save(&user)
		post.Code = 200
		post.Message = "Admin Created/Updated Successfully"
		c.JSON(http.StatusOK, post)
	} else {
		post.Code = 400
		post.Message = "Type should be 0"
		c.JSON(http.StatusBadRequest, post)
	}
}

// PostAdminNotification creates/updates a notification.
//
// Input: Notification Object
//
// Output: Post Response
//
// Example URL: http://localhost:8080/api/v1/admin/notification
func PostAdminNotification(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var notification m.Notification
	var post m.PostResponse
	c.Bind(&notification)
	if notification.StartDate.Before(time.Now()) || notification.EndDate.Before(time.Now()) {
		post.Code = 400
		post.Message = "StartDate/EndDate shouldn't be in the past"
		c.JSON(http.StatusBadRequest, post)
	} else {
		db.Save(&notification)
		post.Code = 200
		post.Message = "Notification Added Successfully."
		c.JSON(http.StatusOK, post)
	}
}

// PostAdminStudent creates/updates a notification.
//
// Input: Notification Object
//
// Output: Post Response
//
// Example URL: http://localhost:8080/api/v1/admin/student
func PostAdminStudent(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var student m.Student
	var post m.PostResponse
	c.Bind(&student)
	if student.FirstName == "" || student.LastName == "" || student.ClassID == "" {
		post.Code = 400
		post.Message = "Missing Parameters"
		c.JSON(http.StatusBadRequest, post)
	} else {
		if student.Username == "" {
			var lastStudent m.Student
			db.Limit(1).Order("LENGTH(username) desc, username desc").Find(&lastStudent)
			id := lastStudent.Username
			id = s.Trim(id, "S")
			num, _ := strconv.Atoi(id)
			num++
			student.Username = "S" + strconv.Itoa(num)
		}
		db.Save(&student)
		post.Code = 200
		post.Message = "Student added/updated successfully."
		c.JSON(http.StatusOK, post)
	}
}

// PostAdminParent creates/updates a notification.
//
// Input: Parent Object
//
// Output: Post Response
//
// Example URL: http://localhost:8080/api/v1/admin/parent
func PostAdminParent(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var parent m.Parent
	var post m.PostResponse
	c.Bind(&parent)
	if parent.FirstName != "" && parent.LastName != "" && parent.Email != "" {
		if parent.Username == "" {
			var lastParent m.Parent
			db.Limit(1).Order("LENGTH(username) desc, username desc").Find(&lastParent)
			id := lastParent.Username
			id = s.Trim(id, "P")
			num, _ := strconv.Atoi(id)
			num++
			parent.Username = "P" + strconv.Itoa(num)
			user := m.User{
				Username: parent.Username,
				Password: "PP" + strconv.Itoa(num),
				Type:     2,
			}
			db.Save(&user)
		}
		db.Save(&parent)
		post.Code = 200
		post.Message = "Parent created/updated successfully."
		c.JSON(http.StatusOK, post)
	} else {
		post.Code = 400
		post.Message = "Missing Parameters"
		c.JSON(http.StatusBadRequest, post)
	}
}

// PostAdminTeacher creates/updates a notification.
//
// Input: Notification Object
//
// Output: Post Response
//
// Example URL: http://localhost:8080/api/v1/admin/teacher
func PostAdminTeacher(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var teacher m.Teacher
	var post m.PostResponse
	c.Bind(&teacher)
	if teacher.FirstName != "" && teacher.LastName != "" && teacher.ProfilePic != "" {
		if teacher.Username == "" {
			var lastTeacher m.Teacher
			db.Limit(1).Order("LENGTH(username) desc, username desc").Find(&lastTeacher)
			id := lastTeacher.Username
			id = s.Trim(id, "T")
			num, _ := strconv.Atoi(id)
			num++
			teacher.Username = "T" + strconv.Itoa(num)
			user := m.User{
				Username: teacher.Username,
				Password: "TP" + strconv.Itoa(num),
				Type:     1,
			}
			db.Save(&user)
		}
		db.Save(&teacher)
		post.Code = 200
		post.Message = "Teacher created/updated successfully."
		c.JSON(http.StatusOK, post)
	} else {
		post.Code = 400
		post.Message = "Missing Parameters"
		c.JSON(http.StatusBadRequest, post)
	}
}

// PostAdminPayment creates/updates a payment.
//
// Input: Payment Object
//
// Output: Post Response
//
// Example URL: http://localhost:8080/api/v1/admin/teacher
func PostAdminPayment(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var payment m.Payment
	var post m.PostResponse
	c.Bind(&payment)
	if payment.StudentID == "" {
		post.Code = 400
		post.Message = "Missing Parameters"
		c.JSON(http.StatusBadRequest, post)
	} else {
		if payment.PaymentID == "" {
			var lastPayment m.Payment
			db.Limit(1).Order("LENGTH(payment_id) desc, payment_id desc").Find(&lastPayment)
			pid := lastPayment.PaymentID
			num, _ := strconv.Atoi(s.Trim(pid, "PID"))
			num++
			payment.PaymentID = "PID" + strconv.Itoa(num)
		}
		db.Save(&payment)
		post.Code = 200
		post.Message = "Payment created/updated successfully."
		c.JSON(http.StatusOK, post)
	}
}
