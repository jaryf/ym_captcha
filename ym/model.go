package ym

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}

type Data struct {
	Code       int     `json:"code"`
	Time       float64 `json:"time"`
	UniqueCode string  `json:"unique_code"`
	CaptchaId  string  `json:"captchaId"`
	RecordId   string  `json:"recordId"`
	Data       string  `json:"data"`
}
