## go application
FROM golang:1.15
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
## Add this go mod download command to pull in any dependencies
# RUN go build
## we run go build to compile the binary
## executable of our Go program
RUN go build cmd/server.go
# Our start command which kicks off
## our newly created binary executable
CMD ["/app/server"]