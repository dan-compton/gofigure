FROM alpine:3.3

RUN apk add --update bash ca-certificates curl
ENV APPENV ""

CMD ["/gofigure"]
