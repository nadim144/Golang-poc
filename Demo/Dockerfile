## Simple
FROM golang:alpine
RUN mkdir /app 
ADD . /app/
WORKDIR /app 
RUN go mod download
COPY *.go ./
RUN go build -o svc .
EXPOSE 8080
CMD ["./svc"]





# # syntax=docker/dockerfile:1
# FROM golang:1.16-alpine
# WORKDIR /app
# COPY go.mod ./
# COPY go.sum ./
# RUN go mod download
# COPY *.go ./
# RUN go build -o /docker-gs-ping
# EXPOSE 8080
# CMD [ "/docker-gs-ping" ]




### Multistage
# FROM golang:alpine as builder
# RUN mkdir /build 
# ADD . /build/
# WORKDIR /build 
# RUN go build -o svc .
# FROM alpine
# COPY --from=builder /build/svc /app/
# WORKDIR /app
# CMD ["./svc"]

# # # Scratch
# FROM golang:alpine as builder
# RUN mkdir /build 
# ADD . /build/
# WORKDIR /build 
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o svc .
# FROM scratch
# COPY --from=builder /build/svc /app/
# WORKDIR /app
# CMD ["./svc"]