FROM golang:alpine as builder

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY ./scraper/ ./scraper

WORKDIR /build/scraper
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o scraper .

WORKDIR /dist

RUN cp /build/scraper/scraper .

CMD ["./scraper"]