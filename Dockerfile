FROM node AS ui
WORKDIR /go/metahub
COPY ./ui/package* ./ui/
RUN cd ui && npm install
COPY ./ui ./ui
COPY ./static ./static
COPY ./templates ./templates
WORKDIR /go/metahub/ui/
RUN npm run build

FROM golang:1.12 AS go
WORKDIR /go/metahub
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download
COPY ./cmd ./cmd
COPY ./pkg ./pkg
WORKDIR /go/metahub/cmd/boltdb
# static build
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -a -ldflags '-extldflags "-static"' .

# Go binary serves the ui web content
FROM scratch
COPY --from=go /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=ui /go/metahub/static /srv/html/static
COPY --from=ui /go/metahub/templates/gen/index.html /srv/html/
COPY --from=go /go/metahub/cmd/boltdb/boltdb /usr/bin/
WORKDIR /usr/bin/
CMD boltdb
