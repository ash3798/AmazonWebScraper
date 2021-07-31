FROM golang:alpine as builder

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./persistence ./persistence

WORKDIR /build/persistence
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o persistence .

WORKDIR /dist

RUN cp /build/persistence/persistence .

CMD ["./persistence"]