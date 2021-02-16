## We specify the base image we need for our
## go application
## FROM golang:1.13.15 is the golang docker version
## FROM golang:1.15.8 is the golang development version
FROM golang:1.15.8
## We create an /app directory within our
## image that will hold our application source
## files
RUN mkdir /app
## We copy everything in the root directory
## into our /app directory
ADD . /app
## We specify that we now wish to execute 
## any further commands inside our /app
## directory
WORKDIR /app
## get the "dependencies"
RUN go get github.com/jordan-wright/email
RUN go get github.com/twinj/uuid
RUN go get modernc.org/sqlite

## we run go build to compile the binary
## executable of our Go program
RUN go build serverGOLANG.go sendmailGOLANG.go databaseGOLANG.go 
## Our start command which kicks off
## our newly created binary executable
CMD ["/app/serverGOLANG"]
