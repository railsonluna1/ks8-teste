FROM golang:latest AS multistage

RUN go get golang.org/x/tools/cmd/cover

WORKDIR /src/main

COPY src/main .

RUN ls -la
RUN go build -o=main

EXPOSE 8000

CMD [ "./main" ]