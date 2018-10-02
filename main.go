package main

import (
	"MTDS-REST/utils"
)

//This is the main function. All it does is create a router, calls a method that sets it up, and run it on port 8080.
func main() {
	r := utils.R
	utils.SetupRoutes()
	r.Run(":8080")
}
