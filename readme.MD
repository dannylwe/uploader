[![Build Status](https://travis-ci.com/dannylwe/uploader.svg?branch=main)](https://travis-ci.com/dannylwe/uploader) [![Maintainability](https://api.codeclimate.com/v1/badges/7d8de23d3cddf3101095/maintainability)](https://codeclimate.com/github/dannylwe/uploader/maintainability) [![codecov](https://codecov.io/gh/dannylwe/uploader/branch/main/graph/badge.svg)](https://codecov.io/gh/dannylwe/uploader)

### How to run the application  
#### 1. SYSTEM INSTALL
- Install mysql running on PORT `3306`. Create user `root` with no password. Create database `test`.  
- Install golang. See [HERE](https://golang.org/doc/install)
- clone repo.
- run `cd uploader` from root of repo
- run `go run main.go`  
- The application runs on PORT `8080`
- The endpoints of interest have been documented in the Postman collection in this repo.  

For frontend:  
- Install `npm` from [HERE](https://nodejs.org/en/download/)
- run `cd uploader/frontend` in a separate terminal from root of repo
- run `npm install`
- create `.env` and set `PORT=9005`
- run `npm start`
- upload valid .csv on frontend

### 2. Run with Docker  
- Install mysql running on PORT `3306`. Create user `root` with no password. Create database `test`.   
- clone repo.
- run `cd uploader` 
- run `docker build -t uploader .` for backend build
- run `cd uploader/frontend` in a separate terminal from root of repo
- run `docker build -t frontend .` for frontend build 
- run frontend `docker run -it -p 9005:80 --rm frontend`
- run backend `docker run -it -p 8080:8080 --rm uploader`
- upload valid .csv to frontend.

### 3. Run with docker compose  [NOT working]
- start application `docker-compose -f docker-compose.yml up`  
- the application does not yet successfully connect to the docker database.
- TODO: Resolve with dynamic variables through `.env` file on backend.

### Run tests with coverage

```
go test -cover ./handlers
```

### Test Coverage
![TestResults](https://github.com/dannylwe/uploader/blob/main/images/tests.PNG)

### Postman Collection  
[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/59b99ec28fe39ad15abe) 

### Part 2: Automation from ERP 
![Automation](https://github.com/dannylwe/uploader/blob/develop/images/update-process.PNG)

### Design Choices
- Backend: Golang  
Golang was chosen as the backend langauge because of it is natively fast, easy to write and it comes out of the box with Concurrency built-in. Every request to the backend API goes off on it's own thread.  

Golang also produces very small Docker images. In the image below, the entire backend is only 10MB. This translates to quick and easy deployments. From a business perspective, less resources are used. If this service is replicated `100X`, it will only consume 100MB.  

![image-deployment](https://github.com/dannylwe/uploader/blob/develop/images/Image%20Info.PNG)

- Frontend: React  
React.js was chosen for the frontend because it has many packages that can be added to easily extend functionality. The modular nature of react makes it easy to plug and play needed functionality. The frontend codebase will however become complex.  

React.js also has a big community which makes it easier to resolve issues. 

- Docker  
In a microservices based eco-system, Docker is the backbone. For this project, docker is the intended module of deployment. The frontend and backend can be built separetely using the Dockerfiles in their respective repositories.

### Known Bugs  
- The desired date on the date picker has to be clicked twice to get correct values.  
- The results from the date picker are appended on top of the previous results(when application loads .csv). This can be handled with improved conditional rendering or a state management like `react -redux`.
