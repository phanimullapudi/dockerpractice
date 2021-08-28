FROM alpine:latest
LABEL maintainer="Phani Mullapudi <phani.mullapudi@gmail.com>"
LABEL description="This is the docker file to have alpine with network tools and etc"
RUN apk add --update mysql-client
ENTRYPOINT ["tail"]
CMD ["-f", "/dev/null"]