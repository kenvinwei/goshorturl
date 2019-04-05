package baidu

import (
	"encoding/json"
	"errors"
	"goshorturl/config"
	"goshorturl/factory"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Baidu struct{}

type Data struct {
	Url string `json:"url"`
}

type Result struct {
	Code     int    `json:"Code"`
	ErrMsg   string `json:"ErrMsg"`
	ShortUrl string `json:"ShortUrl"`
	LongUrl  string `json:"LongUrl"`
}

func (b Baidu) MakeUrl(url string) (r factory.Result, err error) {

	if len(url) == 0 {
		return r, errors.New("invalid url")
	}
	return b.makeUrl(url)
}

func (b Baidu) makeUrl(url string) (ret factory.Result, err error) {
	var (
		req    *http.Request
		resp   *http.Response
		reader io.Reader
		buf    []byte
	)

	h := http.Header{}
	h.Set("Content-Type", "application/json")
	h.Set("Token", config.BaiduToken)

	//add header
	d := Data{Url: url}
	if buf, err = json.Marshal(d); err != nil {
		err = errors.New("baidu json.Marshal err:" + err.Error())
		return
	}

	reader = strings.NewReader(string(buf))

	if req, err = http.NewRequest(http.MethodPost, config.BaiduApiUrl, reader); err != nil {
		err = errors.New("baidu http.NewRequest err:" + err.Error())
		return
	}
	req.Header = h

	cli := http.Client{}
	if resp, err = cli.Do(req); err != nil {
		err = errors.New("baidu cli.Do err:" + err.Error())
		return
	}

	if buf, err = ioutil.ReadAll(resp.Body); err != nil {
		err = errors.New("baidu ioutil.ReadAll err:" + err.Error())
		return
	}

	defer resp.Body.Close()
	r := Result{}
	if err = json.Unmarshal(buf, &r); err != nil {
		err = errors.New("baidu json.Unmarshal err:" + err.Error())
		return
	}

	if r.Code != 0 {
		err = errors.New("baidu api err:" + r.ErrMsg)
		return
	}

	ret = factory.Result{
		LongUrl:  r.LongUrl,
		ShortUrl: r.ShortUrl,
	}

	return ret, err
}
