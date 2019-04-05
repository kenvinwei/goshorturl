package factory

import (
	"encoding/json"
	"log"
)

type Result struct {
	ErrCode  int    `json:"err_code"`
	ShortUrl string `json:"short_url"`
	LongUrl  string `json:"long_url"`
}

type Factory interface {
	MakeUrl(url string) (Result, error)
}

var (
	OK       = 0
	ErrParse = 1
)

/**
	callback:
	{
		"err_code":0,
		"short_url":"http://t.cn/8sFpC4t",
		"long_url":"http://z.qyer.com"
	}
**/
func Create(m Factory, url string) string {
	var (
		err error
		r   Result
		buf []byte
	)

	//default
	r.ErrCode = OK

	if r, err = m.MakeUrl(url); err != nil {
		log.Fatal(err.Error())
		r.ErrCode = ErrParse
	}

	if buf, err = json.Marshal(r); err != nil {
		panic(err)
	}

	return string(buf)
}
