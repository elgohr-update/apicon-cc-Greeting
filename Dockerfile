FROM alpine:latest

ADD . /home/app/
WORKDIR /home/app/src
ENTRYPOINT ["./greeting"]

EXPOSE 9001