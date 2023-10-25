FROM golang:1.21 

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY Makefile ./


RUN go mod download

COPY *.go ./


RUN make run

EXPOSE 8080

CMD ["/bin/GoBank"]
