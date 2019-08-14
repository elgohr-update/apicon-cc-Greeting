FROM alpine:latest

ADD . /home/app/
WORKDIR /home/app
ENTRYPOINT ["./src/greeting"]

EXPOSE 9001