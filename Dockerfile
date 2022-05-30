FROM golang:latest
WORKDIR .
COPY . .

EXPOSE 8082
RUN chmod 777 main
RUN chmod 777 conf
CMD ./main
