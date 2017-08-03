from golang:1.8.3

COPY . /go/src/github.com/chenqinghe/baidu-ai-go-sdk

WORKDIR /go/src/github.com/chenqinghe/baidu-ai-go-sdk/ocr
RUN go test
