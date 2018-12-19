FROM alpine
WORKDIR /usr/src/app
COPY . .
RUN chmod +x ginTemplate
EXPOSE 8080
CMD ./ginTemplate > /var/log/ginTemplate.log
