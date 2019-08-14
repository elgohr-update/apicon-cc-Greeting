FROM alpine:latest

ADD . /home/app/
WORKDIR /home/app
#ENTRYPOINT ["./src/greeting"]
ENTRYPOINT ["/bin/sh"]

EXPOSE 9001