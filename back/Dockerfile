FROM golang:1.18

EXPOSE 8080

ADD . /back

WORKDIR /back
RUN go mod init operations
RUN go mod tidy
RUN go build -o operations .
RUN chmod +x operations

CMD ["./operations"]