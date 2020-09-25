FROM golang:latest as builder

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN env CGO_ENABLED=0 go build -o webapp

# copy webapp from the builder and run it 
FROM alpine:latest

COPY --from=builder /go/src/app/webapp /webapp
CMD [ "./webapp" ]

