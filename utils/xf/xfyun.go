package xf

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type xfOcr struct {
	appId     string
	apiSecret string
	apiKey    string
	webAPI    string
	host      string
	h         http.Client
}

func NewXfOcr(appId string, apiSecret string, apiKey string, webAPI string, host string) *xfOcr {
	return &xfOcr{
		appId:     appId,
		apiSecret: apiSecret,
		apiKey:    apiKey,
		webAPI:    webAPI,
		host:      host,
		h: http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (m *xfOcr) xfAuthorization() {
	// now := time.Now()
	// date := now.Format(time.RFC1123)
	// api_key="$api_key",algorithm="hmac-sha256",headers="host date request-line",signature="$signature"

}

func (m *xfOcr) DefaultReqBody() *XfOcrReqBody {
	return &XfOcrReqBody{
		Header: XfOcrReqHeader{
			AppID:  m.appId,
			Status: 3,
		},
		Parameter: XfOcrReqParameter{
			Sf8E6Aca1: Sf8E6Aca1{
				Category: "ch_en_public_cloud",
				Result: Result{
					Encoding: "utf8",
					Compress: "raw",
					Format:   "json",
				},
			},
		},
		Payload: XfOcrReqPayload{
			Sf8E6Aca1Data1: Sf8E6Aca1Data1{
				Encoding: "jpg",
				Status:   3,
				Image:    "",
			},
		},
	}
}

func (m *xfOcr) parseUrl() (*Url, error) {
	stidx := strings.Index(m.webAPI, "://")
	host := m.webAPI[stidx+3:]
	schema := m.webAPI[:stidx+3]
	edidx := strings.Index(host, "/")
	if edidx <= 0 {
		return nil, errors.New("invalid request url:" + m.webAPI)
	}
	path := host[edidx:]
	host = host[:edidx]
	u := &Url{
		Host:   host,
		Path:   path,
		Schema: schema,
	}
	return u, nil
}

func (m *xfOcr) assembleWsAuthUrl(method string) (string, error) {
	u := Url{
		Host:   "api.xf-yun.com",
		Path:   "/v1/private/sf8e6aca1",
		Schema: "https://",
	}
	now := time.Now().UTC()
	date := now.Format(time.RFC1123)
	signature_origin := fmt.Sprintf("host: %s\ndate: %s\n%s %s HTTP/1.1", u.Host, date, method, u.Path)
	h := hmac.New(sha256.New, []byte(m.apiSecret))
	h.Write([]byte(signature_origin))
	signature_sha := base64.StdEncoding.EncodeToString(h.Sum(nil))
	authorization_origin := fmt.Sprintf("api_key=\"%s\", algorithm=\"%s\", headers=\"%s\", signature=\"%s\"", m.apiKey, "hmac-sha256", "host date request-line", signature_sha)
	authorization := base64.StdEncoding.EncodeToString([]byte(authorization_origin))
	v := url.Values{}
	v.Add("host", u.Host)
	v.Add("date", date)
	v.Add("authorization", authorization)
	v.Encode()
	params := v.Encode()
	reqUrl := fmt.Sprintf("%s?%s", m.webAPI, params)
	return reqUrl, nil
}

func (m *xfOcr) ImgOcr(reqBody *XfOcrReqBody) (wordList []string, err error) {
	reqUrl, err := m.assembleWsAuthUrl("POST")
	if err != nil {
		return
	}
	reqJsonByte, err := json.Marshal(reqBody)
	if err != nil {
		return
	}
	req, err := http.NewRequest(http.MethodPost, reqUrl, bytes.NewReader(reqJsonByte))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("host", "api.xf-yun.com")
	req.Header.Add("app_id", m.appId)
	resp, err := m.h.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var respInfo XfOcrRes
	err = json.Unmarshal(body, &respInfo)
	if err != nil {
		return
	}
	if respInfo.Header.Code != 0 {
		err = errors.New(respInfo.Header.Message)
		return
	}
	text := respInfo.Payload.Result.Text
	textByte, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return
	}
	var pageInfo XfOcrResPage
	err = json.Unmarshal(textByte, &pageInfo)
	if err != nil {
		return
	}
	for _, page := range pageInfo.Pages {
		if page.Exception != 0 {
			continue
		}
		for _, line := range page.Lines {
			if line.Exception != 0 {
				continue
			}
			for _, word := range line.Words {
				// 在这里可以判断置信度
				wordList = append(wordList, word.Content)
			}
		}
	}
	return
}

func (m *xfOcr) getImgBase64FromUrl(imgUrl string) (imgBase64 string, err error) {
	var (
		resp     *http.Response
		respByte []byte
	)
	resp, err = m.h.Get(imgUrl)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New("获取图片响应码非200")
		return
	}
	defer resp.Body.Close()
	respByte, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	imgBase64 = base64.StdEncoding.EncodeToString(respByte)
	return
}

func (m *xfOcr) getImgBase64FromPath(imgPath string) (imgBase64 string, err error) {
	var (
		f     *os.File
		fByte []byte
	)
	if f, err = os.Open(imgPath); err != nil {
		return
	}
	defer f.Close()
	if fByte, err = ioutil.ReadAll(f); err != nil {
		return
	}
	imgBase64 = base64.StdEncoding.EncodeToString(fByte)
	return
}

func (m *xfOcr) ImgOcrXfFromPath(imgPath string) (wordList []string, err error) {
	imgBase64, err := m.getImgBase64FromPath(imgPath)
	if err != nil {
		return
	}
	reqBody := m.DefaultReqBody()
	reqBody.Payload.Sf8E6Aca1Data1.Image = imgBase64
	return m.ImgOcr(reqBody)
}

func (m *xfOcr) ImgOcrXfFromUrl(imgUrl string) (wordList []string, err error) {
	imgBase64, err := m.getImgBase64FromUrl(imgUrl)
	if err != nil {
		return
	}
	reqBody := m.DefaultReqBody()
	reqBody.Payload.Sf8E6Aca1Data1.Image = imgBase64
	return m.ImgOcr(reqBody)
}
