FROM alpine:latest

ADD . /home/app/
WORKDIR /home/app

RUN chmod 777 /home/app/Greeting

ENTRYPOINT ["./Greeting"]
EXPOSE 8080