package main

import (
	"bytes"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"lucky/core/iencrypt/little"
	"lucky/core/iproto"
	"lucky/example/comm/msg"
	"lucky/example/comm/msg/code"
	"lucky/example/comm/protobuf"
	"lucky/log"
	"net/http"
)

func main() {
	client := http.DefaultClient
	helloMsg, err := proto.Marshal(&protobuf.Hello{Hello: "Http test"})
	if err != nil {
		panic(err)
	}
	body, err := proto.Marshal(&iproto.Protocol{Id: code.Hello, Content: helloMsg})
	if err != nil {
		panic(err)
	}
	// 加密
	pwd, err := little.ParsePassword(msg.PwdStr)
	if err != nil {
		panic(err)
	}
	cipher := little.NewCipher(pwd)
	body = cipher.Encode(body)
	// 请求
	req, err := http.NewRequest("POST", "http://localhost:3001", bytes.NewReader(body))
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// 解密
	data = cipher.Decode(data)
	// unmarshal
	var ipro iproto.Protocol
	err = proto.Unmarshal(data, &ipro)
	if err != nil {
		log.Fatal("cant Unmarshal data %v , err %v", data, err)
	}
	log.Debug("received msg %+v", ipro.Id)
}