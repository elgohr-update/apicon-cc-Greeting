FROM alpine:latest

ADD Greeting /home/app/
WORKDIR /home/app/

RUN chmod 655 /home/app/Greeting

ENTRYPOINT ["./Greeting"]
EXPOSE 8080