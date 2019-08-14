FROM alpine:latest

ADD . /home/app/
WORKDIR /home/app
ENTRYPOINT ["/bin/sh"]
#ENTRYPOINT ["./src/greeting"]

EXPOSE 9001