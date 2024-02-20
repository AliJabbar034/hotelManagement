FROM golang:1.21

WORKDIR /myapp

COPY go.mod go.sum ./

RUN go mod download && go mod verify 

COPY . .

RUN go build -o myapp main.go ./

EXPOSE 3000

CMD [ "./myapp" ]

