FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app
ADD ginTemplate .
ADD conf conf
RUN chmod +x ginTemplate

ENTRYPOINT [ "./ginTemplate" ]