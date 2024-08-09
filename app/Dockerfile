FROM golang:alpine

WORKDIR /opt/app

COPY . .

RUN go mod download
RUN go install .
ENTRYPOINT /go/bin/app

EXPOSE 8080
