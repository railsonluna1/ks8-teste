FROM golang:latest AS multistage

RUN go get golang.org/x/tools/cmd/cover

COPY . .

RUN go build -o=main

EXPOSE 8000

CMD [ "./main" ]