/*
* @Time    : 2020-08-20 16:09
* @Author  : CoderCharm
* @File    : request.go
* @Software: GoLand
* @Github  : github/CoderCharm
* @Email   : wg_python@163.com
* @Desc    :
**/
package utils

import (
	"fmt"
	"github.com/imroc/req"
	"log"
	"net/http"
	"strings"
)

// 设置请求头
var header = req.Header{
	"Accept":     "application/json",
	"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.125 Safari/537.36",
}

/*
获取 accessToken
@appId
@appSecret

**/
func GetAccessToken(appId string, appSecret string) string {

	url := fmt.Sprintf(`https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s`, appId, appSecret)

	resp, err := req.Get(url, header)
	if err != nil {
		log.Fatal(err)
	}
	resString := resp.String()

	fmt.Println(resString)

	// 转换map
	resMap := StringToMap(resString)

	// interface转string
	return resMap["access_token"].(string)

}

// json请求体
type ShortBody struct {
	AccessToken string `json:"access_token"`
	Action      string `json:"action"`
	LongUrl     string `json:"long_url"`
}

/*
获取短网址
@accessToken
@targetUrl 目标网址
**/
func GetShortUrl(accessToken string, targetUrl string) string {

	shortJson := ShortBody{accessToken, "long2short", targetUrl}

	url := fmt.Sprintf(`https://api.weixin.qq.com/cgi-bin/shorturl?access_token=%s`, accessToken)

	resp, err := req.Post(url, header, req.BodyJSON(&shortJson))

	if err != nil {
		log.Fatal(err)
	}
	resString := resp.String()

	// 字符串转
	resMap := StringToMap(resString)

	return resMap["short_url"].(string)
}

/*
检测短网址是否被封
@shortUrl 短网址
@return
	true 被封
	false 正常
**/
func CheckUrl(shortUrl string) bool{

	// 禁止重定向
	req.SetClient(&http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	})

	resp, err := req.Get(shortUrl, header)
	if err != nil {
		log.Fatal(err)
	}

	locationUrl := (*resp.Response()).Header["Location"][0]

	// 	fmt.Println((*resp.Response()).Header, reflect.TypeOf((*resp.Response()).Header), "\n响应")
	return strings.HasPrefix(locationUrl, "https://weixin110.qq.com/")


}
