from golang:1.8.3

COPY . /data
RUN　cd /data
RUN ls

RUN go test
