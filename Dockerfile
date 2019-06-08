FROM node AS ui
WORKDIR /go/metahub
COPY ./ui ./ui
COPY ./static ./static
COPY ./templates ./templates
WORKDIR /go/metahub/ui/
RUN npm install
RUN npm run build

FROM golang:1.12 AS go
WORKDIR /go/metahub
COPY ./cmd ./cmd
COPY ./go.mod .
COPY ./go.sum .
COPY ./pkg ./pkg
RUN go mod tidy
WORKDIR /go/metahub/cmd/boltdb
# static build
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -a -ldflags '-extldflags "-static"' .


# Go binary serves the ui web content
FROM scratch
WORKDIR /go/metahub/
COPY --from=ui /go/metahub/static /srv/html/static
COPY --from=ui /go/metahub/templates/gen/index.html /srv/html/
COPY --from=go /go/metahub/cmd/boltdb/boltdb /usr/bin/
CMD boltdb
