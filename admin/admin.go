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

// GetAdminList returns a list of all admins.
//
// Input: None.
//
// Output: []Users
//
// Example URL: http://localhost:8080/api/v1/admin/list
func GetAdminList(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var admins []m.User
	if m.IsAuthorizedUserType(c, db, 0) {
		db.Where("type = 0").Find(&admins)
		c.JSON(http.StatusOK, admins)
	} else {
		c.JSON(http.StatusUnauthorized, m.UnauthorizedResponse)
	}
}

// PostAdminInfo creates/updates an administrator account.
//
// Input: User Object (Type = 0)
//
// Output: Post Response
//
// Example URL: http://localhost:8080/api/v1/admin/A1/info
func PostAdminInfo(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var user m.User
	var post m.PostResponse
	c.Bind(&user)

	if m.IsAuthorizedUserType(c, db, 0) {
		if user.Type == 0 {
			if user.Username == "" {
				var lastAdmin m.User
				db.Limit(1).Order("LENGTH(username) desc, username desc").Find(&lastAdmin)
				id := lastAdmin.Username
				id = s.Trim(id, "A")
				num, _ := strconv.Atoi(id)
				num++
				user.Username = "A" + strconv.Itoa(num)
			}
			db.Save(&user)
			c.JSON(http.StatusOK, user)
		} else {
			post.Code = 400
			post.Message = "Type should be 0"
			c.JSON(http.StatusBadRequest, post)
		}
	} else {
		c.JSON(http.StatusUnauthorized, m.UnauthorizedResponse)
	}
}

// PostAdminNotification creates/updates a notification.
//
// Input: Notification Object
//
// Output: Post Response
//
// Example URL: http://localhost:8080/api/v1/admin/A1/notifications
func PostAdminNotification(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var notification m.Notification
	var post m.PostResponse
	c.Bind(&notification)
	aid := c.Params.ByName("aid")

	if m.IsAuthorized(c, db, aid) {
		if notification.StartDate.Before(time.Now()) && notification.EndDate.Before(time.Now()) && notification.StartDate.Before(notification.EndDate) {
			post.Code = 400
			post.Message = "StartDate/EndDate shouldn't be in the past"
			c.JSON(http.StatusBadRequest, post)
		} else {
			db.Save(&notification)
			c.JSON(http.StatusOK, notification)
		}
	} else {
		c.JSON(http.StatusUnauthorized, m.UnauthorizedResponse)
	}
}

// PostAdminStudent creates/updates a student.
//
// Input: Student Object
//
// Output: Post Response
//
// Example URL: http://localhost:8080/api/v1/admin/A1/students/ || http://localhost:8080/api/v1/admin/A1/students/S1
func PostAdminStudent(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var student m.Student
	var post m.PostResponse
	c.Bind(&student)
	aid := c.Params.ByName("aid")
	urlParam := c.Params.ByName("sid")
	if urlParam != "" && urlParam != student.Username {
		post.Code = 405
		post.Message = "Parameter Mismatch"
		c.JSON(http.StatusBadRequest, post)
	} else {
		if m.IsAuthorized(c, db, aid) {
			if student.FirstName == "" || student.LastName == "" || student.ClassID == "" {
				post.Code = 400
				post.Message = "Missing Parameters"
				c.JSON(http.StatusBadRequest, post)
			} else {
				if urlParam != "" && urlParam != student.Username {
					post.Code = 405
					post.Message = "Parameter Mismatch"
					c.JSON(http.StatusBadRequest, post)
				} else if urlParam == "" && student.Username == "" {
					var lastStudent m.Student
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
		} else {
			c.JSON(http.StatusUnauthorized, m.UnauthorizedResponse)
		}
	}
}

// PostAdminParent creates/updates a parent.
//
// Input: Parent Object
//
// Output: Post Response
//
// Example URL: http://localhost:8080/api/v1/admin/A1/parents || http://localhost:8080/api/v1/admin/A1/parents/P1
func PostAdminParent(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var parent m.Parent
	var post m.PostResponse
	c.Bind(&parent)
	aid := c.Params.ByName("aid")
	urlParam := c.Params.ByName("pid")
	if urlParam != "" && urlParam != parent.Username {
		post.Code = 405
		post.Message = "Parameter Mismatch"
		c.JSON(http.StatusBadRequest, post)
	} else {
		if m.IsAuthorized(c, db, aid) {
			if parent.FirstName != "" && parent.LastName != "" && parent.Email != "" {
				if urlParam != "" && urlParam != parent.Username {
					post.Code = 405
					post.Message = "Parameter Mismatch"
					c.JSON(http.StatusBadRequest, post)
				} else if urlParam == "" && parent.Username == "" {
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
				c.JSON(http.StatusOK, parent)
			} else {
				post.Code = 400
				post.Message = "Missing Parameters"
				c.JSON(http.StatusBadRequest, post)
			}
		} else {
			c.JSON(http.StatusUnauthorized, m.UnauthorizedResponse)
		}
	}
}

// PostAdminTeacher creates/updates a teacher.
//
// Input: Teacher Object
//
// Output: Post Response
//
// Example URL: http://localhost:8080/api/v1/admin/A1/teachers || http://localhost:8080/api/v1/admin/A1/teachers/T1
func PostAdminTeacher(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var teacher m.Teacher
	var post m.PostResponse
	c.Bind(&teacher)
	aid := c.Params.ByName("aid")
	urlParam := c.Params.ByName("tid")
	if urlParam != "" && urlParam != teacher.Username {
		post.Code = 405
		post.Message = "Parameter Mismatch"
		c.JSON(http.StatusBadRequest, post)
	} else {
		if m.IsAuthorized(c, db, aid) {
			if teacher.FirstName != "" && teacher.LastName != "" && teacher.ProfilePic != "" {
				if urlParam != "" && urlParam != teacher.Username {
					post.Code = 405
					post.Message = "Parameter Mismatch"
					c.JSON(http.StatusBadRequest, post)
				} else if urlParam == "" && teacher.Username == "" {
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
				c.JSON(http.StatusOK, teacher)
			} else {
				post.Code = 400
				post.Message = "Missing Parameters"
				c.JSON(http.StatusBadRequest, post)
			}
		} else {
			c.JSON(http.StatusUnauthorized, m.UnauthorizedResponse)
		}
	}
}

// PostAdminPayment creates/updates a payment.
//
// Input: Payment Object
//
// Output: Post Response
//
// Example URL: http://localhost:8080/api/v1/admin/A1/payments || http://localhost:8080/api/v1/admin/A1/payments/PID5
func PostAdminPayment(c *gin.Context) {
	db := initDb()
	defer db.Close()
	var payment m.Payment
	var post m.PostResponse
	c.Bind(&payment)
	aid := c.Params.ByName("aid")
	urlParam := c.Params.ByName("pid")
	if urlParam != "" && urlParam != payment.PaymentID {
		post.Code = 405
		post.Message = "Parameter Mismatch"
		c.JSON(http.StatusBadRequest, post)
	} else {
		if m.IsAuthorized(c, db, aid) {
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
				c.JSON(http.StatusOK, payment)
			}
		} else {
			c.JSON(http.StatusUnauthorized, m.UnauthorizedResponse)
		}
	}
}
