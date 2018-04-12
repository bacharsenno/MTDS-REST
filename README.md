# Middleware Technologies Project

##Initial Configuration
Sent a GET request to http://localhost:8080/api/v1/test to generate the test data; make sure mysql is running and that you specified correctly the database name in model.go InitDb function.

### Initial project choices
|DB | Language | Web Framework | Frontend
|---|---|---|---|
|mySQL|Java|Jersey|PHP|
|mongoDB|Python|Django|Javascript (Node,Angular,React,...)|
||ExpressJS|
||PHP||
||Ruby

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
## To-do 26/02/2018
- Learn about:
    - Go + web framework
    - ReactJS + JS
- Learn design best practices of FE,BE,API,DB
- Prepare design of the database (ER diagram)

## To-do 22/03/2018
- Set up git repository
- Choose GO package manager
- Work at simple REST demo
- Frontend design
- REST API design

## Useful links:
- documentation: [readthedocs](https://readthedocs.org/)
- SoapUI: API tester

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
#### Go WEB frameworks
- [Revel](http://revel.github.io/): 
    - fully featured 
    - no MongoDB
- [Gin Gonic](https://gin-gonic.github.io/gin/): 
    - minimalistic
    - designed after Martini
    - high perfomance (40x faster than Martini)
    - less support for extensions
    - offload to client
    - no enterprise features
- Martini
    - great 3rd party extensions support
    - good docs
    - 40x slower than Gin
    - not maintained since 2014
    - dependency injection (apparently is a bad thing for Golang)
- [Other possibilities, with comparisons](https://blog.usejournal.com/top-6-web-frameworks-for-go-as-of-2017-23270e059c4b)

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





