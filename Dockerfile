FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/mctrialgo
COPY . .
RUN go build .

FROM rocker/r-ver
WORKDIR /usr/src/mctrialgo
COPY --from=builder  /go/src/mctrialgo /usr/src/mctrialgo

RUN apt-get update
RUN apt-get install -y default-libmysqlclient-dev
RUN R -e "install.packages('RMySQL')"

# analysis.R、mctrialgo/repository.go admin.go server.goから参照される
ENV  COOKIE_SEED="cookie seed" \
     RSCRIPT="/usr/local/bin/Rscript" \
     MYSQL_HOST="mysql" \
     MYSQL_USER="oge" \
     MYSQL_PASSWORD="hogehogeA00" \
     MYSQL_DATABASE="studydb"

# docker run -it mctrialgo /bin/bash
# docker rmi mctrialgo
# docker image prune
