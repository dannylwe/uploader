## Build Container
`docker build -t frontend .`

## Run container
` docker run -it -p 9008:80 --rm frontend`

### for Windows
`winpty docker run -it -p 9008:80 --rm frontend`

## Application
React application is running on `localhost:9008/`