FROM golang:1.12
WORKDIR /go/metahub
COPY . ./
RUN go mod tidy
RUN go build
