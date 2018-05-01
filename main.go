package main

import (
	"MTDS-REST/utils"

	_ "github.com/go-sql-driver/mysql"
)

//This is the main function. All it does is create a router, calls a method that sets it up, and run it on port 8080.
func main() {
	r := utils.R
	utils.SetupRoutes()
	r.Run(":8080")
}

//$params = '{"firstname": "Bachar", "lastname": "Senno", "email": "abc@def.com", "username": "username", "password": "password", "phonenumber": "123456789"}'
//Invoke-RestMethod -URI http://localhost:8080/api/v1/teacher/1 -Method "PUT" -Body $params -ContentType 'application/json'
