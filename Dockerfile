from golang:1.8.3

COPY . /data
RUN cd /data/example
RUN go run voice.go

