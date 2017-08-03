from golang:1.8.3

COPY . /go/src/baidu-ai-go-sdk

WORKDIR /go/src/baidu-ai-go-sdk/ocr
RUN go test
