FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN go build -o /basic-auth-app

EXPOSE 8010

CMD [ "/basic-auth-app" ]