FROM golang:1.22rc2-bookworm

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN  go build -o /app/go_chain .cmd/web

CMD [ "make run" ]
