FROM golang:1.12-alpine 

WORKDIR /app

RUN apk update && apk add git

COPY . .
RUN go install

CMD ["library-inventory"]