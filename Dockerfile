FROM golang:alpine3.20
RUN go install github.com/air-verse/air@latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./


EXPOSE 8000
CMD [ "air" ]