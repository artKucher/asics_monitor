FROM golang:1.24

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

WORKDIR /code
COPY . /code
RUN go build -o api *.go

CMD ["./api"]