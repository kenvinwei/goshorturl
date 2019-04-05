package sina

import (
	"encoding/json"
	"errors"
	"fmt"
	"goshorturl/config"
	"goshorturl/factory"
	"io/ioutil"
	"net/http"
)

type Sina struct{}

type Result struct {
	ErrorCode int `json:"error_code"`
	ShortUrl string `json:"url_short"`
	LongUrl  string `json:"url_long"`
}

type Results []Result

func (s Sina) MakeUrl(url string) (r factory.Result, err error) {
	if len(url) == 0 {
		return r, errors.New("invalid url")
	}
	return s.makeUrl(url)
}

func (s Sina) makeUrl(url string) (ret factory.Result, err error) {

	var (
		resp *http.Response
		buf  []byte
	)

	reqUrl := fmt.Sprintf("%s?source=%s&url_long=%s", config.SinaApiUrl, config.SinaAppKey, url)

	if resp, err = http.Get(reqUrl); err != nil {
		err = errors.New("sina http.Get err:" + err.Error())
		return
	}

	if buf, err = ioutil.ReadAll(resp.Body); err != nil {
		err = errors.New("sina ioutil.ReadAll err:" + err.Error())
		return
	}

	defer resp.Body.Close()
	r := Results{}

	if err = json.Unmarshal(buf, &r); err != nil {
		err = errors.New("sina json.Unmarshal err:" + string(buf))
		return
	}

	ret = factory.Result{
		LongUrl:  r[0].LongUrl,
		ShortUrl: r[0].ShortUrl,
	}

	return ret, err
}
