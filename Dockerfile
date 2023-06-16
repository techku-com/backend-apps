FROM golang:1.18-bullseye

ENV APP_NAME=techku-com
ENV GIT_TERMINAL_PROMPT=1

RUN apt-get update && \
    apt-get install -y make curl && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /usr/app

ADD . .
RUN go build main.go

EXPOSE 3000

CMD [ "./main", "--env", "Production_Docker"]
