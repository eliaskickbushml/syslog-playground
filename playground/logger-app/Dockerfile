FROM golang:latest
RUN mkdir /app
ADD ./app /app/
WORKDIR /app
RUN go get github.com/gin-gonic/gin
RUN go build -o main .
EXPOSE 5141
ENV SYSLOG_SERVER syslog-logging
ENV SYSLOG_PORT 5140
CMD ["/app/main"]