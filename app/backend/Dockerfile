FROM golang:1.22-alpine

WORKDIR /gobackend

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o backend .

EXPOSE 8000

CMD ["ls"]
CMD ["./backend"]
