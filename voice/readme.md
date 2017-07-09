# Installation
```GO
go get github.com/chenqinghe/baidu-ai-go-sdk/voice
```

# Usage

- 创建VoiceClient
```GO
    	var apiKey,apiSecret string 
	apiKey = "XXXXXXX"
	apiSecret = "XXXXXXX"
	client := NewVoiceClient(apiKey,apiSecret)
```

- 语音合成

```GO
	bts,err:=client.TextToSpeech("你好")
	if err!=nil{
		log.Fatal(err)
	}
	if err:=writeBytesToFile(bts,filename);err!=nil{//writeBytesToFile需要自己实现
		log.Fatal(err)
	}
```
- 语音识别 
```GO
	var ap ASRParams = ASRParams{
		Format:  "wav",
		Rate:    16000,
		Channel: 1,
		Cuid:    "mac address",
		Token:   client.AccessToken,
		Lan:     "zh",
		Speech:  afterBase64Str,
		Len:     fiLen,
	}
	strs,err:=client.SpeechToText(ap)
	if err!=nil{
		log.Fatal(err)
	}	
	fmt.Println(strs)//[]string，翻译候选结果，5个。
```
