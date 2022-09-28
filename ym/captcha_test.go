package ym

import (
	"encoding/base64"
	"io/ioutil"
	"testing"
)

const (
	token = "ssssssssssss" // 你的token,从用户中心获取
)

func TestYmCaptcha_CommonVerify(t *testing.T) {
	readFileBytes, err := ioutil.ReadFile("../c1.jpg")
	if err != nil {
		t.Fatal(err)
	}
	imgBase64Str := base64.StdEncoding.EncodeToString(readFileBytes)
	var y = NewYmCaptcha(token)
	verify, err := y.CommonVerify(imgBase64Str, "10111")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("验证结果是", verify)
}
