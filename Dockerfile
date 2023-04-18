FROM golang:1.19-alpine AS compiler

ENV APP pifiabot

WORKDIR /go/src

ADD ./ /go/src/

RUN CGO_ENABLE=0 GOOS=linux go install ./...


FROM alpine

ARG BOT_TOKEN
ENV BOT_TOKEN ${BOT_TOKEN}
ENV ENDPOINT ""

WORKDIR /app

COPY --from=compiler /go/bin/pifiabot /app/

ADD ./deploy/cert.pem /app/
ADD ./deploy/private.key /app/
ADD ./deploy/messages.json /app/

CMD ["./pifiabot"]