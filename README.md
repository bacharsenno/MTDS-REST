# Middleware Technologies Project

## Installation
```
go get github.com/gin-gonic/gin
go get -u github.com/pjebs/restgate
go get -u golang.org/x/crypto/bcrypt
go get -u github.com/gin-contrib/sessions
go get -u github.com/utrack/gin-csrf
go run main.go
```

## Initial Configuration

- Drop all of the tables you have after every pull because sometimes I change the table structures
- Send a GET request to http://localhost:8080/api/v1/test to generate the test data; make sure mysql is running and that you specified correctly the database name in model.go InitDb function.

- Implemented calls (so far):
- [http://localhost:8080/api/v1/test](http://localhost:8080/api/v1/test) : GENERATE DATABASE
- [http://localhost:8080/api/v1/login](http://localhost:8080/api/v1/login) :
	- POST Body Example: {'username': 'P1', 'password': 'PP1'}
	- Reply: 404 if not found, otherwise Parent/Teacher object	
	
- Teacher API:				
	- [http://localhost:8080/api/v1/teacher/notifications?id=T1](http://localhost:8080/api/v1/teacher/notifications?id=T1) : returns notifications of teacher T1
	- [http://localhost:8080/api/v1/teacher/appointments?id=T1&scope=day](http://localhost:8080/api/v1/teacher/appointments?id=T1&scope=day) (or scope=week): returns appointments for current day/week
	- [http://localhost:8080/api/v1/teacher/agenda?id=T1&scope=day](http://localhost:8080/api/v1/teacher/agenda?id=T1&scope=day) (or scope=week): returns classes + schedule for current day/week
	- [http://localhost:8080/api/v1/teacher/classes?id=T1](http://localhost:8080/api/v1/teacher/classes?id=T1) : returns all classes of Teacher T1
	
- Parent API:
	- [http://localhost:8080/api/v1/parent/notifications?id=P1](http://localhost:8080/api/v1/parent/notifications?id=P1) : return notifications of parent P1
	- [http://localhost:8080/api/v1/parent/appointments?id=P1&scope=day](http://localhost:8080/api/v1/parent/appointments?id=P1&scope=day) (or scope=week): returns appointments for current day/week
	- [http://localhost:8080/api/v1/parent/students?id=P1](http://localhost:8080/api/v1/parent/students?id=P1) : returns students associated with parent P1
	- [http://localhost:8080/api/v1/parent/payments?id=P1](http://localhost:8080/api/v1/parent/payments?id=P1) : returns payments associated with parent P1
	


__Documentation__: 

### Work phases
- Application design/planning
    - requirements
    - architecture
        - API
        - hypermedia
    - testing
- Development
    - backend + DB
    - Hypermedia

---

# REST Project
## Useful links:
- documentation: [readthedocs](https://readthedocs.org/)
- SoapUI: API tester
- Password Storage: [bcrypt](https://astaxie.gitbooks.io/build-web-application-with-golang/en/09.5.html)

### Front-End Stuff
- graphs: 
    - [chartJS](http://www.chartjs.org/)
    - [Highcharts](https://www.highcharts.com/) (used by Erica's colleagues)
- Javascript
    - [Here's a "quick" overview of JS main characteristics](https://developer.mozilla.org/en-US/docs/Web/JavaScript/A_re-introduction_to_JavaScript)
- React
    - [React Tutorial](https://reactjs.org/tutorial/tutorial.html)
### Go language
- [Go official tour](https://tour.golang.org/welcome/1)
More general than "go by example"
- [examples](https://gobyexample.com/)
- [Webservices with Go in 5 min](https://blog.smartbear.com/web-development/how-to-build-a-web-service-in-5-minutes-with-go/)
- [Build WebApp with GO](https://astaxie.gitbooks.io/build-web-application-with-golang/en/08.0.html)
- Recent article on frameworks: [7 Frameworks to build a Web API](https://nordicapis.com/7-frameworks-to-build-a-rest-api-in-go/)
- [Basic REST example in Go](https://dev.to/codehakase/building-a-restful-api-with-go)
- [GO Project structure](https://golang.org/doc/code.html)
- [GO Package management](https://github.com/golang/go/wiki/PackageManagementTools)
- https://medium.com/@cgrant/developing-a-simple-crud-api-with-go-gin-and-gorm-df87d98e6ed1
- https://medium.com/@thedevsaddam/build-restful-api-service-in-golang-using-gin-gonic-framework-85b1a6e176f3
### Database and ORM with GO
- discussion >>https://forum.golangbridge.org/t/which-orm-is-best-for-golang/6268/2
- http://www.hydrogen18.com/blog/golang-orms-and-why-im-still-not-using-one.html
- http://go-database-sql.org/index.html
### REST API design
- https://byrondover.github.io/post/restful-api-guidelines/
- https://dzone.com/articles/5-basic-rest-api-design-guidelines
- https://dzone.com/articles/common-mistakes-in-rest-api-design

## Design phase

## UI





