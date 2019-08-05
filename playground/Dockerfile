FROM alpine
RUN apk add --update syslog-ng
COPY syslog-ng.conf /etc/syslog-ng
VOLUME [ "/var/log/syslog-logs"]
EXPOSE 514/tcp
EXPOSE 514/udp
ENTRYPOINT /usr/sbin/syslog-ng -F