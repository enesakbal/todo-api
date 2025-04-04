FROM golang

RUN mkdir /todo-app

COPY . /todo-app

WORKDIR /todo-app

RUN go build -o main cmd/main.go

CMD ["/todo-app/main"]

