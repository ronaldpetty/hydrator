FROM golang:1.17 as builder

WORKDIR /go/src/github.com/myself/hydrator/

COPY . . 

RUN go get -u github.com/go-redis/redis/v8
RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/spf13/cobra
RUN go get -u github.com/spf13/viper
RUN go get -u github.com/spf13/pflag
RUN go mod tidy

RUN CGO_ENABLED=0 go build -o /tmp/hydrator .

FROM scratch
WORKDIR /
COPY --from=builder /tmp/hydrator .
USER 1000
ENTRYPOINT ["./hydrator"] 
