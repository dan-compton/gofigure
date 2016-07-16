FROM alpine:3.4

ARG VERSION
ARG PROJECT
ARG ENTRYPOINT
ENV PROJECT  ${PROJECT:-""}
ENV VERSION  ${VERSION:-""}
ENV ENTRYPOINT ${ENTRYPOINT:-""}

RUN apk add --update ca-certificates

ADD ./bin\/${ENTRYPOINT} ${ENTRYPOINT}
ENTRYPOINT ${ENTRYPOINT}
