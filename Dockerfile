from golang:1.8.3

COPY . /data
RUN ls /data
RUN cd /data/example
RUN go test
