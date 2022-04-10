FROM golang:1.18-alpine
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY . .
ENV GIN_MODE=release
RUN go mod download
RUN go build -o /student_app
RUN chmod +x /student_app
CMD [ "/student_app" ]