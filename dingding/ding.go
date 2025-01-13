package dingding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// 钉钉 API 地址
const (
	BaseURL      = "https://oapi.dingtalk.com"
	TokenURL     = "/gettoken"
	TableContent = "/knowledge/get" // 示例接口路径，需查阅钉钉文档确认具体路径
)

// AppKey 和 AppSecret
const (
	AppKey    = "your_app_key"
	AppSecret = "your_app_secret"
)

// AccessTokenResponse 定义获取 Token 的响应结构
type AccessTokenResponse struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// TableRequest 请求获取表格内容的结构体
type TableRequest struct {
	TableID string `json:"table_id"`
}

// TableResponse 响应表格内容的结构体
type TableResponse struct {
	ErrCode int         `json:"errcode"`
	ErrMsg  string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}

func getAccessToken(appKey, appSecret string) (string, error) {
	url := fmt.Sprintf("%s%s?appkey=%s&appsecret=%s", BaseURL, TokenURL, appKey, appSecret)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var tokenResp AccessTokenResponse
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return "", err
	}
	if tokenResp.ErrCode != 0 {
		return "", fmt.Errorf("error getting token: %s", tokenResp.ErrMsg)
	}
	return tokenResp.AccessToken, nil
}

func getTableContent(accessToken, tableID string) (interface{}, error) {
	url := fmt.Sprintf("%s%s?access_token=%s", BaseURL, TableContent, accessToken)
	requestBody, _ := json.Marshal(TableRequest{TableID: tableID})
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var tableResp TableResponse
	if err := json.Unmarshal(body, &tableResp); err != nil {
		return nil, err
	}
	if tableResp.ErrCode != 0 {
		return nil, fmt.Errorf("error fetching table content: %s", tableResp.ErrMsg)
	}
	return tableResp.Data, nil
}
