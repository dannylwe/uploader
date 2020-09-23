[![Build Status](https://travis-ci.com/dannylwe/uploader.svg?branch=main)](https://travis-ci.com/dannylwe/uploader) [![Maintainability](https://api.codeclimate.com/v1/badges/7d8de23d3cddf3101095/maintainability)](https://codeclimate.com/github/dannylwe/uploader/maintainability)

### How to run the application  
- Install mysql running on PORT `3306`. Create user `Root` with no password.  
- Install golang. See [HERE](https://golang.org/doc/install)
- clone repo.
- run `cd uploader`  
- run `go run main.go`  
- The application runs on PORT `8080`
- The endpoints of interest have been documented in the Postman collection in this repo.

#### Run with docker  [NOT working]
- start application `docker-compose -f docker-compose.yml up`

### Run tests with coverage
Note: To run tests locally, uncomment #18 in handlers/render.go and comment out #21 in same file. Revert to run the application.
```
go test -cover ./handlers
```

### Test Coverage
![TestResults](https://github.com/dannylwe/uploader/blob/main/images/tests.PNG)

### Postman Collection  
[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/59b99ec28fe39ad15abe) 