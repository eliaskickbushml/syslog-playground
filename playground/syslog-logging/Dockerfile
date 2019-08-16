FROM alpine
RUN apk add --update syslog-ng
COPY syslog-ng.conf /etc/syslog-ng
VOLUME  ["/var/log/syslog-logs"]
EXPOSE 514/udp
EXPOSE 514/tcp
CMD /usr/sbin/syslog-ng && /bin/ash