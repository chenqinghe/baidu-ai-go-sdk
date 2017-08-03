from golang:1.8.3

COPY . /data

WORKDIR /data/ocr
RUN go test
