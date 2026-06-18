FROM golang:latest

WORKDIR /notesapp

COPY . .

EXPOSE 8000

CMD go run main.go 