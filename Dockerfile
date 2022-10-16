FROM alpine

WORKDIR /usr/src/app

COPY /bin/http/main .

CMD [ "./main" ]