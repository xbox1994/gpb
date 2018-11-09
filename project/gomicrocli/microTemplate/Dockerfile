FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app
ADD microTemplate .
ADD conf conf
RUN chmod +x microTemplate

ENTRYPOINT [ "./microTemplate" ]