
FROM golang:alpine AS builder
RUN apk add --no-cache git gcc libc-dev
WORKDIR /go/src/app
COPY . .
RUN go mod edit -module app
RUN go get -d -v ./...
RUN go install -v ./...

# final stage
FROM alpine:latest
LABEL Name=Quickfire Version=0.0.1
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT ./app
# EXPOSE 80


# FROM golang:1.23 as build
# WORKDIR /app

# COPY go.mod .
# COPY main.go .


# RUN go get
# RUN go build -o bin .
# ENTRYPOINT ["/app/bin"]

# FROM gcr.io/distroless/static-debian12
# COPY --from=build /go/bin/app /
# CMD ["/app"]