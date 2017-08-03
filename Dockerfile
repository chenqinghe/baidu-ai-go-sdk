from golang:1.8.3

COPY . /data
RUN ls /data
RUN cd /data/example
RUN go build /data/example/voice.go -o voice

ENTRYPOINT voice

