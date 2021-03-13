# tuto : https://tutorialedge.net/golang/go-docker-tutorial/#why-docker
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
RUN go build serverGOLANG.go sendmailGOLANG.go databaseGOLANG.go cryptGOLANG.go

# EXPOSE 80 is necessary for production, but does nothing in local
EXPOSE 80

## Our start command which kicks off
## our newly created binary executable
CMD ["/app/serverGOLANG"]

# from the directory where Dockerfile is, run the command :
# ==> 
# ==> ./before_compile_backend.sh; docker build . -t mtt-backend; ./after_compile_backend.sh
# see the results :
# ==> docker images
# run it 
# ==> docker run -p 8090:8090 -it mtt-backend
# delete a docker image (add --force to force)
# ==> docker rmi mtt-backend
# purge
# ==> docker images purge
#
# To build a distant image (will go to Docker Hub, 32681733 is my Docker Hub ID)
# ==> ./before_compile_backend.sh; docker build . -t 32681733/mtt-backend; ./after_compile_backend.sh
# To test locally the distant Docker image
# ==> docker run -it -p 8090:8090 -m 1536m --memory-swap=1536m 32681733/mtt-backend
# To push the image to Docker Hub (32681733 is my Docker Hub ID)
# ==> docker push 32681733/mtt-backend
