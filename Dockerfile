from golang:1.8.3

COPY . /data

WORKDIR /data
RUN go test
