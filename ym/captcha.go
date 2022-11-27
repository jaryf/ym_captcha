package ym

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const CustomUrl = "https://www.jfbym.com/api/YmServer/customApi"
const OkCode = 10000
const DataOkCode = 0

type YmCaptcha struct {
	Token string
	h     http.Client
}

func NewYmCaptcha(token string) *YmCaptcha {
	return &YmCaptcha{
		Token: token,
		h:     http.Client{Timeout: time.Second * 10},
	}
}

// CommonVerify 通用验证
// # 数英汉字类型
// # 通用数英1-4位 10110
// # 通用数英5-8位 10111
// # 通用数英9~11位 10112
// # 通用数英12位及以上 10113
// # 通用数英1~6位plus 10103
// # 定制-数英5位~qcs 9001
// # 定制-纯数字4位 193
// # 中文类型
// # 通用中文字符1~2位 10114
// # 通用中文字符 3~5位 10115
// # 通用中文字符6~8位 10116
// # 通用中文字符9位及以上 10117
// # 定制-XX西游苦行中文字符 10107
// # 计算类型
// # 通用数字计算题 50100
// # 通用中文计算题 50101
// # 定制-计算题 cni 452
func (m *YmCaptcha) CommonVerify(image, captchaType string) (res string, err error) {
	config := map[string]interface{}{}
	config["image"] = image
	config["type"] = captchaType
	config["token"] = m.Token
	configBytes, _ := json.Marshal(config)
	bodyReader := bytes.NewReader(configBytes)
	request, err := http.NewRequest(http.MethodPost, CustomUrl, bodyReader)
	if err != nil {
		return
	}
	request.Header.Add("Content-Type", "application/json;charset=utf-8")
	response, err := m.h.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	resBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	var resData Result
	err = json.Unmarshal(resBytes, &resData)
	if err != nil {
		return
	}
	if resData.Code == OkCode && resData.Data.Code == DataOkCode {
		return resData.Data.Data, nil
	} else {
		return resData.Msg, fmt.Errorf("响应信息是%s", string(resBytes))
	}
}

// SlideVerify # 滑块类型
// # 通用双图滑块  20111
// # slide_image 需要识别图片的小图片的base64字符串
// # background_image 需要识别图片的背景图片的base64字符串(背景图需还原)
func (m *YmCaptcha) SlideVerify(slideImage string, backgroundImage string) (res string, err error) {
	req := map[string]interface{}{}
	req["slide_image"] = slideImage
	req["background_image"] = backgroundImage
	req["type"] = "20111"
	req["token"] = m.Token
	reqBytes, _ := json.Marshal(req)
	resp, err := http.Post(CustomUrl, "application/json;charset=utf-8", bytes.NewReader(reqBytes))
	defer resp.Body.Close()
	resBytes, _ := io.ReadAll(resp.Body)
	var resData Result
	err = json.Unmarshal(resBytes, &resData)
	if err != nil {
		return
	}
	if resData.Code == OkCode && resData.Data.Code == DataOkCode {
		return resData.Data.Data, nil
	} else {
		return resData.Msg, fmt.Errorf("响应信息是%s", string(resBytes))
	}
}

func (m *YmCaptcha) SinSlideVerify(image string) string {
	// # 滑块类型
	// # 通用单图滑块(截图)  20110
	config := map[string]interface{}{}
	config["image"] = image
	config["type"] = "20110"
	config["token"] = m.Token
	configData, _ := json.Marshal(config)
	body := bytes.NewBuffer([]byte(configData))
	resp, err := http.Post(CustomUrl, "application/json;charset=utf-8", body)
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data), err)
	return string(data)
}
