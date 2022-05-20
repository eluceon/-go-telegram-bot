FROM golang:1.18-alpine3.15 AS builder

WORKDIR /go/src/quizwithmebot

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o server ./cmd/server/...

FROM alpine:3.15
WORKDIR /user/local/bin

COPY --from=builder /go/src/quizwithmebot/server .

ENTRYPOINT ["./server"]