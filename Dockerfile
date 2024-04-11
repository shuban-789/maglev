# Just a template as of now
FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping
# CHANGE PORT
EXPOSE 6553

CMD ["/docker-gs-ping"]
