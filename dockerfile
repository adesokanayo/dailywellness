FROM golang:alpine as builder
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
LABEL name=Ayo
LABEL email=adesokanayo@gmail.com 
RUN mkdir /app
WORKDIR /app
COPY . .

RUN go get .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
RUN apk add --no-cache ca-certificates
FROM scratch
WORKDIR /home/innovation
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app .
CMD ["./app"]
