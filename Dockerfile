FROM golang:latest
RUN mkdir /range
ADD . /range/
WORKDIR /range
RUN go build -o main .
CMD ["/range/main"]