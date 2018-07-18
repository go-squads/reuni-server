# create image from the official Go image
FROM golang:1.10.3-alpine3.8

RUN apk add --update tzdata \
    bash wget curl git;

# Create binary directory, install glide and fresh
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh 

# define work directory
ADD . /go/src/github.com/go-squads/reuni-server
WORKDIR /go/src/github.com/go-squads/reuni-server

# serve the app
CMD dep ensure && go test -v ./... && go run main.go
EXPOSE 8080
